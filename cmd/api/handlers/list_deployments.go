package handlers

import (
    "context"
    // "encoding/json"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/client-go/kubernetes"
)

func HandleListDeployments(client *kubernetes.Clientset) (interface{}, error) {
    deployments, err := client.AppsV1().Deployments("default").List(context.Background(), metav1.ListOptions{})
    if err != nil {
        return nil, err
    }

    // Convert deployment list to JSON map
    var deploymentList []map[string]interface{}
    for _, deployment := range deployments.Items {
        deploymentMap := map[string]interface{}{
            "name":      deployment.Name,
            "namespace": deployment.Namespace,
            "replicas":  *deployment.Spec.Replicas,
            // Add more relevant deployment information as needed
        }
        deploymentList = append(deploymentList, deploymentMap)
    }

    return deploymentList, nil
}
