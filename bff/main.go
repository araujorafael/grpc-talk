package main

import (
	"net/http"

	rpc "grpc-talk/bff/rpc"
	pb "grpc-talk/proto"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type CesarMessage struct {
	Message string `json:"message" binding:"required"`
	Shift   int    `json:"shift,omitempty" binding:"required"`
}

type Response struct {
	Data CesarMessage `json:"data,omitempty"`
	Err  string       `json:"error,omitempty"`
}

func getRpcConn(url string) *grpc.ClientConn {
	client := rpc.NewClient(url)
	return client.Connect()
}

func setupRouter() *gin.Engine {
	conn := getRpcConn("localhost:8001")
	cypherConn := pb.NewCypherServiceClient(conn)

	r := gin.Default()

	// Get user value
	r.GET("/cesarcypher/encode", func(c *gin.Context) {
		data := new(CesarMessage)
		err := c.BindJSON(data)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, Response{Err: "Incorrect json values"})
			return
		}

		message := &pb.CypherRequest{
			Options: &pb.CypherOptions{
				Text:  data.Message,
				Shift: int32(data.Shift),
			},
		}

		encoded, err := cypherConn.Encode(c, message)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, err)
			return
		}

		encodedMessage := CesarMessage{Message: encoded.GetEncrypedText()}
		c.JSON(http.StatusOK, Response{Data: encodedMessage})
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
