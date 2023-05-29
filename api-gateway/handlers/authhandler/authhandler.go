package authhandler

import (
	"api-gateway/clients/authclient"
	"api-gateway/dto"
	proto "api-gateway/proto/auth"
	"encoding/json"
	"net/http"
	"strings"
)

func RegisterUser(respWriter http.ResponseWriter, req *http.Request) {
	respWriter.Header().Set("Content-Type", "application/json")

	var newUser dto.RegisterUserRequest

	if err := json.NewDecoder(req.Body).Decode(&newUser); err != nil {
		errMessage := dto.Error{Status: http.StatusBadRequest, Message: err.Error()}
		respWriter.WriteHeader(errMessage.Status)
		json.NewEncoder(respWriter).Encode(errMessage)
		return
	}

	registeredUser, err := authclient.AuthServiceClient.RegisterUser(req.Context(), &proto.RegisterUserRequest{
		Name:     newUser.Name,
		Email:    newUser.Email,
		Password: newUser.Password,
		UserType: newUser.UserType,
	})

	if err != nil {
		errMessage := dto.Error{Status: http.StatusBadRequest, Message: strings.Replace(err.Error(), "rpc error: code = Unknown desc = ", "", 1)}
		respWriter.WriteHeader(errMessage.Status)
		json.NewEncoder(respWriter).Encode(errMessage)
		return
	}

	res := dto.RegisterUserResponse{
		Id:       registeredUser.Id,
		Name:     registeredUser.Name,
		Email:    registeredUser.Email,
		UserType: registeredUser.UserType,
	}

	respWriter.WriteHeader(http.StatusCreated)
	json.NewEncoder(respWriter).Encode(res)
}

func LoginUser(respWriter http.ResponseWriter, req *http.Request) {
	respWriter.Header().Set("Content-Type", "application/json")

	var userCredentials dto.LoginUserRequest

	if err := json.NewDecoder(req.Body).Decode(&userCredentials); err != nil {
		errMessage := dto.Error{Status: http.StatusBadRequest, Message: err.Error()}
		respWriter.WriteHeader(errMessage.Status)
		json.NewEncoder(respWriter).Encode(errMessage)
		return
	}

	token, err := authclient.AuthServiceClient.LoginUser(req.Context(), &proto.LoginUserRequest{
		Email:    userCredentials.Email,
		Password: userCredentials.Password,
	})

	if err != nil {
		errMessage := dto.Error{Status: http.StatusBadRequest, Message: strings.Replace(err.Error(), "rpc error: code = Unknown desc = ", "", 1)}
		respWriter.WriteHeader(errMessage.Status)
		json.NewEncoder(respWriter).Encode(errMessage)
		return
	}

	res := dto.LoginUserResponse{
		Token: token.Token,
	}

	respWriter.WriteHeader(http.StatusOK)
	json.NewEncoder(respWriter).Encode(res)
}
