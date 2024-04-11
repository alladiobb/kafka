#instalando o kafka
#denro do diretório do arqquivo docker-compose.yaml:
docker compose up -d

#pegue o nome em que o kafka está rodando
docker compose ps

#Execute o comando para entrar no TERMINAL do kafka
docker exec -it kafkadocker-kafka-1 bash

#Comando que serão usados:
#bootstrap-server
#topic
#create
#partition
#name

#criado o tópico
kafka-topics --create --topic=teste --bootstrap-server=localhost:9092 --partitions=3

#descrevendo o tópico
kafka-topics --topic=teste --bootstrap-server=localhost:9092 --describe

#criando o consumidor
kafka-console-consumer --bootstrap-server=localhost:9092 --topic=teste

#produtor
kafka-console-producer --bootstrap-server=localhost:9092 --topic=teste

#criando o consumidor e lendo desde o começo as mensagens
kafka-console-consumer --bootstrap-server=localhost:9092 --topic=teste --from-beginning
