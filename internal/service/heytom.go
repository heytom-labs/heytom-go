package service

import (
	"context"

	pb "github.com/heytom-labs/heytom-go/api/heytom/v1"
	"github.com/heytom-labs/heytom-go/internal/biz"
)

type HeytomService struct {
	pb.UnimplementedHeytomServer

	uc *biz.GreeterUsecase
}

func NewHeytomService(uc *biz.GreeterUsecase) *HeytomService {
	return &HeytomService{uc: uc}
}

func (s *HeytomService) CreateGreeter(ctx context.Context, req *pb.CreateGreeterRequest) (*pb.CreateGreeterReply, error) {
	g, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Name: req.Name})
	if err != nil {
		return nil, err
	}
	return &pb.CreateGreeterReply{
		Id:   g.ID,
		Name: g.Name,
	}, nil
}

func (s *HeytomService) GetGreeter(ctx context.Context, req *pb.GetGreeterRequest) (*pb.GetGreeterReply, error) {
	g, err := s.uc.GetGreeter(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetGreeterReply{
		Id:   g.ID,
		Name: g.Name,
	}, nil
}
