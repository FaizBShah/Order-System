package clients

import (
	"api-gateway/clients/orderclient"
	"api-gateway/clients/productclient"
)

func InitClients() {
	productclient.InitProductClient()
	orderclient.InitOrderClient()
}
