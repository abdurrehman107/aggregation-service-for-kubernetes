package handlers

import (
	// client "aggregation-service-cluster-api/cmd/client"
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// handleListNodes responds to GET /nodes requests.
func HandleListNodes(client *kubernetes.Clientset) (interface{}, error) {
	nodes, err := client.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
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
	return nodeList, nil
}
