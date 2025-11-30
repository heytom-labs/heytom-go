package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/heytom-labs/heytom-go/internal/biz"
)

type greeterRepo struct {
	data *Data
	log  *log.Helper
}

func NewGreeterRepo(data *Data, logger log.Logger) biz.GreeterRepo {
	return &greeterRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *greeterRepo) Save(ctx context.Context, g *biz.Greeter) (*biz.Greeter, error) {
	// TODO: implement database save
	g.ID = 1
	return g, nil
}

func (r *greeterRepo) Get(ctx context.Context, id int64) (*biz.Greeter, error) {
	// TODO: implement database get
	return &biz.Greeter{
		ID:   id,
		Name: "Hello Heytom",
	}, nil
}
