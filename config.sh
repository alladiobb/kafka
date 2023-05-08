#instalando o kafka
#denro do diretório do arqquivo docker-compose.yaml:
docker-compose up -d

#pegue o nome em que o kafka está rodando
docker-compose ps

#Execute o comando para entrar no TERMINAL do kafka
docker exec -it kafka-kafaka-1 bash

#Comando que serão usados:
#bootstrap-server
#topic
#create
#partition
#name

kafka-topics --create --topic=teste --bootstrap-server=localhost:9092 --partitions=3
