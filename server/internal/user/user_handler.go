package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *UserService
}

func NewHandler(s *UserService) *Handler {
	return &Handler{service: s}
}

func (h *Handler) CreateUser(c *gin.Context) {
	var u CreateUserRequest
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.service.CreateUser(c.Request.Context(), &u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, res)
}

func (h *Handler) Login(c *gin.Context) {
	var u LoginUserRequest
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.service.Login(c.Request.Context(), &u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("jwt", res.accessToken, 3600, "/", "localhost", false, true)

	res = &LoginUserResponse{
		ID:       res.ID,
		Username: res.Username,
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) Logout(c *gin.Context) {
	c.SetCookie("jwt", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}
