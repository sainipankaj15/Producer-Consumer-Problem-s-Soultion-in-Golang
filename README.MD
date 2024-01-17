# Producer-Consumer Problem

The **Producer-Consumer problem** is a classic synchronization challenge in computer science and concurrent programming. This scenario involves two types of processes: producers and consumers, which interact with a shared, fixed-size buffer or queue. The objective is to enable producers and consumers to work concurrently without causing conflicts or data corruption.

## Terminology:

- **Producers**: Entities or processes responsible for generating data or items and adding them to the shared buffer.

- **Consumers**: Entities or processes tasked with consuming or removing items from the shared buffer.

- **Buffer**: A fixed-size storage area shared between producers and consumers. The buffer has a limited capacity, and coordination is essential to prevent overproduction or underconsumption issues.

## Challenges:

1. **Synchronization**: Ensuring that producers and consumers do not access the buffer simultaneously to prevent data corruption or conflicts.

2. **Communication**: Coordinating between producers and consumers to prevent overfilling by producers or emptying by consumers when the buffer is already empty.

## Illustration:

The Producer-Consumer problem serves as a scenario to demonstrate the challenges of concurrent programming and underscores the significance of synchronization mechanisms in multi-threaded or multi-process environments.

## Solution Approach:

Common solutions involve the use of locks, semaphores, or other synchronization mechanisms to control access to the shared buffer. We are going to use different go routine for both Producer and Consumers and will communticate via channels.
We will take a example of a pizza shop where a new order will work as a producer and delivery the order(doesn't matter it is successully made or not) will work as a consumer. Means in the consumer side , we will read the status of Order. So let's go ...

## How to run 
- Clone the repo.
- Run the command ```go mod tidy ```
- After all module successfully installed , run the command ```go run .```
- Output will command on the terminal with appropriate colors.