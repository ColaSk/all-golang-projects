# go-zero-template

## 启动方式

```bash

etcd

cd mall/user/rpc && go run user.go -f etc/user.yaml

cd mall/order/api && go run order.go -f etc/order.yaml

curl -i -X GET http://localhost:8888/api/order/get/1
````
