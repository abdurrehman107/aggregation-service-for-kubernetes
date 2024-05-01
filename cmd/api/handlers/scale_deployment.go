// here,
// "client" be the kubernetes client
// "cur" be the current deployment object
// "mod" be the modified deployment object,
// make sure to use deep copy before modifying the deployment object.

package handlers

import (
	"context"
	"encoding/json"
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/strategicpatch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/util/retry"
)

func PatchDeploymentObject(ctx context.Context, client kubernetes.Interface, cur, mod *appsv1.Deployment) (*appsv1.Deployment, error) {
	curJson, err := json.Marshal(cur)
	if err != nil {
		return nil, err
	}

	modJson, err := json.Marshal(mod)
	if err != nil {
		return nil, err
	}

	patch, err := strategicpatch.CreateTwoWayMergePatch(curJson, modJson, appsv1.Deployment{})
	if err != nil {
		return nil, err
	}

	if len(patch) == 0 || string(patch) == "{}" {
		return cur, nil
	}

	out, err := client.AppsV1().Deployments(cur.Namespace).Patch(ctx, cur.Name, types.StrategicMergePatchType, patch, metav1.PatchOptions{})
	return out, err
}

func UpdateDeployment(client *kubernetes.Clientset, deploymentName string, replicas int) {
	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		// Retrieve the latest version of Deployment before attempting update
		// RetryOnConflict uses exponential backoff to avoid exhausting the apiserver
		result, getErr := client.AppsV1().Deployments("default").Get(context.TODO(), deploymentName, metav1.GetOptions{})
		if getErr != nil {
			panic(fmt.Errorf("failed to get latest version of Deployment: %v", getErr))
		}

		result.Spec.Replicas = int32Ptr(int32(replicas))             // reduce replica count
		// if image != "" {
		// 	result.Spec.Template.Spec.Containers[0].Image = image // change nginx version
		// }
		_, updateErr := client.AppsV1().Deployments("default").Update(context.TODO(), result, metav1.UpdateOptions{})
		return updateErr
	})
	if retryErr != nil {
		panic(fmt.Errorf("update failed: %v", retryErr))
	}
}
func int32Ptr(i int32) *int32 { return &i }
