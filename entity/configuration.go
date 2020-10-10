package entity

// Configuration for this application
type Configuration struct {
	Server ServerConfiguration
	Database DatabaseConfiguration
}

// DatabaseConfiguration for this application
type DatabaseConfiguration struct {
	ConnectionURI string
	Name string
}

// ServerConfiguration for this application
type ServerConfiguration struct {
	Port int
}