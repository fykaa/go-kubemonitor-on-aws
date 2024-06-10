package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
)

func main() {
	// Loading the default AWS configuration
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	// Creating an ECR service client
	client := ecr.NewFromConfig(cfg)

	// Defining the repository name
	repositoryName := "kubemonitor"

	// Calling CreateRepository
	params := &ecr.CreateRepositoryInput{
		RepositoryName: aws.String(repositoryName),
	}
	resp, err := client.CreateRepository(context.TODO(), params)
	if err != nil {
		log.Fatalf("failed to create repository, %v", err)
	}

	// Retrieving and printing the repository URI
	repositoryURI := aws.ToString(resp.Repository.RepositoryUri)
	fmt.Println(repositoryURI)
}
