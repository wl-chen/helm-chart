package main

// import "github.com/rs/cors/wrapper/gin"
import (
	"net/http"
	// "os"

	"github.com/docker/distribution/uuid"
	"github.com/gin-gonic/gin"
)

var podID string

func init() {
	podID = uuid.Generate().String()
}

func main() {
	r := gin.Default()
	r.GET("/pod/ip", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, getpodID())
	})
	r.Run(":8088")
}

func getpodID() string {
	return podID
}