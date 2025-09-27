package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"tt2/internal/models"
	"github.com/gorilla/mux"
)

type CreateOutPut struct {
	ID int `json:"id"`
}


func (h *Handler) createSubscription(w http.ResponseWriter, r *http.Request) {
	const op = "handler.createSubcription"
	log := h.log.With("op", op)
	var sub models.SubscriptionForCreate
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Error("failed to read request body", "error", err)
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	if err := json.Unmarshal(body, &sub); err != nil {
		log.Error("failed to unmarshal error", "error", err)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	id, err := h.repo.Create(sub)
	if err != nil {
		log.Error("with h.repo.Create -> failed to create subscription", "error", err)
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}
	output := CreateOutPut{ID: id}
	message, err := json.Marshal(output)
	if err != nil {
		log.Error("failed to marshal response", "error", err)
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(message)
}

type GetAllOutPut struct {
	Subscriptions []models.Subscription `json:"subscriptions"`
}

func (h *Handler) getAllSubscriptions(w http.ResponseWriter, r *http.Request) {
	const op = "handler.getAllSubscriptions"
	log := h.log.With("op", op)
	subs, err := h.repo.GetAll()
	if err != nil {
		log.Error("with h.repo.GetAll -> failed to get subscriptions", "error", err)
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}
	output := GetAllOutPut{Subscriptions: subs}
	message, err := json.Marshal(output)
	if err != nil {
		log.Error("failed to marshal response", "error", err)
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(message)
}

func (h *Handler) getSubscription(w http.ResponseWriter, r *http.Request) {
	const op = "handler.getSubscription"
	log := h.log.With("op", op)
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	sub, err := h.repo.GetByID(id)
	if err != nil {
		log.Error("with h.repo.GetByID -> failed to get subscription", "error", err)
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}
	message, err := json.Marshal(sub)
	if err != nil {
		log.Error("failed to marshal response", "error", err)
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(message)
}

func (h *Handler) deleteSubscription(w http.ResponseWriter, r *http.Request) {
	const op = "handler.deleteSubscription"
	log := h.log.With("op", op)
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	if err := h.repo.Delete(id); err != nil {
		log.Error("with h.repo.Delete ->failed to delete subscription", "error", err)
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) updateSubscription(w http.ResponseWriter, r *http.Request) {
	const op = "handler.updateSubscription"
	log := h.log.With("op", op)
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var sub models.SubscriptionForCreate
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Error("failed to read request body", "error", err)
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	if err := json.Unmarshal(body, &sub); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	if err := h.repo.Update(id, sub); err != nil {
		log.Error("with h.repo.Update -> failed to update subscription ", "error", err)
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) totalSubscriptions(w http.ResponseWriter, r *http.Request) {
	const op = "handler.totalSubscriptions"
	log := h.log.With("op", op)
	userID, err := validateUserID(r.URL.Query().Get("user_id"))
	if err != nil {
		// log.Error("with validateUserID -> failed to validate user_id", "error", err)
		http.Error(w, "bad userID attribute", http.StatusBadRequest)
		return
	}
	serviceName := r.URL.Query().Get("service_name")
	startDate, err := validateStartDate(r.URL.Query().Get("start_date"))
	if err != nil {
		// log.Error("with validateStartDate -> failed to validate start_date", "error", err)
		http.Error(w, "bad startDate attribute", http.StatusBadRequest)
		return
	}
	total, err := h.repo.Total(userID, serviceName, startDate)
	if err != nil {
		log.Error("with h.repo.Total -> failed to get total subscriptions", "error", err)
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}
	message, err := json.Marshal(total)
	if err != nil {
		log.Error("failed to marshal response", "error", err)
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(message)
}


 