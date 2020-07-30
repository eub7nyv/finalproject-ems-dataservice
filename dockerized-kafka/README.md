# Kafka running in Docker
This project is to get [Kafka](https://kafka.apache.org/intro) running locally in Docker.  Notes below.

# Running
By default, the `docker-compose` will run a single Zookeeper and Kafka instance.
- Uses the [Zookeeper Docker Image](https://hub.docker.com/_/zookeeper)
- Builds a Kafka image (see `kafka-broker`)

## Compose run command:
`docker-compose up` or `docker-compose up -d` for a daemonized version --> this creates the 2 image : zookeeper and dockerized-kafka_kafka
It also run the above two images in  ports 0.0.0.0:9092->9092/tcp and 0.0.0.0:2181->2181/tcp, 0.0.0.0:2888->2888/tcp, 0.0.0.0:3888->3888/tcp, 0.0.0.0:8080->8080/tcp respectively

## With test box
The Kakfa install comes with a set of test scripts.  You can add a test box with:
```docker-compose -f docker-compose.yml -f docker-compose.tester.yml up -d```
This will create and run an image dockerized-kafka_tester

Then you can validate connectivity by running:
```docker exec dockerized-kafka_tester_1 bin/kafka-topics.sh --create --bootstrap-server kafka:9092 --replication-factor 1 --partitions 1 --topic test```
The above command will create a topic called test. 

to create a test topic and then
```docker exec dockerized-kafka_tester_1 bin/kafka-topics.sh --list --bootstrap-server kafka:9092```

to see it show the `test` topic.

# License
This project is licensed under Apache-2.0.
