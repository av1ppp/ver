package main

import (
	"errors"
	"os"
)

const (
	actionInit = "init"
	actionIncr = "incr"
	actionGet  = "get"
	actionHelp = "help"
)

const (
	appName = "ver"
	appDesc = "Versioning tool"
)

const (
	versionFile = ".version"
)

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		printHelp()
		return
	}

	if err := handleAction(args[0]); err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}
}

func handleAction(action string) error {
	switch action {
	case actionInit:
		return handleActionInit()
	case actionIncr:
		return handleActionIncr()
	case actionGet:
		return handleActionGet()
	case actionHelp:
		printHelp()
		return nil
	}
	panic("Unknown action '" + action + "'")
}

func handleActionInit() error {
	_, err := os.Stat(versionFile)
	if err == nil {
		return errors.New("Version file '" + versionFile + "' already exists")
	}

	if !os.IsNotExist(err) {
		return err
	}

	version := CreateVersionStruct(1)
	return version.WriteToFile(versionFile)
}

func handleActionIncr() error {
	oldVersion, err := ParseVersionStructFromString(versionFile)
	if err != nil {
		if os.IsNotExist(err) {
			return errors.New("Version file '" + versionFile + "' does not exist. Run '" + appName + " " + actionInit + "' first")
		}
		return err
	}

	newVersion := CreateVersionStruct(0)

	if oldVersion.year == newVersion.year && oldVersion.dayOfYear == newVersion.dayOfYear {
		newVersion.micro = oldVersion.micro + 1
	} else {
		newVersion.micro = 1
	}

	return newVersion.WriteToFile(versionFile)
}

func handleActionGet() error {
	version, err := ParseVersionStructFromString(versionFile)
	if err != nil {
		if os.IsNotExist(err) {
			return errors.New("Version file '" + versionFile + "' does not exist. Run '" + appName + " " + actionInit + "' first")
		}
		return err
	}
	os.Stdout.WriteString(version.String())
	return nil
}

func printHelp() {
	s := appName + " - " + appDesc + "\n\n"
	s += "Usage:\n"
	s += "  " + appName + " <action>\n\n"
	s += "Actions:\n"
	s += "  init - Initialize ver in current directory\n"
	s += "  incr - Increment version\n"
	s += "  get  - Get current version\n"
	s += "  help - Print this message\n"
	os.Stdout.WriteString(s + "\n")
}
