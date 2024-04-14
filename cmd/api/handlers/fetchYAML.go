package handlers

import (
	"bytes"
	"context"
	"fmt"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/serializer/json"
	"k8s.io/client-go/kubernetes"
)

func FetchYAML(client *kubernetes.Clientset, namespace string, resourceType string, resourceName string) (string, error) {
	// fetch yaml of resource
	// var yaml string
	// retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
	// 	// Retrieve the latest version of Deployment before attempting update
	// 	// RetryOnConflict uses exponential backoff to avoid exhausting the apiserver
	deploymentsClient := client.AppsV1()
	object, err := deploymentsClient.Deployments("default").Get(context.TODO(), resourceName, metaV1.GetOptions{})
	if err != nil {
		panic(fmt.Errorf("failed to get Deployment: %v", err))
	}
	// Convert Deployment object to YAML
	jsonSerializer := json.NewSerializerWithOptions(json.DefaultMetaFactory, nil, nil, json.SerializerOptions{Yaml: true})
	var yamlBuffer bytes.Buffer

	// write the object into the buffer using the json serializer
	if err := jsonSerializer.Encode(object, &yamlBuffer); err != nil {
		panic(err.Error())
	}
	return yamlBuffer.String(), nil
}