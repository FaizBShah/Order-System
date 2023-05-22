package clients

import "api-gateway/clients/productclient"

func InitClients() {
	productclient.InitProductClient()
}
