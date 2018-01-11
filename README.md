# Go kit Intro
[Go kit](https://gokit.io/) is a self described toolkit for Microservices.
It defines a set of tooling to help in building Microservices that follow good patterns and avoid common pitfalls.

The basic structure of Go kit defines 3 layers which interact through well defined interfaces:
1. Transport layer - HTTP and gRPC support are a couple of included transports. Multiple can be supported in a single service.
2. Endpoint layer - Handlers for RPC in the service are in this layer.
3. Service layer - Contains the business logic and core service definition.

The clean separation of layers which in turn allows for writing of generic components for tasks such as logging and monitoring is the key advantage of this framework.
Additional functionality is added to a service by using a middleware to wrap endpoints or services as needed.
Each layer can depend on inner layers but should never depend on the outer layers.
This graphic from the [FAQ](https://gokit.io/faq/onion.png) is a good representation.
![Go kit model](https://gokit.io/faq/onion.png)
 

## Demo Code

I wrote up a simple API service where you can see the layers in action. The sections to note are:

- The daemon is setup in main.go
- The service code and service level middleware in the service directory.
  Service middleware follows the service interface and so adds a layer that is knowledgeable about the service methods.
- See endpoint.go and transport.go for code related to those layers.

There are some more complete demos available on the [Go kit site](https://gokit.io/examples/)
