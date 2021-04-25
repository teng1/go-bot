// Go practice app testing various interactions with common dev tooling
// such as Git, Bitbucket, Jenkins, Jira
package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
	"github.com/teng1/go-bot/cmd"
	"gopkg.in/yaml.v2"
)

var help = `
  Usage: go-bot [command] [--help]

  Commands:
    docker-repo - Create a template docker project, CI/CD pipeline, Jira Tickets and PR's. 
    helm-repo - Create a template Helm chart project, CI/CD pipeline, Jira Tickets and PR's
  
  Read more:
    https://github.com/teng1/go-bot


`
var dockerHelp = `
  --create [name], Creates a docker project and git repository with specified name. This
  will also create a jira ticket, template Dockerfile and some documentation. PR's will 
  be raised with our team to enable the CI/CD build pipeline and allow publishing to 
  container registires. 


`
var helmHelp = `
  --create [name], Creates a template Helm project and git repository. This will also create
  a Jira ticket, template Helm chart and some documentation to get started. PR's will be raised
  with our team to enable the CI/CD build pipelines and enable publishing to chart repositories
  

`

// Default constants
const (
	defaultTimeout      = 3 * time.Second
	defaultWriteTimeout = time.Duration(0) //wite() will not timeout
)

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
// env vars take precedence over config.yml
func readEnv(cfg *Config) {
	err := envconfig.Process("", cfg)
	if err != nil {
		processError(err)
	}
}

// Args branch, flags set params
func main() {

	version := flag.Bool("version", false, "")
	v := flag.Bool("v", false, "")
	flag.Bool("help", false, "")
	flag.Bool("h", false, "")
	flag.Usage = func() {}
	flag.Parse()

	if *version || *v {
		// TODO: Detect version
		fmt.Println("0.0.1")
		os.Exit(0)
	}

	args := flag.Args()

	subcmd := ""
	if len(args) > 0 {
		subcmd = args[0]
		args = args[1:]
	}

	switch subcmd {
	case "docker-project":
		dockerRepo(args)
	// case "helm":
	// 	client(args)
	default:
		fmt.Print(help)
		os.Exit(0)
	}
}

func dockerRepo(args []string) {

	// dockerProj := flag.String("docker-project", "", dockerProject)
	// helmProj := flag.String("helm-project", "", helmProject)

	// flag.Usage = func() {
	// 	fmt.Printf(help, os.Args[0])
	// 	flag.PrintDefaults()
	// }

	// var cfg Config
	// readFile(&cfg)
	// readEnv(&cfg)

	// flag.Parse()

	// log.Info(*dockerProj)
	// log.Info(*helmProj)

	// log.Info("%v", cfg)
	cmd.Demo()
	cmd.InMemClone()
}

// 	if config.FluentPort == 0 {
// 		config.FluentPort = defaultPort
// 	}
