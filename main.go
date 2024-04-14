package main

import (
	handlers "aggregation-service-cluster-api/cmd/api/handlers"
	watcher "aggregation-service-cluster-api/cmd/api/watcher"
	client "aggregation-service-cluster-api/cmd/client"
	"strconv"

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
	deploymentList, err := handlers.HandleListDeployments(genereated_client)
	if err != nil {
		panic(err)
	}

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

	// Create deployment and set required number of replicas
	// Be sure to set the name
	// e.g. localhost:8081/createDeployment/newdeploy/3 (name: newdeploy, replicas: 3)
	router.GET("/createdeployment/:deployment_name/:replicas", func(c *gin.Context) {
		deployment_name := c.Param("deployment_name")
		if deployment_name == "" {
			c.JSON(400, gin.H{
				"message": "Deployment name is required.",
			})
		}
		replicas := c.Param("replicas")
		replicas_int, err := strconv.Atoi(replicas)
		if err != nil {
			panic("Error with data type of mentioned replicas ")
		}
		handlers.CreateDeploy(genereated_client, deployment_name, replicas_int)
		c.JSON(200, gin.H{
			"message": "Deployment created successfully.",
		})
	})

	// Fetch YAML for a resource
	// e.g. localhost:8081/fetchYAML/pod/pod-name
	router.GET("/fetchyaml/:resource/:name", func(c *gin.Context) {
		resource := c.Param("resource")
		name := c.Param("name")
		if resource == "" || name == "" {
			c.JSON(400, gin.H{
				"message": "Resource and name are required.",
			})
			return
		}
		yaml, err := handlers.FetchYAML(genereated_client, "default", resource, name)
		if err != nil {
			panic(err)
		}
		c.JSON(200, gin.H{
			"yaml": yaml,
		})
	})

	// watcher
	// e.g. localhost:8081/watcher
	router.GET("/watcher", func(c *gin.Context) {
		watcher.Watcher(genereated_client)
		c.JSON(200, gin.H{
			"message": "Watcher started.",
		})
	})

	// Update deployment (change replicas) and change image of deployment
	// e.g. localhost:8081/deploymentscale/deployment-name/replicas/image
	router.GET("/deploymentupdate/:deploymentName/:replicas/:newImage", func(c *gin.Context) {
		deploymentName := c.Param("deploymentName")
		replicas := c.Param("replicas")
		replicas_int, err := strconv.Atoi(replicas)
		if err != nil {
			panic("err")
		}
		newImage := c.Param("newImage")
		handlers.UpdateDeployment(genereated_client, deploymentName, replicas_int, newImage)
		c.JSON(200, gin.H{
			"message": "Scaling deployment " + deploymentName + " to " + replicas + " replicas.",
		})
	})

	// Delete deployment
	// e.g. localhost:8081/deletedeployment/deployment-name
	router.GET("/deletedeployment/:deploymentName", func(c *gin.Context) {
		deploymentName := c.Param("deploymentName")
		handlers.DeleteDeploy(genereated_client, deploymentName)
		c.JSON(200, gin.H{
			"message": "Deployment " + deploymentName + " deleted successfully.",
		})
	})

	// scale deployment (not working)
	// router.POST("/deploymentscale", func(c *gin.Context) {
	// 	// get from json request body
	// 	type Data struct {
	// 		DeploymentName string `json:"deploymentName"`
	// 		Replicas       string `json:"replicas"`
	// 	}
	// 	var data Data
	// 	if err := c.BindJSON(&data); err != nil {
	// 		// DO SOMETHING WITH THE ERROR
	// 		c.JSON(400, gin.H{
	// 			"message": "Invalid request body",
	// 		})
	// 		return
	// 	}
	// 	deploymentName := data.DeploymentName
	// 	replicas := data.Replicas
	// 	// scale deployment
	// 	// handlers.PatchDeploymentObject(ctx, genereated_client, deploymentName, replicas)
	// 	c.JSON(200, gin.H{
	// 		"message": "Scaling deployment " + deploymentName + " to " + replicas + " replicas.",
	// 	})
	// })

	// fetch resources(pods, nodes) according to context name
	// e.g. localhost:8081/context/kind-cluster-1
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
