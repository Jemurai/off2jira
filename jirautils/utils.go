package utils

import (
	"fmt"
	"os"

	"github.com/jemurai/fkit/finding"

	jira "gopkg.in/andygrunwald/go-jira.v1"
)

var server string = os.Getenv("AUTOM8D_JIRA_SERVER")
var user string = os.Getenv("AUTOM8D_JIRA_EMAIL")
var token string = os.Getenv("AUTOM8D_JIRA_TOKEN")
var project string = os.Getenv("AUTOM8D_JIRA_PROJECT")
var issueType string = os.Getenv("AUTOM8D_JIRA_ISSUE_TYPE")

var jiraClient *jira.Client

func init() {
	if server == "" {
		panic("No JIRA server specified in ENV AUTOM8D_JIRA_SERVER")
	}
	if user == "" {
		panic("No user specified in ENV AUTOM8D_JIRA_USER")
	}
	if token == "" {
		panic("No API token specified in ENV AUTOM8D_JIRA_API_TOKEN")
	}
	if project == "" {
		panic("No project specified in ENV AUTOM8D_JIRA_PROJECT")
	}
	if issueType == "" {
		issueType = "Bug"
	}
	tp := jira.BasicAuthTransport{
		Username: user,
		Password: token,
	}
	var err error
	jiraClient, err = jira.NewClient(tp.Client(), server)
	if err != nil {
		panic(err)
	}
}

// ReportFindings - report a set of findings to JIRA
func ReportFindings(findings []finding.Finding) {
	for i := 0; i < len(findings); i++ {
		finding := findings[i]
		if !CheckIfFindingExists(finding) {
			CreateIssue(finding)
		}
	}
}

// CheckIfFindingExists - check if a finding is already in JIRA
func CheckIfFindingExists(finding finding.Finding) bool {
	opt := &jira.SearchOptions{
		MaxResults: 3,
	}
	searchString := "project=" + project + " and description ~ \"Fingerprint: " + finding.Fingerprint + "\""
	_, resp, err := jiraClient.Issue.Search(searchString, opt)
	if err != nil {
		fmt.Printf("Error searching for issues with: %s", searchString)
		panic(err)
	}

	total := resp.Total
	return total > 0
}

// Create a JIRA Issue
func CreateIssue(f finding.Finding) {
	summary := f.Name + " - " + f.Description
	var labels []string
	labels = append(labels, "AUTOM8D", f.Source)
	i := jira.Issue{
		Fields: &jira.IssueFields{
			Description: finding.GetDetailString(f),
			Type: jira.IssueType{
				Name: issueType,
			},
			Project: jira.Project{
				Key: project,
			},
			Summary: summary,
			Labels:  labels,
		},
	}

	_, _, err := jiraClient.Issue.Create(&i)
	if err != nil {
		panic(err)
	}
}
