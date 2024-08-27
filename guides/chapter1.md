# Kubernetes Learning Guide: Chapter 1

## Prerequesite

Before starting with Kubernetes, ensure your environment is set up with the following:

1. Install Docker
- Download and install Docker from Docker's website.
- Enable Kubernetes in Docker Desktop (under Settings > Kubernetes).
2. Install kubectl
- Install kubectl following the [guide](https://pwittrock.github.io/docs/tasks/tools/install-kubectl/).
- Verify the installation with:

```
kubectl version --client
```

## Part 1: Run the SimpleShop Web App in Kubernetes

**Objective**
Let's run some containers in Kubernetes! This is the architecture of the SimpleShop web application you'll be building out:

There’s the main website, two backend REST APIs, and a PostgreSQL database. The goal here is to run each component in Kubernetes, starting with the compute layer:

**Workflow**
Model the application in Kubernetes YAML
You need to define a Pod specification in YAML for each component. There are four components:

- `stefanprifti/simpleshop-products-db:v1`
- `stefanprifti/simpleshop-products-api:v1`
- `stefanprifti/simpleshop-stock-api:v1`
- `stefanprifti/simpleshop-web:v1`

The Pod metadata should include unique names for each Pod and labels indicating they are part of the SimpleShop app.

1. Deploy the application to your cluster
Use `kubectl apply -f` for each of your YAML files. Kubernetes will download the container images from Docker Hub and start the containers in separate Pods.

2. Verify that the application components are running
Use kubectl get pods to confirm that all Pods are in the "Running" state. It's okay if the products API Pod restarts because it can’t access the database yet.

3. Print the application logs
Check the logs of each Pod using `kubectl logs <pod-name>`. You might see errors—this is expected at this stage.

4. Prove the reliability of Pods
Test Kubernetes' self-healing feature by killing the process in the simpleshop-stock-api Pod. Run `kubectl exec <pod-name> -- kill 1`. The Pod should restart automatically.

5. Do not forget to set resources limits for each container.

**Deliverable**
- YAML files modeling the Pods
- The kubectl commands used to deploy and test the app

## Part 2: Route Network Traffic into the Website and the APIs

**Objective**
Now that the individual components are running in Pods, let’s connect them through the network layer.

**Workflow**
- Model internal routes for each component
- Use Kubernetes Service objects to create stable endpoints:

For the PostgreSQL database: DNS simpleshop-products-db, port 5432
For the products API: DNS simpleshop-products-api, port 80
For the stock API: DNS simpleshop-stock-api, port 80
Expose the web app to external traffic
Model a Service to expose the web component externally. The web app listens on port 80, and you can map it to an external port of your choice.

**Deploy the Services**
Apply your Service definitions using `kubectl apply -f`.

**Test the application**
Browse to the external URL of the web component. If everything is working, you should see the SimpleShop homepage.

**Deliverable**
- YAML files modeling the Services
- The kubectl commands used to deploy and test the app

## Part 3: Configure the Application Using Kubernetes Resources

**Objective**
Now, configure the application dynamically using Kubernetes objects like ConfigMaps and Secrets.

**Workflow**
- Configure the database password
- Store the database password in a Secret, and modify the simpleshop-products-db Pod spec to reference the Secret as an environment variable.

**Set API configurations**
Store non-sensitive config files in ConfigMaps and sensitive files in Secrets. For example:
- `simpleshop-products-api` should receive database connection settings via environment variables.
- `simpleshop-stock-api`  should receive config files in /app/config/. The following file content:

```
{
    "name": "SimpleShop Stock API",
    "port": "8082"
}
```

**Deploy and verify the configuration**
Apply the updated YAML specs and verify that the configuration is working by testing the application.

**Deliverable**
- YAML files modeling ConfigMaps, Secrets, and updated Pods
- The kubectl commands used to deploy and test the app

## Part 4: Update the App to Run with High Availability and Scale

**Objective**
Prepare the SimpleShop app for production by making it scalable and highly available.

**Workflow**
- Migrate to Deployment objects
- Replace the Pod specs with Deployment specs for each component (e.g., simpleshop-products-db, simpleshop-products-api, etc.).

**Scale the application**
- Scale the simpleshop-products-api and simpleshop-stock-api Deployments to run 2 Pods each.
- Scale the web component to run 3 Pods and set an environment variable to print the server name in debug mode.

**Upgrade the web component**
Roll out a new version (stefanprifti/simpleshop-web:v2) and ensure the website stays online during the update.

**Test the scaling and update process**
Verify the app is load-balancing correctly and that the update is successful.

**Deliverable**
- YAML files modeling the Deployments and the scaling commands
- The kubectl commands used to deploy, scale, and upgrade the app
