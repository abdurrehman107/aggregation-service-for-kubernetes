package handlers_test

import (
	handlers "aggregation-service-cluster-api/cmd/api/handlers"
	"encoding/json"
	"reflect"
	"testing"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/fake"
)

func TestHandleListNodesSuccess(t *testing.T) {
    // Create mock nodes

fakeNodes := []runtime.Object{
	&v1.Node{ObjectMeta: metav1.ObjectMeta{Name: "node1"}},
	&v1.Node{ObjectMeta: metav1.ObjectMeta{Name: "node2"}},
}

// Create mock client
fakeClient := fake.NewSimpleClientset(fakeNodes...)

// Convert fakeClient to *kubernetes.Clientset
client := kubernetes.NewForConfigOrDie(fakeClient.Config)

// Call the function
response, err := handlers.HandleListNodes(client)
if err != nil {
	t.Errorf("Unexpected error: %v", err)
}

// Convert response to JSON for comparison
jsonData, err := json.Marshal(response)
    if err != nil {
        t.Errorf("Error marshalling response: %v", err)
    }

    // Expected JSON response (adjust based on your desired output)
    expected := `[{"name":"node1","namespace":""},{"name":"node2","namespace":""}]`

    // Compare actual and expected JSON
    if !reflect.DeepEqual(string(jsonData), expected) {
        t.Errorf("Unexpected response: got %s, expected %s", string(jsonData), expected)
    }
}

func TestHandleListNodesError(t *testing.T) {
    // Create mock client with error
    fakeClient := fake.NewSimpleClientset(&v1.Status{Status: metav1.Status{Reason: metav1.StatusReasonForbidden}})

    // Call the function
    _, err := handlers.HandleListNodes(fakeClient)
    if err == nil {
        t.Errorf("Expected error, but got none")
    }

    // Verify expected error message
    expectedErrorMessage := "forbidden"
    if err.Error() != expectedErrorMessage {
        t.Errorf("Unexpected error message: got %v, expected %v", err.Error(), expectedErrorMessage)
    }
}

func TestHandleListNodesEmptyList(t *testing.T) {
    // Create mock client with empty list
    fakeClient := fake.NewSimpleClientset()

    // Call the function
    response, err := handlers.HandleListNodes(fakeClient)
    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }

    // Expected empty list
    expected := "[]"

    // Convert response to JSON for comparison
    jsonData, err := json.Marshal(response)
    if err != nil {
        t.Errorf("Error marshalling response: %v", err)
    }

    // Compare actual and expected JSON
    if !reflect.DeepEqual(string(jsonData), expected) {
        t.Errorf("Unexpected response: got %s, expected %s", string(jsonData), expected)
    }
}
