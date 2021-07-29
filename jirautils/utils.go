package utils

import (
	"fmt"

	"github.com/jemurai/fkit/finding"
)

func ReportFindings(findings []finding.Finding) {

	for i := 0; i < len(findings); i++ {
		fmt.Println(findings[i])
	}
}

func CheckIfFindingExists(finding finding.Finding) bool {
	return false
}
