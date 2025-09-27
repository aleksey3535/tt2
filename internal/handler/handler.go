package handler

import (
	"log/slog"
	"net/http"
	"time"
	"tt2/internal/config"
	"tt2/internal/middleware"
	"tt2/internal/models"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	
	httpswagger "github.com/swaggo/http-swagger"
)

type RepositoryI interface {
	Create(sub models.SubscriptionForCreate) (int, error)
	GetAll() ([]models.Subscription, error)
	GetByID(id int) (models.Subscription, error)
	Update(id int, sub models.SubscriptionForCreate) error
	Delete(id int) error
	Total(userID *uuid.UUID, serviceName string, startDate *time.Time) (int, error)
}

type Handler struct {
	cfg *config.Config
	log *slog.Logger
	mw  *middleware.Middleware
	repo RepositoryI
}

func New(log *slog.Logger, cfg *config.Config, mw *middleware.Middleware, repo RepositoryI) *Handler {
	return &Handler{
		cfg: cfg,
		log: log,
		mw:  mw,
		repo: repo,
	}
}

func (h *Handler) InitRoutes() *mux.Router {
	mux := mux.NewRouter()
	mux.Use(h.mw.UseHeaders)
	mux.HandleFunc("/subscriptions", h.createSubscription).Methods("POST")
	mux.HandleFunc("/subscriptions", h.getAllSubscriptions).Methods("GET")
	mux.HandleFunc("/subscriptions/{id:[0-9]+}", h.getSubscription).Methods("GET")
	mux.HandleFunc("/subscriptions/{id:[0-9]+}", h.deleteSubscription).Methods("DELETE")
	mux.HandleFunc("/subscriptions/{id:[0-9]+}", h.updateSubscription).Methods("PUT")
	mux.HandleFunc("/subscriptions/total", h.totalSubscriptions).Methods("GET")
	mux.PathPrefix("/swagger/").Handler(httpswagger.WrapHandler)
	mux.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error":"not found"}`))
	})
	return mux
}
