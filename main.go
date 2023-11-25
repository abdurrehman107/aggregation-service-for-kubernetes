package main

import (
	"context"
	"flag"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := flag.String("kubeconfig", "/Users/abdurrehman/.kube/config", "location for my kubeconfig file")
	fmt.Println("Kubeconfig: ", kubeconfig)
	// create config object to create Kubernetes clients down the line
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Println("Config error")
	}
	// create a Kubernetes client
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println("clientset error")
	}
	pods, err := clientset.CoreV1().Pods("default").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		fmt.Println("pods error")
	}
	fmt.Println("Pods: ")
	for _, pod := range pods.Items {
		fmt.Printf("%s", pod.Name)
	}
}
