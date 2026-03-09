Ecommerce Microservices Gateway API

This project is an API Gateway for an Ecommerce system that manages communication between multiple microservices, including UserAuth, Product, and Order services. The gateway handles client requests, service-to-service communication, and event management in a scalable and maintainable way.

Features

REST API – Exposes endpoints to clients for user authentication, product browsing, and order management.

Rabbitmq message broker - publish a user created event whcih helps notification service 

gRPC – Facilitates efficient communication between microservices (e.g., checking stock from Product service when creating an order).
