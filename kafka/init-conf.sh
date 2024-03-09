docker exec <name-conatiner> /opt/kafka/bin/kafka-topics.sh --create --topic microservice-topic --bootstrap-server localhost:9092

docker exec <name-conatiner> /opt/kafka/bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic microservice-topic --group microservice-group