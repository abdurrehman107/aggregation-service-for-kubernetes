package handlers

import (
	"bytes"
	"context"
	"fmt"

	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/serializer/json"
	"k8s.io/client-go/kubernetes"
	// "k8s.io/client-go/util/retry"
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

	// var objMap map[string]interface{}
    // if err := json.Unmarshal(yamlBuffer.Bytes(), &objMap); err != nil {
    //     return "", err
    // }

	// fmt.Println(yamlBuffer.String())
	// yaml = yamlBuffer.String()
		// result.Spec.Replicas = int32Ptr(1)                           // reduce replica count
		// result.Spec.Template.Spec.Containers[0].Image = "nginx:1.13" // change nginx version
		//_, updateErr := deploymentsClient.Update(context.TODO(), result, metav1.UpdateOptions{})
		// return updateErr
		// return nil
	// })
	// if retryErr != nil {
	// 	panic(fmt.Errorf("update failed: %v", retryErr))
	// }
	// generate client for resource
	// fetch resource
	// print yaml
}