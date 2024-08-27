# SimpleShop Application

## Overview

The `SimpleShop` is a sample microservice-based e-commerce application. It consists of several services that work together to form the SimpleShop backend and frontend.

### Architecture

The application is built using a microservices architecture and consists of the following services:

- **Web Service**: The frontend of the SimpleShop.
- **Products API**: A backend service that manages product listings, prices, and descriptions.
- **Stock API**: A backend service responsible for managing product availability and stock levels.
- **Products Database**: A SQL-based database service that stores product data.

Each service is containerized and can be run independently using Docker and Kubernetes.

## Prerequisites

To build and run the application, you need to have the following installed:

- [Docker](https://docs.docker.com/get-docker/)
- [Kubernetes (kubectl)](https://kubernetes.io/docs/tasks/tools/install-kubectl/)
- [Make](https://www.gnu.org/software/make/)

## Setup and Installation

### Building images

```
make all
```
This will build the following Docker images:

- Web: your-dockerhub-username/web
- Products API: your-dockerhub-username/products-api
- Stock API: your-dockerhub-username/stock-api
- Products Database: your-dockerhub-username/products-db

You can also build individual images with:

```
make build-web
make build-products-api
make build-stock-api
make build-products-db
```

Push Docker Images to Docker Hub
Once your images are built, you can push them to Docker Hub:

```
make push-all DOCKER_USERNAME=your-dockerhub-username
```

Alternatively, push specific images:

```
make push-web DOCKER_USERNAME=your-dockerhub-username
```

Ensure you have logged into Docker Hub via docker login before pushing.