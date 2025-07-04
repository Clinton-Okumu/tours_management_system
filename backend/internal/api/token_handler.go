package api

import (
	"backend/internal/store"
	"backend/internal/tokens"
	"backend/internal/utils"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type TokenHandler struct {
	tokenStore store.TokenStore
	userStore  store.UserStore
	logger     *log.Logger
}

type createTokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewTokenHandler(tokenStore store.TokenStore, userStore store.UserStore, logger *log.Logger) *TokenHandler {
	return &TokenHandler{
		tokenStore: tokenStore,
		userStore:  userStore,
		logger:     logger,
	}
}

// HandleCreateToken godoc
// @Summary Generate an authentication token
// @Description Validates email & password and returns an authentication token
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body createTokenRequest true "User credentials"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /tokens/authentication [post]
func (h *TokenHandler) HandleCreateToken(w http.ResponseWriter, r *http.Request) {
	var req createTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Printf("ERROR decoding createTokenRequest: %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "Invalid request body"})
		return
	}

	user, err := h.userStore.GetUserByEmail(r.Context(), req.Email)
	if err != nil {
		h.logger.Printf("ERROR finding user: %v", err)
		utils.WriteJSON(w, http.StatusUnauthorized, utils.Envelope{"error": "Invalid credentials"})
		return
	}

	match, err := user.CheckPassword(req.Password)
	if err != nil || !match {
		utils.WriteJSON(w, http.StatusUnauthorized, utils.Envelope{"error": "Invalid credentials"})
		return
	}

	tokenModel, plaintext, err := h.tokenStore.CreateNewToken(r.Context(), user.ID, 24*time.Hour, tokens.ScopeAuth)
	if err != nil {
		h.logger.Printf("ERROR creating token: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "Failed to generate token"})
		return
	}

	utils.WriteJSON(w, http.StatusCreated, utils.Envelope{
		"token":  plaintext,
		"expiry": tokenModel.Expiry,
	})
}
