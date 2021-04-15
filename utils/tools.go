package utils

import "os"

// InCI return whether in github acitons
func InCI() bool {
	// https://help.github.com/en/actions/configuring-and-managing-workflows/using-environment-variables
	e := os.Getenv("GITHUB_ACTION")
	if e == "" {
		return false
	}
	return true
}

// IsDev ...
func IsDev() bool {
	return *modeFlag == "dev"
}
