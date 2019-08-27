package rpc

import (
	"context"
	"time"

	"github.com/labstack/gommon/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

// Client contains rpc connection confs
type Client struct {
	address    string
	connection *grpc.ClientConn
	ctx        context.Context
}

// NewClient Create a new rpc client
func NewClient(address string) Client {
	return Client{
		address: address,
		ctx:     context.Background(),
	}
}

// Connect connect to another aplication via rpc
func (cl Client) Connect() *grpc.ClientConn {
	k := keepalive.ClientParameters{
		Time:                10 * time.Second, // send pings every 10 seconds if there is no activity
		Timeout:             time.Second,      // wait 1 second for ping ack before considering the connection dead
		PermitWithoutStream: true,             // send pings even without active streams
	}
	opts := grpc.WaitForReady(false)
	conn, err := grpc.Dial(
		cl.address,
		grpc.WithInsecure(),
		grpc.WithKeepaliveParams(k),
		grpc.WithDefaultCallOptions(opts),
	)
	if err != nil {
		log.Error("Falhei. Nao consegui conectar em "+cl.address, err.Error())
	} else {
		log.Info("Conectando em", cl.address)
	}
	cl.connection = conn
	return conn
}

// Disconnect
func (cl Client) Disconnect() {
	if cl.connection != nil {
		log.Info("Desconectei de " + cl.address)
		cl.connection.Close()
	}
}
