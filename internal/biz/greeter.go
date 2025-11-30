package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Greeter struct {
	ID   int64
	Name string
}

type GreeterRepo interface {
	Save(context.Context, *Greeter) (*Greeter, error)
	Get(context.Context, int64) (*Greeter, error)
}

type GreeterUsecase struct {
	repo GreeterRepo
	log  *log.Helper
}

func NewGreeterUsecase(repo GreeterRepo, logger log.Logger) *GreeterUsecase {
	return &GreeterUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *GreeterUsecase) CreateGreeter(ctx context.Context, g *Greeter) (*Greeter, error) {
	uc.log.WithContext(ctx).Infof("CreateGreeter: %v", g.Name)
	return uc.repo.Save(ctx, g)
}

func (uc *GreeterUsecase) GetGreeter(ctx context.Context, id int64) (*Greeter, error) {
	uc.log.WithContext(ctx).Infof("GetGreeter: %d", id)
	return uc.repo.Get(ctx, id)
}
