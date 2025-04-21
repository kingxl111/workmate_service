package apihandler

import (
	"github.com/kingxl111/workmate_service/pkg/api/oapigen/tasks"
	openapiTypes "github.com/oapi-codegen/runtime/types"
	"net/http"
)

var _ tasks.ServerInterface = (*Handler)(nil)

type Handler struct {
	taskService TaskService
}

func NewHandler(service TaskService) *Handler {
	return &Handler{
		taskService: service,
	}
}

func (h *Handler) PostTasks(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) GetTasksId(w http.ResponseWriter, r *http.Request, id openapiTypes.UUID) {

}
