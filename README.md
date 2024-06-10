# go-kubemonitor-on-aws

## Overview

This project is a Cloud Native system resource monitoring application built with the Go Programming Language and deployed on Kubernetes (K8s) with AWS Elastic Container Registry (ECR).

## Architecture Diagram

(project flow diagram here)

## Setting Up Your Development Environment

1. **Install Go**: [Go installation guide](https://golang.org/doc/install)
2. **Install Docker**: [Docker installation guide](https://docs.docker.com/get-docker/)
3. **Install AWS CLI**: [AWS CLI installation guide](https://docs.aws.amazon.com/cli/latest/userguide/install-cliv2.html)
4. **Install kubectl**: [kubectl installation guide](https://kubernetes.io/docs/tasks/tools/install-kubectl/)
5. **Configure AWS CLI**: Run `aws configure`

You will need a code editor (e.g., VS Code).

## Building the Project

### Step 1: Build the Go Application

#### Tools and Packages Used
- Go
- gopsutil package
- HTML, JavaScript, jQuery, Plotly.js

1. **Create a Monitoring Application using Go**:
    - Fetch CPU and memory usage using the `gopsutil` package.
    - Serve the data as a `JSON API`.
    - Create a dynamic front-end using `Plotly.js` and `jQuery` for real-time updates.

#### Backend
- **Go**: Used with `gopsutil` to fetch CPU & memory usage.
- **JSON API**: Endpoint to serve system metrics as `JSON`.

#### Frontend
- **HTML & JavaScript**: For UI.
- **Plotly.js**: For dynamic graphs.
- **jQuery**: For AJAX calls.

**Code Files**:
- `main.go`: Contains the Go backend code.
- `index.gohtml`: Contains the HTML code for the front-end.

**Steps to Run the Code**:
1. Clone the repository and navigate to the root directory.
2. Run the Go application:
    ```sh
    go run main.go
    ```
3. Open a web browser and navigate to `http://localhost:8080` to see the UI.

### Step 2: Containerize the Application

#### Tools and Packages Used
- Docker

#### Key Learning Points
- Learning Docker and how to containerize a Go application.
- Creating a Dockerfile.
- Building a Docker image.
- Running a Docker container.
- Docker commands.

1. **Create Dockerfile**:
    - The Dockerfile is located in the root directory and named `Dockerfile`.
    - **What the Dockerfile does**:
        - **Stage 1**: Uses a Golang image to build the Go application. [For more information](https://hub.docker.com/_/golang).
        - **Stage 2**: Creates a minimal image containing only the necessary files to run the application.
    - **Optimisation**: We are using multi-staged builds to optimise by separating the builds from final runtime env, and reducing the size of the final image.


2. **Build Docker Image**:
    ```sh
    docker build -t kubemonitor .
    ```

3. **Run Docker Container Locally**:
    ```sh
    docker run -p 8080:8080 kubemonitor
    ```

## Step 3: Push to AWS ECR using Go

#### Tools and Packages Used
- AWS SDK for Go

#### Key Learning Points
- Creating an ECR repository using Go.
- Pushing a Docker image to ECR.

1. **Create ECR Repository using Go**:
    - Write Go code to interact with AWS SDK and create an ECR repository.
    - Refer to the code in `ecr.go` file in this repository.

## Step 4: Deploy on Kubernetes (EKS)

(steps and details on deploying the application on eks remaining)

### Key Learning Points
- Learning Kubernetes.
- Creating an EKS cluster and node groups.
- Creating Kubernetes deployments and services using Go.


## Tip for Efficient Learning
To gain deeper insights into the project's development, check out the [branches](https://github.com/fykaa/go-kubemonitor-on-aws/branches) and [PR commits](https://github.com/fykaa/go-kubemonitor-on-aws/pulls?q=is%3Apr+label%3Aoptimization+). Each pull request documents the iterative improvements and thought process behind optimizing the project. This approach will help you understand the evolution of the project from the initial version to the optimized version.
