package warehouse_grpc

import (
	"context"
	"google.golang.org/grpc"
)

type GRPCServer struct{}

func (s *GRPCServer) mustEmbedUnimplementedAdderServer() {
	//TODO implement me
	panic("implement me")
}

func (s *GRPCServer) Add(ctx context.Context, req *AddRequest) (*AddResponce, error) {
	return &AddResponce{Result: req.GetX() + req.GetY()}, nil
}

func (s *GRPCServer) TakeDelivery(ctx context.Context, in *Delivery, opts ...grpc.CallOption) (*Responce, error) {
	return &Responce{Anc: "good"}, nil
}
