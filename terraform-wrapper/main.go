package main

import (
	"fmt"
	"github.com/nsbno/cloud-tools/config"
	"github.com/nsbno/cloud-tools/wrapper"
	"os"
	"os/exec"
	"regexp"
	"time"
)

// terraform-wrapper will get secrets from your pass password store,
// setup an environment containing secrets and execute terraform,
// passing command-line arguments to terraform as-is
func main() {
	config := config.ParseDefaultCloudConfig()

	if !isRequiredTerraformVersion(config.TerraformVersion) {
		fmt.Fprintf(os.Stderr, "Bad Terraform version - cloud-config.yml requires %s\n", config.TerraformVersion)
		os.Exit(1)
	}

	start := time.Now()
	fmt.Println("Started terraform operation at:", start)
	secEnv := wrapper.GetEnvironmentVariablesForSecrets(config.SecretVariables[:])
	env := wrapper.GetEnvironmentVariablesForValues(config.Variables[:])
	wrapper.ExecuteCommand("terraform", os.Args[1:], append(secEnv, env...))
	stop := time.Now()
	fmt.Println("Started terraform operation at:", start)
	fmt.Println("Finished terraform operation at:", stop)
	duration := stop.Unix() - start.Unix()
	fmt.Println("Total duration (seconds): ", duration)
}

func isRequiredTerraformVersion(requiredVersion string) bool {
	cmd := exec.Command("terraform", "--version")
	output, _ := cmd.Output()
	text := string(output)
	terraformVersion, _ := regexp.Compile(`Terraform v(\d{1,2}\.\d{1,2}\.\d{1,2})`)
	result_slice := terraformVersion.FindStringSubmatch(text)
	return result_slice[1] == requiredVersion
}
