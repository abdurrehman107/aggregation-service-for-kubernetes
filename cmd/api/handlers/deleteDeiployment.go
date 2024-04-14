package handlers

import (
	"context"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// delete deployment
func DeleteDeploy(client *kubernetes.Clientset, deploymentName string) {
	// generate client for deployment
	deployClient := client.AppsV1().Deployments("default")
	// delete deployment
	err := deployClient.Delete(context.TODO(), deploymentName, metaV1.DeleteOptions{})
	if err != nil {
		panic(err)
	}
}
