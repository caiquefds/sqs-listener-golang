# SQS Listener (Golang + LocalStack)

This project is a **Golang-based SQS listener** that consumes messages from an AWS SQS queue, processes them through a defined service, and acknowledges (deletes) them after successful processing.  
It is designed to run against **LocalStack** for local development.

---

## 📌 Features

- Connects to SQS  
- Receives messages continuously  
- Processes messages through a custom service (`OrderService`)  
- Deletes messages after processing (acknowledgement)  
- Uses Goroutines + Channels for concurrency  
---
## 🛠️ Requirements

- Go 1.20+
- LocalStack
- AWS CLI or awslocal

---

## 🚀 Running

Install dependencies:
```bash
go mod tidy
```

Run localstack:

```bash
docker compose up
```
Run application:

```bash
go run main.go
```
