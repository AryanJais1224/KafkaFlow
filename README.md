# KafkaFlow

KafkaFlow is a minimal event-driven system built using Go and Apache Kafka. It demonstrates real-world streaming patterns such as retry handling and dead letter queues while keeping the codebase simple and easy to understand.

## Features

* Event production and consumption
* Consumer groups
* Retry mechanism
* Dead letter queue
* JSON-based messaging
* Concurrent processing using goroutines

## Architecture

The system uses three Kafka topics:

* events
* events-retry
* events-dlq

Flow:

1. Producer sends events to the events topic
2. Consumer processes events
3. Failed events are sent to events-retry
4. After multiple failures, events go to events-dlq

## Tech Stack

* Go
* Apache Kafka
* kafka-go library

## How it Works

* Producer continuously generates events
* Consumer reads and processes events
* Random failures simulate real-world scenarios
* Retry logic ensures reprocessing
* DLQ captures permanently failed events

## Use Cases

* Order processing systems
* Payment pipelines
* Notification systems
* Log processing pipelines

