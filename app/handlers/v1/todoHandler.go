package v1

import (
	"github.com/fernandochristyanto/retsgo/app/core"
	appservice "github.com/fernandochristyanto/retsgo/services"
	"github.com/gorilla/mux"
	"github.com/json-iterator/go"
	"net/http"
	"strconv"
)

type TodoHandler struct {
	todoService *appservice.TodoService
}

func (handler TodoHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	todos := handler.todoService.GetAll()
	json := jsoniter.ConfigCompatibleWithStandardLibrary

	todosJSON, _ := json.Marshal(todos)
	w.Header().Set("Content-Type", "application/json")
	w.Write(todosJSON)
}

func (handler TodoHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(core.NewHttpException(http.StatusBadRequest, "ID must be an int"))
	}
	todo := handler.todoService.GetByID(id)
	json := jsoniter.ConfigCompatibleWithStandardLibrary

	todoJSON, _ := json.Marshal(todo)
	w.Header().Set("Content-Type", "application/json")
	w.Write(todoJSON)
}

func NewTodoHandler(todoService *appservice.TodoService) *TodoHandler {
	return &TodoHandler{
		todoService: todoService,
	}
}
