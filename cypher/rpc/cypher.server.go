package rpc

import (
	"context"

	"google.golang.org/grpc"

	pb "grpc-talk/proto"
)

type CypherServer struct{}

// NewCypherServer cria o serviço que contera as funçoes do contrato
func NewCypherServer(s *grpc.Server) *CypherServer {
	server := &CypherServer{}

	if s != nil {
		pb.RegisterCypherServiceServer(s, server)
	}

	return server
}

// Encode encripta string de acordo com cifra de cesar
func (cypher CypherServer) Encode(ctx context.Context, req *pb.CypherRequest) (*pb.CypherResponse, error) {
	options := req.GetOptions()
	encode := rotate(options.GetText(), int(options.GetShift()))
	resp := &pb.CypherResponse{
		EncrypedText: encode,
		Options:      options,
	}
	return resp, nil
}

// Decode descripta string de acordo com cifra de cesar
func (cypher CypherServer) Decode(ctx context.Context, req *pb.CypherRequest) (*pb.CypherResponse, error) {
	options := req.GetOptions()
	encode := rotate(options.GetText(), -int(options.GetShift()))
	resp := &pb.CypherResponse{
		EncrypedText: encode,
		Options:      options,
	}
	return resp, nil
}

func rotate(text string, shift int) string {
	shift = (shift%26 + 26) % 26 // [0, 25]
	b := make([]byte, len(text))
	for i := 0; i < len(text); i++ {
		t := text[i]
		var a int
		switch {
		case 'a' <= t && t <= 'z':
			a = 'a'
		case 'A' <= t && t <= 'Z':
			a = 'A'
		default:
			b[i] = t
			continue
		}
		b[i] = byte(a + ((int(t)-a)+shift)%26)
	}
	return string(b)
}
