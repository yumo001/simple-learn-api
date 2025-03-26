package svc

import (
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/yumo001/simple-learn-api/internal/config"
	i18n2 "github.com/yumo001/simple-learn-api/internal/i18n"
	"github.com/yumo001/simple-learn-api/internal/middleware"
	"github.com/yumo001/simple-learn-rpc/exampleclient"
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/casbin/casbin/v2"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config     config.Config
	Casbin     *casbin.Enforcer
	Authority  rest.Middleware
	Trans      *i18n.Translator
	ExampleRpc exampleclient.Example
}

func NewServiceContext(c config.Config) *ServiceContext {

	rds := c.RedisConf.MustNewUniversalRedis()

	cbn := c.CasbinConf.MustNewCasbinWithOriginalRedisWatcher(c.CasbinDatabaseConf.Type, c.CasbinDatabaseConf.GetDSN(), c.RedisConf)

	trans := i18n.NewTranslator(c.I18nConf, i18n2.LocaleFS)

	return &ServiceContext{
		Config:    c,
		Authority: middleware.NewAuthorityMiddleware(cbn, rds).Handle,
		Trans:     trans,

		ExampleRpc: exampleclient.NewExample(zrpc.MustNewClient(c.ExampleRpc)),
	}

}
