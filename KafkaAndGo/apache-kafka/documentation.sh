docker run -it -p 9092:9092 --privileged  apache/kafka:3.7.0 bash

opt/kafka/bin/zookeeper-server-start.sh config/zookeeper.properties

/opt/kafka/bin/kafka-topics.sh --create --topic quickstart-events --bootstrap-server localhost:9092

