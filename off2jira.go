package main

import (
	"os"

	cmd "github.com/jemurai/off2jira/cmd"
)

func main() {
	var offfile string = os.Args[1]
	cmd.Report(offfile)
}
