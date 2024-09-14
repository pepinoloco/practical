package handlers

import (
	"net/http"

	"practical/appendix-b/http-server/config"
	"practical/appendix-b/http-server/types"
)

func SetupHandlers(mux *http.ServeMux, config *config.AppConfig) {
	mux.Handle(
		"/api/packages",
		&types.App{Config: config, Handler: packageHandler},
	)
}
