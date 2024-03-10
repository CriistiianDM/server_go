/**
 *
 * @autor: Cristian Machado <cristian.machado@correounivalle.edu.co>
 * @copyrigth: 2024
 * @license: GPL-3.0
*/
package kafkaPlugin

import (
    "fmt"
	
	"server_go/kafka/kafkaPlugin/utils"
    "github.com/confluentinc/confluent-kafka-go/kafka"
)

/**
  * Kafka Conf
*/
func KafkaConf(args map[string]interface{}) {

	if utils.CheckExpectedKeys(args,[]string{"host","group_id","topic"}) {
		consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
			"bootstrap.servers": args["host"],
			"group.id": args["group_id"],
			"auto.offset.reset": "earliest",
		})
		fmt.Println(consumer,err)
	} else { 
		fmt.Println("No have all keys for create conf kafka")
	}
}