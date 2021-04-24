package cmd

import (
	"fmt"

	"github.com/andygrunwald/go-jira"
	log "github.com/sirupsen/logrus"
)

// Create ticket if not exist, return ticket number,
// Demo This is just the first example from the go-jira readme
func Demo() {
	jiraClient, _ := jira.NewClient(nil, "https://issues.apache.org/jira/")
	issue, _, _ := jiraClient.Issue.Get("MESOS-3325", nil)

	fmt.Printf("%s: %+v\n", issue.Key, issue.Fields.Summary)
	fmt.Printf("Type: %s\n", issue.Fields.Type.Name)
	fmt.Printf("Priority: %s\n", issue.Fields.Priority.Name)

	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")
}

func CreateTicket() {

}

func AssignToBot() {

}
