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

// BookingHandler handles booking-related HTTP requests
type BookingHandler struct {
	bookingStore store.BookingStore
	logger       *log.Logger
}

func NewBookingHandler(store store.BookingStore, logger *log.Logger) *BookingHandler {
	return &BookingHandler{
		bookingStore: store,
		logger:       logger,
	}
}

// CreateBooking godoc
// @Summary Create a new booking
// @Description Create a new booking with tour ID, user ID, and booking details
// @Tags bookings
// @Accept json
// @Produce json
// @Param booking body models.Booking true "Booking object"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /booking [post]
func (bh *BookingHandler) CreateBooking(w http.ResponseWriter, r *http.Request) {
	var booking models.Booking

	err := json.NewDecoder(r.Body).Decode(&booking)
	if err != nil {
		bh.logger.Printf("ERROR: decoding booking: %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid request format"})
		return
	}

	createdBooking, err := bh.bookingStore.CreateBooking(r.Context(), &booking)
	if err != nil {
		bh.logger.Printf("ERROR: creating booking: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "could not create booking"})
		return
	}

	utils.WriteJSON(w, http.StatusCreated, utils.Envelope{"booking": createdBooking})
}

// GetBookingByID godoc
// @Summary Get a booking by ID
// @Description Retrieve a booking using its ID
// @Tags bookings
// @Produce json
// @Param id path int true "Booking ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /booking/{id} [get]
func (bh *BookingHandler) GetBookingByID(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ReadIDParam(r)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid booking id"})
		return
	}

	booking, err := bh.bookingStore.GetBookingByID(r.Context(), uint(id))
	if err != nil {
		if errors.Is(err, store.ErrBookingNotFound) {
			utils.WriteJSON(w, http.StatusNotFound, utils.Envelope{"error": "booking not found"})
			return
		}
		bh.logger.Printf("ERROR: fetching booking: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "could not retrieve booking"})
		return
	}

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"booking": booking})
}

// UpdateBooking godoc
// @Summary Update a booking by ID
// @Description Update an existing booking with new information
// @Tags bookings
// @Accept json
// @Produce json
// @Param id path int true "Booking ID"
// @Param booking body models.Booking true "Updated booking object"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /booking/{id} [put]
func (bh *BookingHandler) UpdateBooking(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ReadIDParam(r)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid booking id"})
		return
	}

	var booking models.Booking
	err = json.NewDecoder(r.Body).Decode(&booking)
	if err != nil {
		bh.logger.Printf("ERROR: decoding booking: %v", err)
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid request format"})
		return
	}

	booking.ID = uint(id)

	err = bh.bookingStore.UpdateBooking(r.Context(), &booking)
	if err != nil {
		bh.logger.Printf("ERROR: updating booking: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "could not update booking"})
		return
	}

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"message": "booking updated"})
}

// DeleteBooking godoc
// @Summary Delete a booking by ID
// @Description Delete a specific booking using its ID
// @Tags bookings
// @Produce json
// @Param id path int true "Booking ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /booking/{id} [delete]
func (bh *BookingHandler) DeleteBooking(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ReadIDParam(r)
	if err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, utils.Envelope{"error": "invalid booking id"})
		return
	}

	err = bh.bookingStore.DeleteBooking(r.Context(), uint(id))
	if err != nil {
		bh.logger.Printf("ERROR: deleting booking: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "could not delete booking"})
		return
	}

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"message": "booking deleted"})
}
