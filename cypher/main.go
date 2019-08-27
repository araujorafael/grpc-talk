package main

import (
	"github.com/labstack/gommon/log"
	"grpc-talk/libs/remoteprocedurecall"
	"grpc-talk/cypher/rpc"
)

func main() {
	port := ":8001"
	rpcServer := remoteprocedurecall.NewServer(port)
	if rpcServer == nil {
		log.Fatal("Nao consigo escutar na porta:", port)
	}

	rpc.NewCypherServer(rpcServer.Grpc)

	log.Info("Listening on", port)
	log.Fatal(rpcServer.Start())
}
