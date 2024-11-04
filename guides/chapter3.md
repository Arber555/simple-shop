# Automated Deployments

## Templated Deployment Using Helm

**Objective**

Package your entire application—including the web frontend, databases, and backend APIs—into a single Helm chart. This chart will act as a deployable artifact that can be used across all environments, enabling easy application of environment-specific configurations.

---

**Why It Matters**

- **Consistency**: Maintain the same code version across all environments, with configuration adjustments specific to each environment.
- **Simplification**: Easily manage updates to image versions and settings using Helm variables.
- **Scalability**: Deploy multiple instances of the application in the same cluster without naming conflicts.
- **Automation**: Build the foundation for automated deployments as part of your CI/CD pipeline.

---

**Steps**

1. **Deploy the Current App**

   - Use your existing Kubernetes YAML files (from Chapter 2) to deploy the application and verify accessibility.
   - Install [Helm](https://helm.sh/docs/intro/install/) if you haven't already.
   - Delete the existing application deployed in Kubernetes.

2. **Create a Helm Chart**

   - Organize your YAML files into the Helm chart structure: `Chart.yaml`, `values.yaml`, and a `templates/` directory containing your Kubernetes manifests. Here's the recommended structure:
     ```
     ├── Chart.yaml
     ├── templates
     │   ├── simpleshop-products-api.yaml
     │   ├── simpleshop-products-db.yaml
     │   ├── simpleshop-stock-api.yaml
     │   └── simpleshop-web.yaml
     └── values.yaml
     ```
   - Populate `Chart.yaml` with metadata about your chart.
   - Deploy the app using `helm install` and ensure everything functions correctly. Refer to the Helm documentation for exact command usage.

3. **Template Resource Names**

   - Modify resource names to include `{{ .Release.Name }}` to avoid naming conflicts.
   - Update `metadata.name`, `labels`, and `selectors` in your templates to use the release name.
   - Adjust internal service URLs in ConfigMaps and Secrets to match these new names.

4. **Make Configurations Optional**

   - Introduce variables to toggle features like:
     - `includeReadiness` (default: `true`)
     - `includeLiveness` (default: `true`)
   - Use Helm’s conditional statements to include or exclude probes based on these variables.

5. **Parameterize Images and Replicas**

   - Use variables for image names, replica counts, service configurations, and resource settings. For example, the Products API could have:
     ```yaml
     includeLiveness: true
     includeReadiness: true

     productsConnectionString: "your-database-dsn"
     productsApiImage: "stefanprifti/simpleshop-products-api:v1"
     productsApiReplicaCount: 1
     productsApiResources:
       requests:
         memory: "256Mi"
         cpu: "200m"
       limits:
         memory: "512Mi"
         cpu: "500m"
     ```
   - Follow a similar approach for the other templates.
   - Deploy the Helm chart and verify that the application still functions as expected.

---

**Deliverables**

- **Helm Chart Folder**: A folder containing all templated YAML files and a `values.yaml` file with default settings.
- **Deployment Commands**: A list of `helm` commands used for deployment. 


## Deployment Strategies

**Objective**

Enable on-demand releases for the entire SimpleShop application, with the ability to roll back immediately if issues arise. Ensure each component has a tailored rollout strategy to prevent disruptions and maintain availability.

---

**Why It Matters**

- **Custom Rollouts**: Default Kubernetes strategies may not suit all components. Adjusting strategies ensures smooth deployments.
- **Coordinated Updates**: Helm allows for unified, fail-safe deployments, minimizing the risk of partial updates.
- **Rollback Control**: Helm makes rollbacks straightforward, offering greater control over failed deployments compared to Kubernetes alone.

---

### Steps

1. **Deploy the Current SimpleShop App with Helm**

   - Use your existing Helm chart as a base. Docker images use the tag `:v1` for all components.
   - Deploy and confirm everything works.

2. **Configure Fast Rollouts for the Stock API**

   - **Requirement**: Stock API can handle multiple versions simultaneously and higher Pod counts.
   - **Setup**: Configure an update strategy that doubles the number of running Pods during updates. Ensure old Pods aren't removed until new ones are ready.
   - **Make Configurable**: Allow these settings to be adjustable in the Helm chart.

   ```
    strategy:
    type: ...
    rollingUpdate:
        maxSurge:       ... # Double the Pods during the update
        maxUnavailable: ... # Ensure all existing Pods remain available
    ```

3. **Set Up Replacement Upgrades for the Products API**

   - **Requirement**: Products API should not run multiple versions simultaneously.
   - **Setup**: Configure a strategy to replace existing Pods only after termination. Make this behavior configurable.

   ```
    strategy:
    type: ...  # Terminate old Pods before creating new ones
    ```

4. **Implement Partial Updates for the Products Database**

   - **Requirement**: Update secondaries first and pause before updating the primary.
   - **Setup**: Use Helm to pause the update at a specific Pod number, defaulting to leave the primary untouched unless specified.

   ```
    updateStrategy:
    type: ...
    rollingUpdate:
        partition: ...  # Update secondaries first, pause before updating the primary
    ```

5. **Enable Blue/Green Deployments for the Web**

   - **Requirement**: Run both blue and green versions simultaneously and switch traffic easily.
   - **Setup**: Deploy the v2 as the blue release and the v1 as the green release. Direct traffic to one version using the Service.
   - **Customizable**: Make the image versions configurable in the Helm chart. Make the active release configurable.  Verify seamless traffic switching between blue and green releases with no downtime.

---

## Deliverables

- **Helm Chart**: Complete with templated YAML files and configurable rollout strategies.
- **Testing Evidence**: Demonstrate that updates and rollbacks work as intended, including blue/green deployment switches without downtime.
