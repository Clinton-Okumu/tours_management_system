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
