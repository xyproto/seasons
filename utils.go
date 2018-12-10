package main

import (
	"os"
	"os/exec"
	"strings"
)

func which(executable string) string {
	p, err := exec.LookPath(executable)
	if err != nil {
		return ""
	}
	return p
}

func run(shellCommand string) error {
	cmd := exec.Command("sh", "-c", shellCommand)
	if _, err := cmd.CombinedOutput(); err != nil {
		return err
	}
	return nil
}

func output(shellCommand string) string {
	cmd := exec.Command("sh", "-c", shellCommand)
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		return ""
	}
	return string(stdoutStderr)
}

func containsE(envVar, subString string) bool {
	return strings.Contains(os.Getenv(envVar), subString)
}

func hasE(envVar string) bool {
	return os.Getenv(envVar) != ""
}
