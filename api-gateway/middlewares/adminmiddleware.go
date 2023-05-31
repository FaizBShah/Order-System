package middlewares

import (
	"api-gateway/dto"
	"encoding/json"
	"net/http"
)

func AdminMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(respWriter http.ResponseWriter, req *http.Request) {
		userType := req.Context().Value(USER_TYPE).(string)

		if userType != "ADMIN" {
			errMessage := dto.Error{Status: http.StatusForbidden, Message: "user does not have admin priviledges to access this resource"}
			respWriter.WriteHeader(errMessage.Status)
			json.NewEncoder(respWriter).Encode(errMessage)
			return
		}

		next.ServeHTTP(respWriter, req)
	})
}
