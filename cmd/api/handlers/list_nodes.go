package handlers

import (
	client "aggregation-service-cluster-api/cmd/client"
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
)

func handleListNodes(w http.ResponseWriter, r *http.Request) (map[string]interface{}, error) {
	nodes, err := client.Client().CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, err
	}

	// Convert pod list to a JSON map
	var nodeList []map[string]interface{}
	for _, node := range nodes.Items {
		nodeMap := map[string]interface{}{
			"name":      node.Name,
			"namespace": node.Namespace,
			"status":    node.Status.Phase,
		}
		nodeList = append(nodeList, nodeMap)
	}

	return map[string]interface{}{
		"pods": nodeList,
	}, nil
}
