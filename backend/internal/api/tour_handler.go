package api

import (
	"backend/internal/models"
	"backend/internal/store"
	"backend/internal/utils"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

// TourHandler handles tour-related requests
type TourHandler struct {
	tourStore store.TourStore
	logger    *log.Logger
}

func NewTourHandler(store store.TourStore, logger *log.Logger) *TourHandler {
	return &TourHandler{
		tourStore: store,
		logger:    logger,
	}
}

// CreateTour godoc
// @Summary Create a new tour
// @Description Add a new tour to the database
// @Tags tours
// @Accept json
// @Produce json
// @Param tour body models.Tour true "Tour data"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /tour [post]
func (th *TourHandler) CreateTour(w http.ResponseWriter, r *http.Request) {
	var tour models.Tour

	err := json.NewDecoder(r.Body).Decode(&tour)
	if err != nil {
		th.logger.Printf("ERROR: decoding tour: %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid request format"})
		return
	}

	createdTour, err := th.tourStore.CreateTour(r.Context(), &tour)
	if err != nil {
		th.logger.Printf("ERROR: creating tour: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "could not create tour"})
		return
	}

	utils.WriteJSON(w, http.StatusCreated, utils.Envelope{"tour": createdTour})
}

// GetTourByID godoc
// @Summary Get tour by ID
// @Description Retrieve a specific tour by its ID
// @Tags tours
// @Produce json
// @Param id path int true "Tour ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /tour/{id} [get]
func (th *TourHandler) GetTourByID(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ReadIDParam(r)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid tour id"})
		return
	}

	tour, err := th.tourStore.GetTourByID(r.Context(), uint(id))
	if err != nil {
		if errors.Is(err, store.ErrTourNotFound) {
			utils.WriteJSON(w, http.StatusNotFound, utils.Envelope{"error": "tour not found"})
			return
		}
		th.logger.Printf("ERROR: fetching tour: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "could not retrieve tour"})
		return
	}

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"tour": tour})
}

// UpdateTour godoc
// @Summary Update a tour
// @Description Update an existing tour by ID
// @Tags tours
// @Accept json
// @Produce json
// @Param id path int true "Tour ID"
// @Param tour body models.Tour true "Updated tour data"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /tour/{id} [put]
func (th *TourHandler) UpdateTour(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ReadIDParam(r)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid tour id"})
		return
	}

	var tour models.Tour
	err = json.NewDecoder(r.Body).Decode(&tour)
	if err != nil {
		th.logger.Printf("ERROR: decoding tour: %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid request format"})
		return
	}

	tour.ID = uint(id)

	err = th.tourStore.UpdateTour(r.Context(), &tour)
	if err != nil {
		th.logger.Printf("ERROR: updating tour: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "could not update tour"})
		return
	}

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"message": "tour updated"})
}

// DeleteTour godoc
// @Summary Delete a tour
// @Description Delete a specific tour by ID
// @Tags tours
// @Produce json
// @Param id path int true "Tour ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /tour/{id} [delete]
func (th *TourHandler) DeleteTour(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ReadIDParam(r)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid tour id"})
		return
	}

	err = th.tourStore.DeleteTour(r.Context(), uint(id))
	if err != nil {
		th.logger.Printf("ERROR: deleting tour: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "could not delete tour"})
		return
	}

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"message": "tour deleted"})
}
