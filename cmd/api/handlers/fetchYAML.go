package handlers

import (
	"bytes"
	"context"
	"fmt"

	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/serializer/json"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/util/retry"
)

func FetchYAML(client *kubernetes.Clientset, namespace string, resourceType string, resourceName string) {
	// fetch yaml of resource
	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		// Retrieve the latest version of Deployment before attempting update
		// RetryOnConflict uses exponential backoff to avoid exhausting the apiserver
		deploymentsClient := client.AppsV1()
		result, getErr := deploymentsClient.Deployments("default").Get(context.TODO(), "demo-deployment", metaV1.GetOptions{})
		if getErr != nil {
			panic(fmt.Errorf("failed to get latest version of Deployment: %v", getErr))
		}
		// Convert Deployment object to YAML
		jsonSerializer := json.NewSerializerWithOptions(json.DefaultMetaFactory, nil, nil, json.SerializerOptions{Yaml: true})
		var yamlBuffer bytes.Buffer
		if err := jsonSerializer.Encode(result, &yamlBuffer); err != nil {
			panic(err.Error())
		}
		fmt.Println(yamlBuffer.String())
		// result.Spec.Replicas = int32Ptr(1)                           // reduce replica count
		// result.Spec.Template.Spec.Containers[0].Image = "nginx:1.13" // change nginx version
		//_, updateErr := deploymentsClient.Update(context.TODO(), result, metav1.UpdateOptions{})
		// return updateErr
		return nil
	})
	if retryErr != nil {
		panic(fmt.Errorf("Update failed: %v", retryErr))
	}
	// generate client for resource
	// fetch resource
	// print yaml
}

func int32Ptr(i int) {
	panic("unimplemented")
}

/*
retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
	// Retrieve the latest version of Deployment before attempting update
	// RetryOnConflict uses exponential backoff to avoid exhausting the apiserver
	result, getErr := deploymentsClient.Get(context.TODO(), "demo-deployment", metav1.GetOptions{})
	if getErr != nil {
		panic(fmt.Errorf("Failed to get latest version of Deployment: %v", getErr))
	}

	result.Spec.Replicas = int32Ptr(1)                           // reduce replica count
	result.Spec.Template.Spec.Containers[0].Image = "nginx:1.13" // change nginx version
	_, updateErr := deploymentsClient.Update(context.TODO(), result, metav1.UpdateOptions{})
	return updateErr
})
if retryErr != nil {
	panic(fmt.Errorf("Update failed: %v", retryErr))
}

*/
