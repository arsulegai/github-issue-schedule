package configs

import (
	"github-issue-schedule/internal/pkg/utils"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

// Configuration to run the program
type Configuration struct {
	BufferWindowDays uint8     `yaml:"buffer_window_days"`
	Projects         []Project `yaml:"projects"`
}

// Project related configuration parameters
type Project struct {
	Name        string     `yaml:"name"`
	GitHubOrg   string     `yaml:"github_org"`
	GitHubRepo  string     `yaml:"github_repo"`
	Maintainers string     `yaml:"maintainers"`
	Schedules   []Schedule `yaml:"schedules"`
}

// Schedule for sending emails
type Schedule struct {
	Date        string `yaml:"date"`
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
}

// ReadConfiguration returns the object by reading the file
func ReadConfiguration() Configuration {
	// expect an environment variable to know the location
	log.Println("reading the configuration")
	var config Configuration
	var configFile = utils.GetEnvOrDefault(ConfigFile, "config.yaml")
	fileContents, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatalf("Couldn't read the config file %v, Err: %v", configFile, err)
	}
	err = yaml.Unmarshal(fileContents, &config)
	if err != nil {
		log.Fatalf("Error while parsing the config file %v, Err: %v", configFile, err)
	}
	return config
}
