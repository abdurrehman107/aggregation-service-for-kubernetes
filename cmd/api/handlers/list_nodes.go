package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	client "aggregation-service-cluster-api/cmd/client"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)
func handleListNodes(w http.ResponseWriter, r *http.Request) {
    nodes, err := client.Client().CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Encode pods into JSON
    data, err := json.Marshal(nodes.Items)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Write JSON data to response
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(data)
}
