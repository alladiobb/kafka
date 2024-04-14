#instalando o kafka
#denro do diretório do arqquivo docker-compose.yaml:
docker compose up -d

#pegue o nome em que o kafka está rodando
docker compose ps

#Execute o comando para entrar no TERMINAL do kafka
docker exec -it kafkadocker-kafka-1 bash

#entendendo o kafka-topics
kafka-topics

#Comando que serão usados:
#bootstrap-server
#topic
#create
#partition
#name

#criado o tópico
kafka-topics --create --topic=teste --bootstrap-server=localhost:9092 --partitions=3

#Listando o tópico criado e outros que existem
kafka-topics --list --bootstrap-server=localhost:9092

#descrevendo o tópico
kafka-topics --topic=teste --bootstrap-server=localhost:9092 --describe

# Topic: teste    PartitionCount: 3       ReplicationFactor: 1    Configs:
#         Topic: teste    Partition: 0    Leader: 1       Replicas: 1     Isr: 1
#         Topic: teste    Partition: 1    Leader: 1       Replicas: 1     Isr: 1
#         Topic: teste    Partition: 2    Leader: 1       Replicas: 1     Isr: 1


#Ouvindo o consumidor
kafka-console-consumer --bootstrap-server=localhost:9092 --topic=teste

#produtor
kafka-console-producer --bootstrap-server=localhost:9092 --topic=teste

#criando o consumidor e lendo desde o começo as mensagens
kafka-console-consumer --bootstrap-server=localhost:9092 --topic=teste --from-beginning

#Consumer Groups
kafka-console-consumer --bootstrap-server=localhost:9092 --topic=teste group=x

#Descrição dos consumers
kafka-consumer-groups --bootstrap-server=localhost:9092 --group=x --describe