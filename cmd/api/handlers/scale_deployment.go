// here,
// "client" be the kubernetes client
// "cur" be the current deployment object
// "mod" be the modified deployment object,
// make sure to use deep copy before modifying the deployment object.

package handlers

import (
	"context"
	"encoding/json"

	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/strategicpatch"
	"k8s.io/client-go/kubernetes"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func PatchDeploymentObject(ctx context.Context, client kubernetes.Interface, cur, mod *v1.Deployment) (*v1.Deployment, error) {
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