package cmd

import (
	utils "github.com/jemurai/fkit/utils"
	jirautils "github.com/jemurai/off2jira/jirautils"
)

func Report(file string) {
	findings := utils.BuildFindingsFromFile(file)
	jirautils.ReportFindings(findings)
}
