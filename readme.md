# A Practical Introduction to the Event Driven Architecture using Redis Pub/Sub

![Event Driven Architecture](https://www.cncf.io/wp-content/uploads/2023/10/Screenshot-2023-10-27-at-16.36.16.png)

> THIS ARTICLE IS STILL A WORK IN PROGRESS

After working with microservices on a variety of projects, I have developed a strong appreciation for a specific software design architecture, The Event Driven Architecture(EDA). The Event Driven Architecture is a software design pattern that allows systems to detect, process, manage, and react to real-time events as they happen. An event-driven architecture uses events to trigger and communicate between decoupled services. This ability to communicate between decoupled services is why the EDA is perfect for microservices.

In this article, I will explore event-driven architecture in depth, with practical code snippets that demonstrate how to utilize Redis's Pub/Sub functionality.

Redis Pub/Sub (publish/subscribe) is a messaging pattern provided by Redis that facilitates communication between different components of a system. In this pattern, publishers send messages to specific channels without needing to know who receives them, while subscribers listen to these channels to receive the messages in real-time. This decoupling of publishers and subscribers makes Redis Pub/Sub a lightweight and efficient solution for building event-driven system

## How Event-Driven Architecture Works

At a high level point of view the EDA is quite straight forward. Events representing occurrences or changes in the system drive the flow are generated by various sources (Publishers), published to an event bus or message broker, and consumed by interested components (Consumers) asynchronously. This approach promotes flexibility, scalability, and resilience.

An event is any change in state of the system and can be anything ranging from customer requests, inventory updates, sensor readings, and so on.

The event-driven architecture enables different components of a system to run independently of each other by introducing a middleman known as an **Event broker** that routes events to different intended destination components. This means that applications and devices do not need to know where they are sending information or where the information they are consuming comes from. Rather they focus on carrying out there intended functionality.

In Redis Pub/Sub, we have a *Publisher* that publishes messages across different *Message Channels* and *Consumers* that can listen to different message channels for events/messages sent.

![Image showing producer](https://jim-junior.github.io/eda1.jpg)

> Note: In this article we use Redis's Pub/Sub because of its simplicity and realatively easy learning curve but Redis Pub/Sub wont demonstrate the full functionality of an Event Driven System, however there are other more complex software systems such as RabbitMQ or Apache Kafka that provide a wide range of functionality and Protocals such as AMQP and MQTT(Commonly used in IoT application) that can be used to build more large scale systems, however these have a steep learning curve and wont be ideal for demonstration purposes in this article.




## References

Event-Driven Architecture. Amazon Web Services. [https://aws.amazon.com/event-driven-architecture/](https://aws.amazon.com/event-driven-architecture/)

Event-Driven Architecture (EDA) - A Complete Introduction. Confluent. [https://www.confluent.io/learn/event-driven-architecture/](https://www.confluent.io/learn/event-driven-architecture/)

The Complete Guide to Event-Driven Architecture. Medium. [https://medium.com/@seetharamugn/the-complete-guide-to-event-driven-architecture-b25226594227](https://medium.com/@seetharamugn/the-complete-guide-to-event-driven-architecture-b25226594227)
