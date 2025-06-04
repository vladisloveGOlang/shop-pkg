package warehouse_grpc

import (
	"context"
)

type GRPCServer struct{}

func (s *GRPCServer) mustEmbedUnimplementedAdderServer() {
	//TODO implement me
	panic("implement me")
}

func (s *GRPCServer) Add(ctx context.Context, req *AddRequest) (*AddResponce, error) {
	return &AddResponce{Result: req.GetX() + req.GetY()}, nil
}

func (s *GRPCServer) TakeDelivery(ctx context.Context, delivery *Delivery) (*Responce, error) {
	return &Responce{Anc: "good"}, nil
}
