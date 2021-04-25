// Go practice app testing various interactions with common dev tooling
// such as Git, Bitbucket, Jenkins, Jira
package main

import (
	"os"

	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
	"github.com/teng1/go-bot/cmd"
	"gopkg.in/yaml.v2"
)

var help = `
  Usage: go-bot [command] [--help]

  Commands:
    --docker-project - Create a template docker project, CI/CD pipeline, Jira Tickets and PR's. 
    --helm-project - Create a template Helm chart project, CI/CD pipeline, Jira Tickets and PR's
  
  Read more:
    https://github.com/teng1/go-bot


`
var dockerProject = `
  --create [name], Creates a docker project and git repository with specified name. This
  will also create a jira ticket, template Dockerfile and some documentation. PR's will 
  be raised with our team to enable the CI/CD build pipeline and allow publishing to 
  container registires. 


`
var helmProject = `
  --create [name], Creates a template Helm project and git repository. This will also create
  a Jira ticket, template Helm chart and some documentation to get started. PR's will be raised
  with our team to enable the CI/CD build pipelines and enable publishing to chart repositories
  

`

// Default constants
// const (
// 	defaultGitURL        = "https://github.com"
// 	defaultGitPort       = 22
// 	defaultGitRepo       = "/teng1/gobot"
// 	defaultGitSSHKeyPath = "~/.ssh/id_rsa"
// 	defaultJiraURL       = "https://issues.apache.org/jira/"
// 	defaultJiraToken     = ""
// 	defaultTimeout       = 3 * time.Second
// 	defaultWriteTimeout  = time.Duration(0) //wite() will not timeout
// )

type Config struct {
	GoBotSettings struct {
		// Timeout       time.Duration `yaml:"timeout"`
		// WriteTimeout  time.Duration `yaml:"write_timeout"`
		GitConfig struct {
			Host           string `yaml:"host" envconfig:"GOBOT_GIT_HOST"`
			Port           int    `yaml:"port" envconfig:"GOBOT_GIT_PORT"`
			SshKeyLocation string `yaml:"sshKeyLocation" envconfig:"GOBOT_GIT_SSH_KEY"`
		} `yaml:"gitConfig"`
		JiraConfig struct {
			Host     string `yaml:"host" envconfig:"GOBOT_JIRA_HOST"`
			Port     int    `yaml:"port" envconfig:"GOBOT_JIRA_PORT"`
			ApiToken string `yaml:"apiToken" envconfig:"GOBOT_JIRA_TOKEN"`
		} `yaml:"jiraConfig"`
	} `yaml:"goBotSettings"`
}

func processError(err error) {
	log.Error(err)
	os.Exit(2)
}

// Read configfile
func readFile(cfg *Config) {
	f, err := os.Open("config.yml")
	if err != nil {
		processError(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		processError(err)
	}
}

// Read environment variables
func readEnv(cfg *Config) {
	err := envconfig.Process("", cfg)
	if err != nil {
		processError(err)
	}
}

// Read command line arguments

func main() {
	var cfg Config
	readFile(&cfg)
	readEnv(&cfg)
	log.Info("%v", cfg)
	cmd.Demo()
	cmd.InMemClone()
}

// func newWithDialer(config Config)  {
// 	if config.FluentNetwork == "" {
// 		config.FluentNetwork = defaultNetwork
// 	}
// 	if config.FluentHost == "" {
// 		config.FluentHost = defaultHost
// 	}
// 	if config.FluentPort == 0 {
// 		config.FluentPort = defaultPort
// 	}
// 	if config.FluentSocketPath == "" {
// 		config.FluentSocketPath = defaultSocketPath
// 	}
// 	if config.WriteTimeout == 0 {
// 		config.WriteTimeout = defaultWriteTimeout
// 	}
// 	if config.BufferLimit == 0 {
// 		config.BufferLimit = defaultBufferLimit
// 	}
// 	if config.RetryWait == 0 {
// 		config.RetryWait = defaultRetryWait
// 	}
// 	if config.MaxRetry == 0 {
// 		config.MaxRetry = defaultMaxRetry
// 	}
// 	if config.MaxRetryWait == 0 {
// 		config.MaxRetryWait = defaultMaxRetryWait
// 	}
// 	if config.AsyncConnect {
// 		fmt.Fprintf(os.Stderr, "fluent#New: AsyncConnect is now deprecated, please use Async instead")
// 		config.Async = config.Async || config.AsyncConnect
// 	}
