package config

import (
	"flag"
	"fmt"

	"github.com/jinzhu/configor"
)

// Config represents the composition of yml settings.
type Config struct {
	Database struct {
		Dialect   string `default:"sqlite3"`
		Host      string `default:"book.db"`
		Port      string
		Dbname    string
		Username  string
		Password  string
		Migration bool `default:"false"`
	}
	Extension struct {
		MasterGenerator bool `yaml:"master_generator" default:"false"`
		CorsEnabled     bool `yaml:"cors_enabled" default:"false"`
	}
}

const (
	// DEV represents development environment
	DEV = "develop"
	// PRD represents production environment
	PRD = "production"
	// DOC represents docker container
	DOC = "docker"
)

var config *Config
var env *string

// Load reads the settings written to the yml file
func Load() {
	env = flag.String("env", "develop", "To switch configurations.")
	flag.Parse()
	config = &Config{}
	configor.Load(config, "application."+*env+".yml")
	fmt.Println("Loaded this configuration : " + "application." + *env + ".yml")
	fmt.Println(*config)
}

// GetConfig returns the configuration data.
func GetConfig() *Config {
	return config
}

// GetEnv returns the environment variable.
func GetEnv() *string {
	return env
}