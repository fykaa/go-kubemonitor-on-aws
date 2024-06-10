# Deploy on Kubernetes (EKS) Programmatically Using client-go Library

Now that we have our Image deployed on Docker, we can create an Amazon EKS cluster and deploy applications on it using the client-go library for Kubernetes. The client-go library allows you to interact with your Kubernetes cluster programmatically using Go.

## Step 1: Create an EKS Cluster in AWS

### 1.1 Create an EKS Kubernetes Cluster with Nodes
1. **Open EKS Console**
    - Navigate to the Amazon EKS console.

2. **Add a Cluster**
    - Click on `Add cluster` and select `Create cluster`.

3. **Cluster Configuration**
    - Provide a name for your cluster.
    - Leave the settings to defaults.

4. **Networking Configuration**
    - For this we will ensure four subnets are selected.
    - Create a security group with the appropriate ports opened.

5. **Logging and Add-ons**
    - Disable logging and add-ons for simplicity.

6. **Review and Create**
    - Review the settings.
    - Use the defaults and click `Create`.

   ![Cluster Creation](../assets/eksClusterCreated.png)

7. **Wait for Cluster Creation**
    - The creation process may take several minutes.

### 1.2 Create a Node Group
1. **Go to Compute**
    - After the cluster is created, navigate to `Compute`.

2. **Create Node Group**
    - Click on `Create node group`.
    - Provide a name for the node group.
    - Attach an IAM role suitable for EKS nodes.
    - Select an instance type, such as `t2.micro`.
    - Keep the desired size and subnets as defaults.

3. **Create Node Group**
    - Click `Create`.

   ![Node Group Creation](../assets/eksNodeGroup.png)

4. **Wait for Node Group Creation**
    - The node group should now have 2 nodes.

## Step 2: Create Kubernetes Deployments and Services Using client-go

### 2.1 Prerequisites
- Ensure you have `kubectl` CLI installed.
- Make sure your cluster is up and running for testing purposes.

### 2.2 Using client-go Library
- Typically, Kubernetes deployments and services are created using YAML files. However, in this tutorial, we will use the client-go library for programmatic creation.

### 2.3 Sample Go Code
- Create a Go file like `eks.go` which defines and creates a Kubernetes deployment and Service.

### 2.4 Running the Code
1. **Run the Go Program**
   ```bash
   go run eks.go
   ```

2. **Check Deployments and Services**
   ```bash
   kubectl get deployments
   kubectl get services
   ```

3. **Port-Forward to Access the Service**
   ```bash
   kubectl port-forward svc/<svc-name> 8080:8080
   ```

## Key Learning Points
- Understanding Kubernetes.
- Using the standard HTTP server provided by Kubernetes.
- Creating an EKS cluster and node groups.
- Programmatically creating Kubernetes deployments and services using Go.
- Utilizing the client-go library.

## IAM Roles Used
Here's a detailed list of all the IAM roles used in this project:
![IAM Roles](../assets/iamRoles.png)

## Cleanup
After completing the project, ensure to delete everything you created:
1. **Delete Node Groups in EKS**
2. **Delete EKS Cluster**
3. **Delete ECR Repository**

Ensuring proper cleanup will help maintain a tidy and cost-effective AWS environment.

## The End
This page provided a guide to creating an EKS cluster, deploying nodes, and using the client-go library to manage Kubernetes resources programmatically.