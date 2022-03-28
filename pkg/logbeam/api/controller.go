package api

import (
	"net/http"

	"github.com/barklan/logbeam/pkg/logbeam/config"
	"github.com/go-chi/render"
	"go.uber.org/zap"
)

type Controller struct {
	log  *zap.Logger
	conf *config.Config
}

func NewController(lg *zap.Logger, conf *config.Config) *Controller {
	return &Controller{
		log:  lg,
		conf: conf,
	}
}

type ErrorResp struct {
	Msg string `json:"msg"`
}

func (e ErrorResp) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (c *Controller) errorResp(w http.ResponseWriter, r *http.Request, msg string, code int) {
	w.WriteHeader(code)
	if render.Render(w, r, &ErrorResp{Msg: msg}) != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (c *Controller) internalError(w http.ResponseWriter, r *http.Request, msg string, err error) {
	c.log.Error(msg, zap.Error(err))
	c.errorResp(w, r, msg, http.StatusInternalServerError)
}

// func AllowCors(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Access-Control-Allow-Origin", "*")
// 		next.ServeHTTP(w, r)
// 	})
// }

// func (c *Controller) Serve() error {
// 	r := chi.NewRouter()
// 	r.Use(middleware.Logger)
// 	r.Use(render.SetContentType(render.ContentTypeJSON))
// 	// r.Use(AllowCors)
// 	r.Route("/api", func(r chi.Router) {
// 		r.Route("/test", func(r chi.Router) {
// 			r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
// 				c.helloHandler(w, r)
// 			})
// 		})
// 	})

// 	c.log.Info("logbeam rest server is listening", zap.Int64("port", port))

// 	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), r); err != nil {
// 		return fmt.Errorf("failed to listen and serve: %w", err)
// 	}

// 	return nil
// }
