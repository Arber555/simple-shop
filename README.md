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
make all DOCKER_USERNAME=your-dockerhub-username
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


## Guides

These guides provide a step-by-step approach to deploying, configuring, and securing the SimpleShop web application in a Kubernetes environment.

### [Chapter 1: Running and Configuring SimpleShop in Kubernetes](guides/chapter1.md)

1. **[Prerequisites](guides/chapter1.md#prerequisites)**
   - Ensures Docker and Kubernetes environment are set up, and kubectl is installed.

2. **[Part 1: Run the SimpleShop Web App in Kubernetes](guides/chapter1.md#part-1-run-the-simpleshop-web-app-in-kubernetes)**
   - Guides through deploying the core components of SimpleShop (web app, APIs, database) as Pods in Kubernetes.

3. **[Part 2: Route Network Traffic into the Website and the APIs](guides/chapter1.md#part-2-route-network-traffic-into-the-website-and-the-apis)**
   - Shows how to expose internal services using `ClusterIP` and external services using `NodePort` or `LoadBalancer`.

4. **[Part 3: Configure the Application Using Kubernetes Resources](guides/chapter1.md#part-3-configure-the-application-using-kubernetes-resources)**
   - Explains using ConfigMaps and Secrets to configure the application dynamically.

5. **[Part 4: Update the App to Run with High Availability and Scale](guides/chapter1.md#part-4-update-the-app-to-run-with-high-availability-and-scale)**
   - Focuses on scaling and ensuring high availability using Kubernetes Deployment objects.

---

### [Chapter 2: Enhancing Reliability, Security, and Storage for SimpleShop with Kubernetes](guides/chapter2.md)

1. **[Using Probes to Make SimpleShop Self-Healing](guides/chapter2.md#using-probes-to-make-simpleshop-self-healing)**
   - Details adding liveness and readiness probes to ensure services restart on failure and only healthy Pods receive traffic.

2. **[Enhanced Security for SimpleShop Containers](guides/chapter2.md#enhanced-security-for-simpleshop-containers)**
   - Introduces security practices like disabling Kubernetes API token access, limiting resources, and running containers as non-root users.

3. **[Using Volumes for Storage](guides/chapter2.md#using-volumes-for-storage)**
   - Covers how to surface logs with ephemeral storage and configure persistent storage for the Product Database.
