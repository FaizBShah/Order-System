package middlewares

import (
	"api-gateway/clients/authclient"
	"api-gateway/dto"
	proto "api-gateway/proto/auth"
	"context"
	"encoding/json"
	"net/http"
	"strings"
)

type key int

const (
	USER_ID key = iota
	USER_EMAIL
	USER_TYPE
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(respWriter http.ResponseWriter, req *http.Request) {
		header := req.Header.Get("Authorization")

		if header == "" || !strings.HasPrefix(header, "Bearer ") {
			errMessage := dto.Error{Status: http.StatusBadRequest, Message: "authentication token missing from the request"}
			respWriter.WriteHeader(errMessage.Status)
			json.NewEncoder(respWriter).Encode(errMessage)
			return
		}

		token := strings.TrimPrefix(header, "Bearer ")

		user, err := authclient.AuthServiceClient.AuthencticateUser(req.Context(), &proto.AuthenticateUserRequest{
			Token: token,
		})

		if err != nil {
			errMessage := dto.Error{Status: http.StatusBadRequest, Message: strings.Replace(err.Error(), "rpc error: code = Unknown desc = ", "", 1)}
			respWriter.WriteHeader(errMessage.Status)
			json.NewEncoder(respWriter).Encode(errMessage)
			return
		}

		ctx := context.WithValue(req.Context(), USER_ID, user.Id)
		ctx = context.WithValue(ctx, USER_EMAIL, user.Email)
		ctx = context.WithValue(ctx, USER_TYPE, user.UserType.String())

		req = req.WithContext(ctx)

		next.ServeHTTP(respWriter, req)
	})
}
