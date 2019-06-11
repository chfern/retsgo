package v1

import (
	"github.com/fernandochristyanto/retsgo/app/core"
	apprepo "github.com/fernandochristyanto/retsgo/repositories"
	"github.com/gorilla/mux"
	"github.com/json-iterator/go"
	"net/http"
	"strconv"
)

type UserHandler struct {
	userRepo *apprepo.UserRepository
}

func (handler UserHandler) GetTodos(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(core.NewHttpException(http.StatusBadRequest, "ID must be an int"))
	}
	todos := handler.userRepo.GetTodos(int64(id))
	json := jsoniter.ConfigCompatibleWithStandardLibrary

	todosJSON, _ := json.Marshal(todos)
	w.Header().Set("Content-Type", "application/json")
	w.Write(todosJSON)
}

func NewUserHandler(userRepo *apprepo.UserRepository) *UserHandler {
	return &UserHandler{
		userRepo: userRepo,
	}
}
