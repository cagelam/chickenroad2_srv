main:
	goctl rpc protoc chickenroad2_srv.proto		--go_out=..		--go-grpc_out=..	--zrpc_out=..  --style=go_zero

gw:
	goctl rpc protoc conn_gw.proto		--go_out=conn_gw		--go-grpc_out=conn_gw	--zrpc_out=conn_gw  --style=go_zero

login:
	goctl rpc protoc login_srv.proto			--go_out=login_srv			--go-grpc_out=login_srv			--zrpc_out=login_srv --style=go_zero


playercenter:
	goctl rpc protoc playercenter_srv.proto		--go_out=playercenter_srv		--go-grpc_out=playercenter_srv	--zrpc_out=playercenter_srv	 --style=go_zero

order:
	goctl rpc protoc order_srv.proto			--go_out=order_srv			--go-grpc_out=order_srv	--zrpc_out=order_srv --style=go_zero

operatorproxy:
	goctl rpc protoc operatorproxy_srv.proto	--go_out=operatorproxy_srv	--go-grpc_out=operatorproxy_srv	--zrpc_out=operatorproxy_srv