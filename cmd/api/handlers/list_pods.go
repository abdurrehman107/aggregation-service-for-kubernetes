package handlers

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)
func HandleListPods(client *kubernetes.Clientset) (interface{}, error) {
    pods, err := client.CoreV1().Pods("default").List(context.Background(), metav1.ListOptions{})
    if err != nil {
        return nil, err
    }

    // Convert pod list to a JSON map
    var podList []map[string]interface{}
    for _, pod := range pods.Items {
        podMap := map[string]interface{}{
            "name":       pod.Name,
            "namespace":  pod.Namespace,
            "status":     pod.Status.Phase,
        }
        podList = append(podList, podMap)
    }
    return podList, nil
    // response := map[string]interface{}{
    //     "nodes": podList,
    // }
    // return response, nil
}
