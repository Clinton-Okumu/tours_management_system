package api

import (
	"backend/internal/models"
	"backend/internal/store"
	"backend/internal/utils"
	"encoding/json"
	"log"
	"net/http"
)

// ReviewHandler handles review-related HTTP requests
type ReviewHandler struct {
	reviewStore store.ReviewStore
	logger      *log.Logger
}

// NewReviewHandler returns a new ReviewHandler
func NewReviewHandler(store store.ReviewStore, logger *log.Logger) *ReviewHandler {
	return &ReviewHandler{
		reviewStore: store,
		logger:      logger,
	}
}

// CreateReview godoc
// @Summary Create a new review
// @Description Submit a review for a tour
// @Tags reviews
// @Accept json
// @Produce json
// @Param review body models.Review true "Review object"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /review [post]
func (rh *ReviewHandler) CreateReview(w http.ResponseWriter, r *http.Request) {
	var review models.Review

	err := json.NewDecoder(r.Body).Decode(&review)
	if err != nil {
		rh.logger.Printf("ERROR: decoding review: %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid request format"})
		return
	}

	createdReview, err := rh.reviewStore.CreateReview(r.Context(), &review)
	if err != nil {
		rh.logger.Printf("ERROR: creating review: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "could not create review"})
		return
	}

	utils.WriteJSON(w, http.StatusCreated, utils.Envelope{"review": createdReview})
}

// DeleteReview godoc
// @Summary Delete a review by ID
// @Description Delete a specific review using its ID
// @Tags reviews
// @Produce json
// @Param id path int true "Review ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /review/{id} [delete]
func (rh *ReviewHandler) DeleteReview(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ReadIDParam(r)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid review id"})
		return
	}

	err = rh.reviewStore.DeleteReview(r.Context(), uint(id))
	if err != nil {
		rh.logger.Printf("ERROR: deleting review: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "could not delete review"})
		return
	}
	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"message": "review deleted"})
}
