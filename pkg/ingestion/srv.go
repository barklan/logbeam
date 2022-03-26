package ingestion

import (
	"fmt"
	"net/http"

	"github.com/barklan/logdip/pkg/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"go.uber.org/zap"
)

const port = 9010

type PublicCtrl struct {
	log  *zap.Logger
	conf *config.Config
}

func (c *PublicCtrl) internalError(w http.ResponseWriter, msg string, err error) {
	c.log.Error(msg, zap.Error(err))
	http.Error(w, msg, http.StatusInternalServerError)
}

func NewCtrl(lg *zap.Logger, conf *config.Config) *PublicCtrl {
	return &PublicCtrl{
		log:  lg,
		conf: conf,
	}
}

func AllowCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}

func (c *PublicCtrl) Serve() error {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Use(AllowCors)
	r.Route("/api/logdip", func(r chi.Router) {
		r.Route("/test", func(r chi.Router) {
			r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
				c.helloHandler(w, r)
			})
		})
	})

	c.log.Info("logdip rest server is listening", zap.Int64("port", port))

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), r); err != nil {
		return fmt.Errorf("failed to listen and serve: %w", err)
	}

	return nil
}
