package api

import (
	"backend/internal/models"
	"backend/internal/store"
	"backend/internal/utils"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type UserHandler struct {
	userStore store.UserStore
	logger    *log.Logger
}

func NewUserHandler(store store.UserStore, logger *log.Logger) *UserHandler {
	return &UserHandler{
		userStore: store,
		logger:    logger,
	}
}

type registerRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Register godoc
// @Summary Register a new user
// @Description Register a new user with name, email, and password
// @Tags users
// @Accept json
// @Produce json
// @Param user body registerRequest true "User registration details"
// @Success 201 {object} object{user=object{id=int,name=string,email=string,role=string}}
// @Failure 400 {object} object{error=string} "Invalid request format or missing fields"
// @Failure 500 {object} object{error=string} "Internal server error"
// @Router /users/register [post]
func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var input registerRequest

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid request format"})
		return
	}

	input.Name = strings.TrimSpace(input.Name)
	input.Email = strings.TrimSpace(input.Email)
	input.Password = strings.TrimSpace(input.Password)

	if input.Name == "" || input.Email == "" || len(input.Password) < 6 {
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{
			"error": "name, email, and password (min 6 chars) are required",
		})
		return
	}

	user := &models.User{
		Name:  input.Name,
		Email: input.Email,
		Role:  "user",
	}

	if err := user.SetPassword(input.Password); err != nil {
		h.logger.Printf("ERROR: hashing password: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "internal error"})
		return
	}

	err := h.userStore.CreateUser(r.Context(), user)
	if err != nil {
		h.logger.Printf("ERROR: creating user: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "could not create user"})
		return
	}

	utils.WriteJSON(w, http.StatusCreated, utils.Envelope{
		"user": map[string]any{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
			"role":  user.Role,
		},
	})
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Login godoc
// @Summary Log in a user
// @Description Log in a user with email and password
// @Tags users
// @Accept json
// @Produce json
// @Param user body loginRequest true "User login credentials"
// @Success 200 {object} object{user=object{id=int,name=string,email=string,role=string}}
// @Failure 400 {object} object{error=string} "Invalid request format or missing fields"
// @Failure 401 {object} object{error=string} "Invalid email or password"
// @Failure 500 {object} object{error=string} "Internal server error"
// @Router /users/login [post]
func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var input loginRequest

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid request format"})
		return
	}

	input.Email = strings.TrimSpace(input.Email)
	input.Password = strings.TrimSpace(input.Password)

	if input.Email == "" || input.Password == "" {
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "email and password are required"})
		return
	}

	user, err := h.userStore.GetUserByEmail(r.Context(), input.Email)
	if err != nil {
		h.logger.Printf("ERROR: retrieving user: %v", err)
		utils.WriteJSON(w, http.StatusUnauthorized, utils.Envelope{"error": "invalid email or password"})
		return
	}

	ok, err := user.CheckPassword(input.Password)
	if err != nil {
		h.logger.Printf("ERROR: checking password: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "internal error"})
		return
	}

	if !ok {
		utils.WriteJSON(w, http.StatusUnauthorized, utils.Envelope{"error": "invalid email or password"})
		return
	}
	utils.WriteJSON(w, http.StatusOK, utils.Envelope{
		"user": map[string]any{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
			"role":  user.Role,
		},
	})
}
