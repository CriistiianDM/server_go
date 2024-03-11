/**
 *
 * @autor: Cristian Machado <cristian.machado@correounivalle.edu.co>
 * @copyrigth: 2024
 * @license: GPL-3.0
*/
package kafkaPlugin

import (
    "fmt"
	"os"
	"os/signal"
	"encoding/json"

	"server_go/kafka/kafkaPlugin/utils"
    "github.com/confluentinc/confluent-kafka-go/kafka"
)

/**
  * Kafka Conf
*/
func KafkaConsumer(args map[string]interface{}) (*kafka.Consumer, error) {
	var (
		consumer *kafka.Consumer
		is_error error
		key_available = []string{"host","group","topic"}
	)

	if utils.CheckExpectedKeys(args,key_available) {

		consumer, is_error = kafka.NewConsumer(&kafka.ConfigMap{
			"bootstrap.servers": args["host"].(string),
			"group.id": args["group"].(string),
			"auto.offset.reset": "earliest",
		})

		if is_error == nil {
			consumer.SubscribeTopics([]string{args["topic"].(string)}, nil)
        }
	} else {
		fmt.Println("No have all keys for create conf kafka")
	}

	return consumer, is_error
}

/**
  * Listener changes menssages
*/
func KafakaListener(args map[string]interface{}) {
	consumer := args["consumer"].(*kafka.Consumer);
	is_error := args["is_error"];
	run := true
	
	if is_error == nil {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, os.Interrupt)
	
		// Loop para manejar mensajes y señales
		for run {
			select {
				case sig := <-signals:
					fmt.Printf("Terminating due to signa: %v\n", sig)
					run = false
		
				default:
					ev := consumer.Poll(100)
					if ev == nil {
						continue
					}
		
					switch e := ev.(type) {
						case *kafka.Message:
							fmt.Printf("Message received: %s\n", e.Value)

						case kafka.Error:
							fmt.Fprintf(os.Stderr, "Error: %v\n", e)
							run = false
			
						default:
							fmt.Printf("Event ignored: %v\n", e)
					}
			}
		}
	}
}

/**
  * Send menssage to topic in formatter json
*/
func KafkaSendMennsage(args map[string]interface{}) {
	var (
		key_available = []string{"host","message","topic"}
		producer *kafka.Producer
		jsonMessage []uint8
		is_error error
	)

	if !utils.CheckExpectedKeys(args,key_available) {
		fmt.Println("No have all keys")
		return
	}

	producer, is_error = kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": args["host"].(string),
	})

	if is_error == nil {
		go manageEventsProducer(map[string]interface{}{
			"producer": producer,
		})

		topic := args["topic"].(string)
		message := args["message"].(map[string]interface{})
		jsonMessage, is_error = json.Marshal(message)

		if is_error == nil {
			producer.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
				Value: jsonMessage,
			}, nil)
	
			producer.Flush(15 * 1000)
		}
	}

	if is_error != nil {fmt.Println(is_error)}
}

/**
  * Event of producer
*/
func manageEventsProducer(args map[string]interface{}) {
  producer := args["producer"].(*kafka.Producer)
  for e := range producer.Events() {
    switch ev := e.(type) {
        case *kafka.Message:
            if ev.TopicPartition.Error != nil {
              fmt.Printf("Error send menssage: %v\n", ev.TopicPartition.Error)
            } else {
              fmt.Printf("Mesage send to %s [%d] in the position %d\n",
                *ev.TopicPartition.Topic, ev.TopicPartition.Partition, ev.TopicPartition.Offset)
            }
    }
  }
}