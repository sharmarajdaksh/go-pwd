package config

import (
	"github.com/kkyr/fig"
)

type config struct {
	DatabaseFile string `fig:"DatabaseFile" default:"go-pwd.db"`
	MasterSecret string `fig:"MasterSecret" default:"QAfsw8@dQbz@K#EdTSZandLz6W*Wt61G"`
}

// C represents a global config object
var c config

// LoadConfig loads up the global config struct from file on startup
func LoadConfig() error {
	// TODO: config file location
	return fig.Load(&c,
		fig.File("./config.yaml"),
	)
}

// GetDatabaseFile returns the name of the sqlite db file from the config file
func GetDatabaseFile() string {
	return c.DatabaseFile
}

// GetMasterSecret returns the master secret from the config file
func GetMasterSecret() string {
	return c.MasterSecret
}
