package handler

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"tt2/internal/models"
	"tt2/internal/repository"

	"github.com/gorilla/mux"
)

type CreateOutPut struct {
	ID int `json:"id" example:"123"`
}



// @Summary Создать новую подписку
// @Description Создает новую подписку и возвращает ID
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param subscription body models.SubscriptionForCreate true "Данные подписки"
// @Success 201 {object} CreateOutPut
// @Failure 400 {string} string "Неверный запрос"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /subscriptions [post]
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

// GetAllOutPut содержит список подписок
type GetAllOutPut struct {
	Subscriptions []models.Subscription `json:"subscriptions"`
}


// @Summary Получить все подписки
// @Description Возвращает список всех подписок
// @Tags subscriptions
// @Produce json
// @Success 200 {object} GetAllOutPut
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /subscriptions [get]
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


// @Summary Получить подписку по ID
// @Description Возвращает данные подписки по ID
// @Tags subscriptions
// @Produce json
// @Param id path int true "ID подписки"
// @Success 200 {object} models.Subscription
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Failure 404 {string} string "Подписка не найдена"
// @Router /subscriptions/{id} [get]
func (h *Handler) getSubscription(w http.ResponseWriter, r *http.Request) {
	const op = "handler.getSubscription"
	log := h.log.With("op", op)
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	sub, err := h.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			http.Error(w, "subscription not found", http.StatusNotFound)
			return
		}
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


// @Summary Удалить подписку по ID
// @Description Удаляет подписку по ID
// @Tags subscriptions
// @Param id path int true "ID подписки"
// @Success 204 {string} string "Прошло успешно"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /subscriptions/{id} [delete]
func (h *Handler) deleteSubscription(w http.ResponseWriter, r *http.Request) {
	const op = "handler.deleteSubscription"
	log := h.log.With("op", op)
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	if err := h.repo.Delete(id); err != nil {
		log.Error("with h.repo.Delete -> failed to delete subscription", "error", err)
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}


// @Summary Обновить подписку
// @Description Обновляет данные подписки по ID
// @Tags subscriptions
// @Accept json
// @Param id path int true "ID подписки"
// @Param subscription body models.SubscriptionForCreate true "Обновленные данные подписки"
// @Success 204 {string} string "Прошло успешно"
// @Failure 400 {string} string "Неверный запрос"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /subscriptions/{id} [put]
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


// @Summary Получить суммарную стоимость подписок
// @Description Считает суммарную стоимость подписок за период с фильтрами по пользователю, названию подписки и дате начала периода(все вместе или на выбор). При отсутствии фильтров считает полную сумму всех подписок
// @Tags subscriptions
// @Produce json
// @Param user_id query string false "ID пользователя" format(uuid)
// @Param service_name query string false "Название подписки"
// @Param start_date query string false "Дата начала периода" format(date)
// @Success 200 {integer} int "Суммарная стоимость"
// @Failure 400 {string} string "Неверный запрос"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /subscriptions/total [get]
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


 