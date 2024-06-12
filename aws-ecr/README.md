# Create an Elastic Container Registry Programmatically using Go

## Overview

This page walks you through creating an AWS Elastic Container Registry (ECR) programmatically using Go. You will learn to set up your development environment, write Go code to create an ECR repository, and push Docker images to it.

## Architecture Diagram

*Project flow diagram here*

## Setting Up Your Development Environment

### Prerequisites
- AWS CLI
- Programmatic access
- Access key matching

### Steps

1. **Set up AWS CLI**  
   Follow the instructions in the [AWS CLI getting started guide](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-quickstart.html).

2. **Install the AWS SDK for Go V2**  
   If you have a Go Module project, you can retrieve the SDK and its dependencies using the `go get` command. These dependencies will be recorded in the `go.mod` file. Run the following commands to retrieve the standard set of SDK modules:
   ```sh
   go get github.com/aws/aws-sdk-go-v2
   go get github.com/aws/aws-sdk-go-v2/config
   ```
   Next, install the required AWS service API clients. For this project, to retrieve the Amazon ECR API client, run:
   ```sh
   go get github.com/aws/aws-sdk-go-v2/service/ecr
   ```

3. **Get your AWS access keys**  
   Access keys consist of an access key ID and secret access key, used to sign programmatic requests to AWS. You can create these keys using the AWS Management Console. It's recommended to use IAM access keys instead of AWS root account access keys.

   **To get your access key ID and secret access key:**
    - Open the [IAM console](https://aws.amazon.com/console/.
    - On the navigation menu, choose Users.
    - Choose your IAM user name (not the check box).
    - Open the Security credentials tab, and then choose Create access key.
    - To see the new access key, choose Show. Your credentials will resemble:
        - Access key ID: `AKIAIOSFODNN7EXAMPLE`
        - Secret access key: `wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY`
    - Download the key pair as a `.csv` file and store it securely.
   
   **Things to Ensure**
   - Use IAM user with the correct policy attached. The required policy for this project is shared in `/resources`.
   - Define a permissions boundary.

For specifying credentials programmatically, follow the instructions at [Specifying Credentials Programmatically](https://aws.github.io/aws-sdk-go-v2/docs/configuring-sdk/#specify-credentials-programmatically).

## Tools and Packages Used
- [AWS SDK for Go](https://aws.github.io/aws-sdk-go-v2/)

## Key Learning Points
- Creating an ECR repository using Go.
- Pushing a Docker image to ECR.

## Building the Project

### Step 1: Writing Go Code

1. **Create ECR Repository using Go**:
    - Write Go code to interact with AWS SDK and create an ECR repository.
    - Refer to the code in the `ecr.go` file in this repository.
    - Run `go get` for the following packages:
      ```go
      go get github.com/aws/aws-sdk-go-v2/aws
      go get github.com/aws/aws-sdk-go-v2/config
      go get github.com/aws/aws-sdk-go-v2/service/ecr
      ```

### Step 2: Push to AWS ECR using Go

1. Use the code file `ecr.go` to create an ECR repository programmatically in AWS by running `go run ecr.go`.
2. Follow the link displayed on the screen.

### Step 3: Deploy Images to This ECR Repo Using Commands

1. Visit the ECR repositories management console.
2. Select the `kubemonitor` repository created programmatically.
3. Click on the "View push commands" button.
4. Copy and paste the push commands into your terminal to push your project's Docker image to ECR.
