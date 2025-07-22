package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	model "test-assignment/internal/app/models"
	"time"

	"github.com/go-chi/chi/v5"
)

// @Summary Создать подписку
// @Description Создает новую подписку для пользователя
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param subscription body requests.CreateSubscriptionRequest true "Данные подписки"
// @Success 201 {object} model.Subscription "Созданная подписка"
// @Failure 400 {object} response.ErrorResponse "Некорректные данные"
// @Failure 500 {object} response.ErrorResponse "Внутренняя ошибка сервера"
// @Router /subscriptions [post]
func (h *SubscriptionHandler) Create(w http.ResponseWriter, r *http.Request) {
	var sub model.Subscription
	if err := json.NewDecoder(r.Body).Decode(&sub); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.Svc.Create(&sub); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(sub)
}

// @Summary Получить список подписок
// @Description Возвращает список всех подписок в системе
// @Tags subscriptions
// @Produce json
// @Success 200 {array} model.Subscription "Список подписок"
// @Failure 500 {object} response.ErrorResponse "Внутренняя ошибка сервера"
// @Router /subscriptions [get]
func (h *SubscriptionHandler) List(w http.ResponseWriter, r *http.Request) {
	subs, err := h.Svc.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(subs)
}

// @Summary Получить подписку по ID
// @Description Возвращает подписку по указанному идентификатору
// @Tags subscriptions
// @Produce json
// @Param id path int true "ID подписки"
// @Success 200 {object} model.Subscription "Подписка"
// @Failure 400 {object} response.ErrorResponse "Некорректный ID"
// @Failure 404 {object} response.ErrorResponse "Подписка не найдена"
// @Router /subscriptions/{id} [get]
func (h *SubscriptionHandler) Get(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}

	sub, err := h.Svc.GetByID(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(sub)
}

// @Summary Обновить подписку
// @Description Обновляет существующую подписку по ID
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param id path int true "ID подписки"
// @Param subscription body requests.UpdateSubscriptionRequest true "Обновленные данные подписки"
// @Success 200 {object} model.Subscription "Обновленная подписка"
// @Failure 400 {object} response.ErrorResponse "Некорректные данные"
// @Failure 500 {object} response.ErrorResponse "Внутренняя ошибка сервера"
// @Router /subscriptions/{id} [put]
func (h *SubscriptionHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}

	var sub model.Subscription
	if err := json.NewDecoder(r.Body).Decode(&sub); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	sub.ID = uint(id)

	if err := h.Svc.Update(&sub); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(sub)
}

// @Summary Удалить подписку
// @Description Удаляет подписку по указанному ID
// @Tags subscriptions
// @Param id path int true "ID подписки"
// @Success 204 "Подписка успешно удалена"
// @Failure 400 {object} response.ErrorResponse "Некорректный ID"
// @Failure 500 {object} response.ErrorResponse "Внутренняя ошибка сервера"
// @Router /subscriptions/{id} [delete]
func (h *SubscriptionHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid ID", http.StatusBadRequest)
		return
	}

	if err := h.Svc.Delete(uint(id)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// @Summary Получить общую стоимость подписок
// @Description Вычисляет общую стоимость подписок по заданным критериям за период
// @Produce json
// @Param user_id query string false "ID пользователя" example:"550e8400-e29b-41d4-a716-446655440000"
// @Param service_name query string false "Название сервиса" example:"Netflix"
// @Param from query string true "Дата начала периода в формате MM-YYYY" example:"01-2024"
// @Param to query string true "Дата окончания периода в формате MM-YYYY" example:"12-2024"
// @Success 200 {object} response.TotalCostResponse "Общая стоимость"
// @Failure 400 {object} response.ErrorResponse "Некорректные параметры запроса"
// @Failure 500 {object} response.ErrorResponse "Внутренняя ошибка сервера"
// @Router /total [get]
func (h *SubscriptionHandler) TotalCost(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	serviceName := r.URL.Query().Get("service_name")
	fromStr := r.URL.Query().Get("from")
	toStr := r.URL.Query().Get("to")

	layout := "01-2025"
	from, err := time.Parse(layout, fromStr)
	if err != nil {
		http.Error(w, "invalid from date", http.StatusBadRequest)
		return
	}

	to, err := time.Parse(layout, toStr)
	if err != nil {
		http.Error(w, "invalid to date", http.StatusBadRequest)
		return
	}

	total, err := h.Svc.TotalCost(userID, serviceName, from, to)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]int{"total": total})
}
