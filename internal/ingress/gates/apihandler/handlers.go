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
	
	//var req merchstoreapi.AuthRequest
	//if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
	//	h.respondWithError(w, http.StatusBadRequest, err.Error())
	//	return
	//}
	//ctx := r.Context()
	//resp, err := h.userService.Authenticate(ctx, &users.AuthRequest{
	//	Username: req.Username,
	//	Password: req.Password,
	//})
	//if err != nil {
	//	if errors.Is(err, users.ErrorWrongPassword) {
	//		h.respondWithError(w, http.StatusBadRequest, "wrong password")
	//		return
	//	}
	//	h.respondWithError(w, http.StatusInternalServerError, intServerError)
	//	return
	//}
	//h.respondWithJSON(w, http.StatusOK, merchstoreapi.AuthResponse{Token: &resp.Token})
}

func (h *Handler) GetTasksId(w http.ResponseWriter, r *http.Request, id openapiTypes.UUID) {

}
