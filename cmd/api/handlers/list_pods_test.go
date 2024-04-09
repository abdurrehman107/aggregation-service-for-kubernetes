package handlers_test

import (
    // "context"
    "encoding/json"
    // "errors"
    "reflect"
    "testing"

    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/api/core/v1"
    "k8s.io/apimachinery/pkg/runtime"
    "k8s.io/client-go/kubernetes/fake"
	handlers "aggregation-service-cluster-api/cmd/api/handlers"

)

func TestHandleListPodsSuccess(t *testing.T) {
    // Create mock pods
    fakePods := []runtime.Object{
        &v1.Pod{ObjectMeta: metav1.ObjectMeta{Name: "pod1", Namespace: "default"}},
        &v1.Pod{ObjectMeta: metav1.ObjectMeta{Name: "pod2", Namespace: "default"}},
    }

    // Create mock client
    fakeClient := fake.NewSimpleClientset(fakePods...)

    // Call the function
    response, err := handlers.HandleListPods(fakeClient)
    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }

    // Convert response to JSON for comparison
    jsonData, err := json.Marshal(response)
    if err != nil {
        t.Errorf("Error marshalling response: %v", err)
    }

    // Expected JSON response (adjust based on your desired output)
    expected := `[{"name":"pod1","namespace":"default","status":""},{"name":"pod2","namespace":"default","status":""}]`

    // Compare actual and expected JSON
    if !reflect.DeepEqual(string(jsonData), expected) {
        t.Errorf("Unexpected response: got %s, expected %s", string(jsonData), expected)
    }
}

func TestHandleListPodsError(t *testing.T) {
    // Create mock client with error
    fakeClient := fake.NewSimpleClientset(&v1.Status{Status: metav1.Status{Reason: metav1.StatusReasonForbidden}})

    // Call the function
    _, err := handlers.HandleListPods(fakeClient)
    if err == nil {
        t.Errorf("Expected error, but got none")
    }

    // Verify expected error message
    expectedErrorMessage := "forbidden"
    if err.Error() != expectedErrorMessage {
        t.Errorf("Unexpected error message: got %v, expected %v", err.Error(), expectedErrorMessage)
    }
}

func TestHandleListPodsEmptyList(t *testing.T) {
    // Create mock client with empty list
    fakeClient := fake.NewSimpleClientset()

    // Call the function
    response, err := handlers.HandleListPods(fakeClient)
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
