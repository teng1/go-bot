// Go practice app testing various interactions with common dev tooling
// such as Git, Bitbucket, Jenkins, Jira
package main

import (
	"github.com/teng1/go-bot/cmd"
)

func main() {
	cmd.Demo()
	cmd.InMemClone()
}
