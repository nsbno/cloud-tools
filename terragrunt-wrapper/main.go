package main

import (
	"fmt"
        "strings"
	"regexp"
	"github.com/nsbno/cloud-tools/config"
	"github.com/nsbno/cloud-tools/wrapper"
	"os"
	"os/exec"
	"time"
)

// terragrunt-wrapper will get secrets from your pass password store,
// setup an environment containing secrets and execute terragrunt,
// passing command-line arguments to terragrunt as-is
func main() {
        if !isInstalled("terraform") {
                fmt.Fprintf(os.Stderr, "Terraform is not installed!\n")
                os.Exit(1)
        }
        if !isInstalled("terragrunt") {
                fmt.Fprintf(os.Stderr, "Terragrunt is not installed!\n")
                os.Exit(1)
        }
        if !isInstalled("pass") {
                fmt.Fprintf(os.Stderr, "pass is not installed!\n")
                os.Exit(1)
        }

	config := config.ParseDefaultCloudConfig()

//	if !isRequiredTerraformVersion(config.TerraformVersion) {
//		fmt.Fprintf(os.Stderr, "Bad Terraform version - cloud-config.yml requires %s\n", config.TerraformVersion)
//		os.Exit(1)
//	}

	start := time.Now()
	fmt.Println("Started terragrunt operation at:", start)
	secEnv := wrapper.GetEnvironmentVariablesForSecrets(config.SecretVariables[:])
	env := wrapper.GetEnvironmentVariablesForValues(config.Variables[:])
        wrapper.RunCmds(config.Commands[:])
        fullEnv := append(env, secEnv...)
	wrapper.ExecuteCommand("terragrunt", os.Args[1:], fullEnv)
	stop := time.Now()
	fmt.Println("Started terragrunt operation at:", start)
	fmt.Println("Finished terragrunt operation at:", stop)
	duration := stop.Unix() - start.Unix()
	fmt.Println("Total duration (seconds): ", duration)
}

func isInstalled(execName string) bool {
	cmd := exec.Command(execName, "--version")
	output, _ := cmd.Output()
        totest :=  strings.ToLower(string(output))
	if strings.Contains(totest, execName) {
		return true
	}
	return false
}



func isRequiredTerraformVersion(requiredVersion string) bool {
	cmd := exec.Command("terraform", "--version")
	output, _ := cmd.Output()
	text := string(output)
	terraformVersion, _ := regexp.Compile(`Terraform v(\d{1,2}\.\d{1,2}\.\d{1,2})`)
	result_slice := terraformVersion.FindStringSubmatch(text)
	return result_slice[1] == requiredVersion
}
