package userservice

import (
	"net/http"

	"github.com/EliasSantiago/api-contatos/adapter/http/errors"
	"github.com/EliasSantiago/api-contatos/core/dto"
)

// @Summary Create new user
// @Description Create new user
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body dto.CreateUserRequest true "user"
// @Success 200 {object} domain.User
// @Router /user [post]
func (service service) Publish(response http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()
	userRequest, err := dto.FromJSONCreateUserRequest(request.Body)
	if err != nil {
		errors.InvalidParameters()
		return
	}
	err = service.usecase.Publish(userRequest)
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

}
