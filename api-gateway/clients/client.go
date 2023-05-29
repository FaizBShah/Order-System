package clients

import (
	"api-gateway/clients/authclient"
	"api-gateway/clients/orderclient"
	"api-gateway/clients/productclient"
)

func InitClients() {
	productclient.InitProductClient()
	orderclient.InitOrderClient()
	authclient.InitAuthClient()
}
