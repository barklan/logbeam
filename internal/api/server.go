package api

import (
	"net/http"

	"github.com/barklan/logbeam/internal/config"
	"github.com/barklan/logbeam/internal/ingestion"
	"go.uber.org/zap"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

const applicationNDJSON = "application/x-ndjson"

type Server struct {
	Router *chi.Mux
	log    *zap.Logger
	conf   *config.Config
	agent  ingestion.Agent
}

func NewServer(logger *zap.Logger, conf *config.Config, agent ingestion.Agent) *Server {
	return &Server{
		Router: chi.NewRouter(),
		log:    logger,
		conf:   conf,
		agent:  agent,
	}
}

type Error struct {
	Message string `json:"message"`
}

type ErrorResp struct {
	Error Error
}

func (e ErrorResp) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *Server) mountHandlers() {
	s.Router.Use(middleware.Logger)
	s.Router.Use(allowCors)
	s.Router.Use(middleware.Recoverer)

	s.Router.Post("/input", s.input)
	s.Router.Route("/input", func(r chi.Router) {
		r.Use(middleware.BasicAuth("logbeam", map[string]string{
			s.conf.Username: s.conf.Password,
		}))
		r.Use(middleware.AllowContentType(applicationNDJSON))
		r.Post("/", s.input)
	})
	s.Router.Get("/auth/token", s.getAuthToken)
}

func (s *Server) errorResp(w http.ResponseWriter, r *http.Request, message string, httpCode int) {
	w.WriteHeader(httpCode)
	body := ErrorResp{
		Error{
			Message: message,
		},
	}
	if render.Render(w, r, &body) != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (s *Server) internalError(w http.ResponseWriter, r *http.Request, message string, err error) {
	s.log.Error(message, zap.Error(err))
	s.errorResp(w, r, message, http.StatusInternalServerError)
}

func allowCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}
