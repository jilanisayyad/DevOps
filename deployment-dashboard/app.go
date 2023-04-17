package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jilanisayyad/DevOps/deployment-dashboard/handlers"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.LoadHTMLGlob("index.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/api/v1/deployments/namespace/:namespace", func(ctx *gin.Context) {
		namespace := ctx.Param("namespace")
		if namespace != "" {
			handlers.GetDeployments(ctx, namespace)
		}
	})

	r.GET("/api/v1/deployments/", func(ctx *gin.Context) {
		handlers.GetDeployments(ctx, "")
	})

	r.GET("/api/v1/namespaces/", func(ctx *gin.Context) {
		handlers.GetNamespaces(ctx)
	})

	r.GET("/api/v1/pods/namespace/:namespace", func(ctx *gin.Context) {
		namespace := ctx.Param("namespace")
		if namespace != "" {
			handlers.GetPods(ctx, namespace)
		}
	})

	r.GET("/api/v1/pods/", func(ctx *gin.Context) {
		handlers.GetPods(ctx, "")
	})

	r.GET("/api/v1/replicasets/namespace/:namespace", func(ctx *gin.Context) {
		namespace := ctx.Param("namespace")
		if namespace != "" {
			handlers.GetReplicaSets(ctx, namespace)
		}
	})

	r.GET("/api/v1/replicasets/", func(ctx *gin.Context) {
		handlers.GetReplicaSets(ctx, "")
	})

	r.GET("/api/v1/services/namespace/:namespace", func(ctx *gin.Context) {
		namespace := ctx.Param("namespace")
		if namespace != "" {
			handlers.GetServices(ctx, namespace)
		}
	})

	r.GET("/api/v1/services/", func(ctx *gin.Context) {
		handlers.GetServices(ctx, "")
	})

	return r
}

func main() {
	r := SetupRoutes()
	r.Run(":8080")
}
