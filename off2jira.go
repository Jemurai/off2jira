package main

import (
	"os"

	cmd "github.com/jemurai/off2jira/cmd"
)

func main() {
	if len(os.Args) < 2 {
		panic("No file of OFF findings provided.")
	}
	var offfile string = os.Args[1]
	if offfile == "" {
		panic("No file of OFF findings provided")
	}
	// fmt.Printf("Arg = %s", offfile)
	cmd.Report(offfile)
}
