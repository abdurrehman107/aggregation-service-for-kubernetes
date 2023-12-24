package handlers

import (
	client "aggregation-service-cluster-api/cmd/client"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)
func HandleListPods(c *gin.Context) {
    pods, err := client.Client().CoreV1().Pods("default").List(context.Background(), metav1.ListOptions{})
    if err != nil {
        c.Error(err)
        return
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

    c.JSON(http.StatusOK, map[string]interface{}{
        "nodes": podList,
    })
}
