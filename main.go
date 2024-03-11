package main

import (
	handlers "aggregation-service-cluster-api/cmd/api/handlers"
	client "aggregation-service-cluster-api/cmd/client"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// create config object
	config := client.GenerateDefaultConfig()
	// create a client
	genereated_client := client.Client(config)
	// set nodes to handleListNodes function in cmd/api/handlers/list_nodes.go
	nodes, err := handlers.HandleListNodes(genereated_client)
	if err != nil {
		panic(err)
	}
	pods, err := handlers.HandleListPods(genereated_client)
	if err != nil {
		panic(err)
	}
	// deploymentClient, err := handlers.DeploymentClient(genereated_client)
	if err != nil {
		panic(err)
	}
	deploymentList, err := handlers.HandleListDeployments(genereated_client)
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

	router.GET("/deployments", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"deployments": deploymentList,
			"message":     "Do you wish to scale?.",
		})
	})

	router.POST("/deploymentscale", func(c *gin.Context) {

		// get from json request body
		type Data struct {
			DeploymentName string `json:"deploymentName"`
			Replicas       string `json:"replicas"`
		}

		var data Data

		if err := c.BindJSON(&data); err != nil {
			// DO SOMETHING WITH THE ERROR
			c.JSON(400, gin.H{
				"message": "Invalid request body",
			})
			return
		}

		deploymentName := data.DeploymentName
		replicas := data.Replicas

			// scale deployment
		// handlers.PatchDeploymentObject(ctx, genereated_client, deploymentName, replicas)	

		c.JSON(200, gin.H{
			"message": "Scaling deployment " + deploymentName + " to " + replicas + " replicas.",
		})
	})

	// fetch resources(pods, nodes) according to context name
	router.GET("/context/:context", func(c *gin.Context) {
		context := c.Param("context")
		config, err := client.BuildConfigWithContextFromFlags(context, "/Users/abdurrehman/.kube/config")
		if err != nil {
			panic(err)
		}
		genereated_client := client.Client(config)
		nodes, err := handlers.HandleListNodes(genereated_client)
		if err != nil {
			panic(err)
		}
		pods, err := handlers.HandleListPods(genereated_client)
		if err != nil {
			panic(err)
		}
		c.JSON(200, gin.H{
			"nodes": nodes,
			"pods":  pods,
		})
	})

	router.Run("localhost:8081")
}
