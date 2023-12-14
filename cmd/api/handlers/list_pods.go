package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	client "aggregation-service-cluster-api/cmd/client"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)
func handleListPods(w http.ResponseWriter, r *http.Request) {
    pods, err := client.Client().CoreV1().Pods("default").List(context.Background(), metav1.ListOptions{})
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Encode pods into JSON
    data, err := json.Marshal(pods.Items)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Write JSON data to response
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(data)
}
