package worker

import (
	"github.com/Zerohated/tools/pkg/dao"
	"github.com/Zerohated/tools/pkg/logger"
)

var (
	pgConn = dao.PgConn
	log    = logger.Logger
)

const (
	evoEnvProd = "prod_external"
	evoEnvTest = "pp_external"
	evoAppID   = "mpcoupon"
)

const (
	StatusStop    = -1
	StatusWaiting = 0
	StatusRunning = 1
)

const (
	ProviderByCSV = iota + 1
	ProviderByDB
)

type Provider interface {
	GetFactory() *Factory
	GetStatus() int
	GetType() int
	GetDetail() map[string]interface{}
	Supply()
}

type BaseProvider struct {
	Factory *Factory
	Status  int
	Type    int
}
type CSVProvider struct {
	BaseProvider
	SourceFileDir string
}
