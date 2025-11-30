package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/heytom-labs/heytom-go/internal/conf"
)

var ProviderSet = wire.NewSet(NewData, NewGreeterRepo)

type Data struct {
	// TODO: add database connection
}

func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{}, cleanup, nil
}
