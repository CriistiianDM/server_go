#Kafka image
FROM apache/kafka:3.7.0

# Add Kafka to PATH
ENV PATH="/opt/kafka/bin:${PATH}"

# Create Directory
WORKDIR /kafka


# Generate 
RUN mkdir -p /kafka/config

# Copy configurations
# COPY ./conf/server.properties /kafka/config/server.properties # For build image
# COPY ./conf/consumer.properties /kafka/config/consumer.properties # For build image
COPY kafka/conf/server.properties /kafka/config/server.properties
COPY kafka/conf/consumer.properties /kafka/config/consumer.properties

# Generate SSL certificates
RUN keytool -genkeypair -keyalg RSA -alias kafka-ssl -keystore /kafka/config/keystore.jks -storepass keystore_password -keypass key_password -validity 365 -keysize 2048 -dname "CN=localhost" -ext SAN=dns:localhost

RUN keytool -export -alias kafka-ssl -keystore /kafka/config/keystore.jks -file /kafka/config/kafka_ssl.cer -storepass keystore_password

RUN keytool -import -alias kafka-ssl -file /kafka/config/kafka_ssl.cer -keystore /kafka/config/truststore.jks -storepass truststore_password -noprompt

# Init Conf
ENV KAFKA_OPTS -Djava.security.auth.login.config=/kafka/config/server.properties

# Port
EXPOSE 9092