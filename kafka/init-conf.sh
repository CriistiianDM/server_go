docker exec <nam_container> /opt/kafka/bin/kafka-topics.sh --create --topic microservice-topic --bootstrap-server localhost:9092

docker exec <name_container> /opt/kafka/bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic microservice-topic --group microservice-group

# all menssage in topic
# docker exec kafka-server /opt/kafka/bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic microservice-topic --from-beginning