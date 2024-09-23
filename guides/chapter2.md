# Chapter 2: Enhancing Reliability, Security, and Storage for SimpleShop with Kubernetes

## Preface

This guide outlines the steps to enhance the SimpleShop application's reliability, security, and storage management using Kubernetes. The primary objectives include implementing self-healing mechanisms through liveness and readiness probes, enhancing container security to prevent breaches, and configuring persistent and ephemeral storage for efficient log management and data retention.

## Using Probes to make SimpleShop self-healing
To make the SimpleShop app self-healing using Kubernetes probes, we'll add liveness and readiness probes for each service:

1. Products API:
- Add a `/health` endpoint check.
- Liveness probe restarts pods if they fail.
- Readiness probe ensures traffic is routed only to healthy pods.

2. Stock API:
- Add a `/health` endpoint probe.
- Similar to Products API but with shorter probe intervals.

3. Web
- Use the `/` endpoint.
- Probes run aggressively (every 5 seconds) to ensure high availability.

4. Products DB:
- Use the `pg_isready` command to check the database health.


**Use Named Ports to Reduce Duplication**
To avoid referencing port numbers in multiple places, define named ports in the container spec and use them across the deployment and service YAML configurations.

```
ports:
  - name: http
    containerPort: 8080
  - name: db
    containerPort: 5432
```

This allows easier maintenance when port numbers change, as updates need to be done only in one place (the container definition), and other configurations (e.g., services, probes) can reference the named port directly.

**Deliverable**
- YAML files modeling the deployments.
- The kubectl commands used to deploy and test the app

## Enhanced Security for SimpleShop Containers
To secure the SimpleShop app and prevent security breaches, we’ll implement several security measures for the products-api, stock-api, web, and products-db services.

All services use Go, and each service has tailored security measures applied to ensure they are secure without breaking functionality.

Workflow and Security Requirements:

**1. Block Kubernetes API Access:**
All services have a token loaded by default, but they don’t need Kubernetes API access. Disable token auto-mounting to remove unnecessary access.

Hint:
```
automountServiceAccountToken: false
```

**2. Products API:**
- Concern: CPU/memory overuse under load.
- Solution: Limit CPU and memory resources.

Hint:
```
resources:
  requests:
    memory: "256Mi"
    cpu: "500m"
  limits:
    memory: "500Mi"
    cpu: "1"
```

**3. Stock API:**
- Concern: Compromise leads to file system changes.
- Solution: Make the file system read-only.

Hint
```
readOnlyRootFilesystem: true 
```

**4. Web:**
- Concern: Public-facing, prone to attack.
- Solution: Run with non-root user, drop Linux capabilities, and prevent privilege escalation.

Hint
```
securityContext:
    runAsUser: 1000  # Use non-root user
    runAsGroup: 1000  # Matching group ID
    runAsNonRoot: true  # Ensure the container runs as non-root
    readOnlyRootFilesystem: true  # Ensure the filesystem is read-only
    capabilities:
    drop:
        - ALL  # Drop all unnecessary Linux capabilities
```

**5. Products DB:**

Note: Do not modify except for disabling Kubernetes API token access.

**Deliverable**
- YAML files modeling the deployments.


## Using volumes for storage
In this milestone, we will configure Kubernetes storage volumes to meet the data and logging requirements for SimpleShop. Specifically, we will surface application logs for the Product API service and configure persistent storage for the Products Database.

**Objective**
Kubernetes offers robust storage management, allowing data to persist across pod restarts and enabling log and file sharing between containers. In this guide, we will:

- Surface application logs for the Product API using ephemeral storage.
- Ensure Products Database data persists using persistent storage.

**Workflow**

1. Surface Application Logs from the Product API Pods
The Product API logs to a file inside the container, but these logs aren't visible to Kubernetes. To surface the logs:

- Step 1: Create a shared volume in the pod to store logs at /logs/app.log.
- Step 2: Add a secondary container in the pod that reads and prints the logs using tail -f.
- Step 3: Include an init container to create an empty log file before the main container starts, preventing errors due to missing logs.

Testing
- Check that logs are written to `/logs/app.log` and are viewable through the logger container.

2. Persistent Storage for Products Database

The Products Database needs persistent storage to ensure that product data is retained across pod restarts. We'll use a PersistentVolumeClaim (PVC) for this.

Testing
- Test persistence by modifying data in the database and then restarting the pod. Verify that data is retained after restart.

**Deliverable**
- YAML files modeling the deployments and volumes.
- The kubectl commands used to deploy and test the app.
