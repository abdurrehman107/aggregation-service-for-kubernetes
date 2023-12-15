package handlers

import (
	client "aggregation-service-cluster-api/cmd/client"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// handleListNodes responds to GET /nodes requests.
func HandleListNodes(c *gin.Context) {
    nodes, err := client.Client().CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
    if err != nil {
        c.Error(err)
        return
    }

    // Convert node list to a JSON map
    var nodeList []map[string]interface{}
    for _, node := range nodes.Items {
        nodeMap := map[string]interface{}{
            "name":      node.Name,
            "namespace": node.Namespace,
            "status":    node.Status.Phase,
        }
        nodeList = append(nodeList, nodeMap)
    }

    c.JSON(http.StatusOK, map[string]interface{}{
        "nodes": nodeList,
    })
}

