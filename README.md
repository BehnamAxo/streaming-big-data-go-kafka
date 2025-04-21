# Streaming Big Data with Go + Kafka

This project is a simple example of producing and consuming messages with Go using Apache Kafka. It also includes a Kafka UI (Kafdrop) to help visualize what’s happening under the hood.


## 🚀 How to Run This Project

Follow these steps to get everything up and running smoothly:

### 1. Make sure you have the basics installed:
- **Docker** and **Docker Compose**
- **Go** installed locally (`go run` needs to work in your terminal)

---

### 2. Start the Kafka stack using Docker

Spin up Zookeeper, Kafka, and Kafdrop:

```bash
docker-compose up --build
```

> This will start all the containers and expose:
> - Kafka on `localhost:9092`
> - Kafdrop (Kafka UI) on `localhost:9000`

Wait a few seconds until you see logs showing the services are up.

---

### 3. Create the `logs` topic (if not auto-created)

If Kafka auto topic creation is off or flaky, you can manually create the topic:

```bash
docker exec -it streaming-big-data-go-kafka-kafka-1 \
kafka-topics --create \
--topic logs \
--bootstrap-server localhost:9092 \
--partitions 1 \
--replication-factor 1
```

---

### 4. Run the Producer

In a new terminal window, run the producer:

```bash
go run producer/main.go
```

This will start generating fake messages like `"User 123 clicked button"` and push them to the `logs` topic.

---

### 5. Run the Consumer

In another terminal, run the consumer:

```bash
go run consumer/main.go
```

This will consume messages from Kafka and log how many it processes every second.

---

### 6. Open Kafdrop UI (Optional but Awesome)

Visit [http://localhost:9000](http://localhost:9000) in your browser to inspect Kafka topics and messages.

---

### 7. Stop Everything When You're Done

To stop the Kafka stack:

```bash
docker-compose down
```
