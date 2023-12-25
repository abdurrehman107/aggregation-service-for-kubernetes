package main

import (
	handlers "aggregation-service-cluster-api/cmd/api/handlers"
	client "aggregation-service-cluster-api/cmd/client"
	// "context"

	"github.com/gin-gonic/gin"
	// metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {
	router := gin.Default()
	// create a client
	client := client.Client()
	// set nodes to handleListNodes function in cmd/api/handlers/list_nodes.go
	nodes, err := handlers.HandleListNodes(client)
	if err != nil {
		panic(err)
	}
	pods, err := handlers.HandleListPods(client)
	if err != nil {
		panic(err)
	}
	// podWatcher, err := client.CoreV1().Pods("").Watch(context.TODO(), metav1.ListOptions{})
	// go func() {
	// 	for {
	// 		select {
	// 		case event := <-podWatcher.ResultChan():
	// 			// Handle pod event
	// 			// Update the pods variable or trigger a refresh
	// 		case event := <-nodeWatcher.ResultChan():
	// 			// Handle node event (if needed)
	// 		}
	// 	}
	// }()
	// if err != nil {	
	// 	panic(err)
	// }
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hey there! Go on to /nodes or /pods to get the list of nodes and pods respectively.",
		})
	})
	router.GET("/nodes", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"nodes": nodes,
		})
	})
	router.GET("/pods", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"pods": pods,
		})
	})
	router.Run("localhost:8081")
}
