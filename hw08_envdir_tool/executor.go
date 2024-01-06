package main

import (
	"errors"
	"os"
	"os/exec"
	"strings"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	command := exec.Command(cmd[0], cmd[1:]...)
	command.Env = prepareEnvironments(env)
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err := command.Run()
	if err != nil {
		exitError := new(exec.ExitError)
		if errors.As(err, &exitError) {
			return exitError.ExitCode()
		}
		return errorReturnCode
	}

	return successReturnCode
}

func prepareEnvironments(environments Environment) []string {
	systemEnvironments := os.Environ()
	for environmentName, value := range environments {
		systemEnvironments = changeEnvironments(systemEnvironments, environmentName)
		if !value.NeedRemove {
			systemEnvironments = append(systemEnvironments, environmentName+"="+value.Value)
		}
	}

	return systemEnvironments
}

func changeEnvironments(systemEnvironments []string, environmentName string) []string {
	number := 0
	for _, systemEnvironment := range systemEnvironments {
		systemEnvironmentName := strings.SplitN(systemEnvironment, "=", 2)[0]
		if systemEnvironmentName != environmentName {
			systemEnvironments[number] = systemEnvironment
			number++
		}
	}

	return systemEnvironments[:number]
}
