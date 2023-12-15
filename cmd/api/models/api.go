package models

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type resource struct {
    CLUSTER     string  `json:"cluster"`
    NODE  string  `json:"node"`
    POD string  `json:"pod"`
}
// resources to pull | sample resources
var resources = []resource{
    {CLUSTER: "1", NODE: "n1", POD: "pod1"},
    {CLUSTER: "2", NODE: "n2", POD: "pod2"},
    {CLUSTER: "3", NODE: "n3", POD: "pod3"},
}
func run() {
    router := gin.Default()
    router.GET("/albums", getResource)

    router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
func getResource(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, resources)
}