package main

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"

	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func int32Ptr(i int32) *int32 { return &i }

func main() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// Building the Kubernetes client configuration
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// Creating a Kubernetes client
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// Defining the deployment
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "go-http-app",
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "go-http-app",
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "go-http-app",
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  "go-http-container",
							Image: "474802787127.dkr.ecr.us-west-2.amazonaws.com/kubemonitor",
							Ports: []apiv1.ContainerPort{
								{
									ContainerPort: 8080,
								},
							},
						},
					},
				},
			},
		},
	}

	// Creating the deployment
	deploymentsClient := clientSet.AppsV1().Deployments(apiv1.NamespaceDefault)
	fmt.Println("Creating deployment...")
	_, err = deploymentsClient.Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Deployment created.")

	// Define the service
	service := &apiv1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "go-http-service",
		},
		Spec: apiv1.ServiceSpec{
			Selector: map[string]string{
				"app": "go-http-app",
			},
			Ports: []apiv1.ServicePort{
				{
					Port: 8080,
				},
			},
		},
	}

	// Creating the service
	servicesClient := clientSet.CoreV1().Services(apiv1.NamespaceDefault)
	fmt.Println("Creating service...")
	_, err = servicesClient.Create(context.TODO(), service, metav1.CreateOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Service created.")
}
