package version

import "fmt"

var (
	GitCommit = "unknown"
	BuildDate = "unknown"
	Version   = "unreleased"
)

func ShowVersion() {
	fmt.Printf("Version:\t %s\n", Version)
	fmt.Printf("Git commit:\t %s\n", GitCommit)
	fmt.Printf("Date:\t\t %s\n", BuildDate)
}
