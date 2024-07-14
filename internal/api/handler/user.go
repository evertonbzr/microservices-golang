package handler

import (
	"net/http"

	"github.com/evertonbzr/microservices-golang/internal/service"
	"github.com/evertonbzr/microservices-golang/pkg/utils"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type UserHandler struct {
	Service *service.UserService
}

func NewUserHandler(db *gorm.DB, cache *redis.Client) *UserHandler {
	return &UserHandler{
		Service: service.NewUserService(db, cache),
	}
}

func (h *UserHandler) ListUsers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := h.Service.ListUsers()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, users)
	}
}
