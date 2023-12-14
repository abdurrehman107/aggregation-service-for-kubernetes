package handlers

import (
	"context"
	"net/http"
	client "aggregation-service-cluster-api/cmd/client"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)
func handleListPods(w http.ResponseWriter, r *http.Request) (map[string]interface{}, error) {
    pods, err := client.Client().CoreV1().Pods("default").List(context.Background(), metav1.ListOptions{})
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

    return map[string]interface{}{
        "pods": podList,
    }, nil
}
