package entity

type key string

const (
	// AppCtx global context
	AppCtx key = "app"
	// ConfigCtx context name for config
	ConfigCtx key = "config"
	// DBCtx context name for database adapter
	DBCtx key = "db"
)