package entity

import (
	"context"
	"gorm.io/gorm"
	"net/http"
)

// Application configuration
type Application struct {
	Config Configuration
	DB *gorm.DB
}

// AppCtx set app context
func (app *Application) AppCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), AppCtx, app)
		next.ServeHTTP(w, r.WithContext(ctx))
  })
}