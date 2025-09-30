package svc

import (
	"cocogame-max/chickenroad2_srv/internal/config"
	"cocogame-max/chickenroad2_srv/proto/conn_gw/conngwservice"
	"cocogame-max/chickenroad2_srv/proto/login_srv/loginsrv"
	"cocogame-max/chickenroad2_srv/proto/operatorproxy_srv/operatorproxysrv"
	"cocogame-max/chickenroad2_srv/proto/order_srv/ordersrv"
	"cocogame-max/chickenroad2_srv/proto/playercenter_srv/playercentersrv"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config           config.Config
	GW               conngwservice.ConnGwService
	LoginSrv         loginsrv.LoginSrv
	PlayerCenterSrv  playercentersrv.PlayerCenterSrv
	OrderSrv         ordersrv.OrderSrv
	OperatorProxySrv operatorproxysrv.OperatorProxySrv
	GameID           int32
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:           c,
		GW:               conngwservice.NewConnGwService(zrpc.MustNewClient(c.GWSrvConf)),
		LoginSrv:         loginsrv.NewLoginSrv(zrpc.MustNewClient(c.LoginSrvConf)),
		PlayerCenterSrv:  playercentersrv.NewPlayerCenterSrv(zrpc.MustNewClient(c.PlayerCenterSrvConf)),
		OrderSrv:         ordersrv.NewOrderSrv(zrpc.MustNewClient(c.OrderSrvConf)),
		OperatorProxySrv: operatorproxysrv.NewOperatorProxySrv(zrpc.MustNewClient(c.OperatorProxySrvConf)),
		GameID:           300,
	}
}
