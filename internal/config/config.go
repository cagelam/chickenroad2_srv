package config

import "github.com/zeromicro/go-zero/zrpc"

type JwtAuth struct {
	AccessSecret string
	AccessExpire int64
}
type Config struct {
	zrpc.RpcServerConf
	GWSrvConf            zrpc.RpcClientConf
	LoginSrvConf         zrpc.RpcClientConf
	PlayerCenterSrvConf  zrpc.RpcClientConf
	OrderSrvConf         zrpc.RpcClientConf
	OperatorProxySrvConf zrpc.RpcClientConf
	JwtAuth              JwtAuth
}
