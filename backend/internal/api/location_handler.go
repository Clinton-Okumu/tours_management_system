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

type LocationHandler struct {
	locationStore store.LocationStore
	logger        *log.Logger
}

func NewLocationHandler(store store.LocationStore, logger *log.Logger) *LocationHandler {
	return &LocationHandler{
		locationStore: store,
		logger:        logger,
	}
}

func (lh *LocationHandler) CreateLocation(w http.ResponseWriter, r *http.Request) {
	var location models.Location

	err := json.NewDecoder(r.Body).Decode(&location)
	if err != nil {
		lh.logger.Printf("ERROR: decoding location: %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid request format"})
		return
	}

	createdLocation, err := lh.locationStore.CreateLocation(r.Context(), &location)
	if err != nil {
		lh.logger.Printf("ERROR: creating location: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "could not create location"})
		return
	}

	utils.WriteJSON(w, http.StatusCreated, utils.Envelope{"location": createdLocation})
}

func (lh *LocationHandler) GetLocationByID(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ReadIDParam(r)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid location id"})
		return
	}

	location, err := lh.locationStore.GetLocationByID(r.Context(), uint(id))
	if err != nil {
		if errors.Is(err, store.ErrLocationNotFound) {
			utils.WriteJSON(w, http.StatusNotFound, utils.Envelope{"error": "location not found"})
			return
		}
		lh.logger.Printf("ERROR: fetching location: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "could not retrieve location"})
		return
	}

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"location": location})
}

func (lh *LocationHandler) UpdateLocation(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ReadIDParam(r)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid location id"})
		return
	}

	var location models.Location
	err = json.NewDecoder(r.Body).Decode(&location)
	if err != nil {
		lh.logger.Printf("ERROR: decoding location: %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid request format"})
		return
	}

	location.ID = uint(id)

	err = lh.locationStore.UpdateLocation(r.Context(), &location)
	if err != nil {
		lh.logger.Printf("ERROR: updating location: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "could not update location"})
		return
	}

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"message": "location updated"})
}

func (lh *LocationHandler) DeleteLocation(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ReadIDParam(r)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid location id"})
		return
	}

	err = lh.locationStore.DeleteLocation(r.Context(), uint(id))
	if err != nil {
		lh.logger.Printf("ERROR: deleting location: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "could not delete location"})
		return
	}
}
