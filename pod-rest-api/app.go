package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func getPodsJSON() (map[string]interface{}, error) {
	cmd := exec.Command("kubectl", "get", "pods", "-o", "json")
	stdout, err := cmd.Output()
	if err != nil {
		log.Printf("Error running kubectl command: %v", err)
		return nil, err
	}
	var podsJSON map[string]interface{}
	err = json.Unmarshal(stdout, &podsJSON)
	if err != nil {
		log.Printf("Error unmarshaling JSON: %v", err)
		return nil, err
	}
	return podsJSON, nil
}

func getPods(c *gin.Context) {
	podsJSON, err := getPodsJSON()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting pods in JSON format"})
		return
	}
	pods, ok := podsJSON["items"].([]interface{})
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting pods in JSON format"})
		return
	}
	c.JSON(http.StatusOK, pods)
}

func main() {
	r := gin.Default()
	r.GET("/api/v1/pods", getPods)
	err := r.Run(":5000")
	if err != nil {
		fmt.Printf("Failed to start server: %v", err)
	}
}
