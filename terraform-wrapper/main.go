package main

import (
	"fmt"
        "strings"
	"github.com/nsbno/cloud-tools/config"
	"github.com/nsbno/cloud-tools/wrapper"
	"os"
	"os/exec"
        "bufio"
	"time"
)

// terraform-wrapper will get secrets from your pass password store,
// setup an environment containing secrets and execute terraform,
// passing command-line arguments to terraform as-is
func main() {
        if !isInstalled("terraform") {
                fmt.Fprintf(os.Stderr, "Terraform is not installed!\n")
                os.Exit(1)
        }
        if !isInstalled("pass") {
                fmt.Fprintf(os.Stderr, "pass is not installed!\n")
                os.Exit(1)
        }

        terraformArgs := os.Args[1:]
        terraformArgs = append(terraformArgs, "-input=false")
        if (os.Args[1] == "apply" ) {
                terraformArgs = append(terraformArgs, "-auto-approve")
                terraformArgs = append(terraformArgs, "-lock=true")
        }

	config := config.ParseDefaultCloudConfig()

        // Always add AWS credentials
        var credentials []string
        credentials = append(credentials, "AWS_ACCESS_KEY_ID="+os.Getenv("AWS_ACCESS_KEY_ID"))
        credentials = append(credentials, "AWS_SECRET_ACCESS_KEY="+os.Getenv("AWS_SECRET_ACCESS_KEY"))

        env := wrapper.GetEnvironmentVariablesForValues(config.Variables[:])
	secEnv := wrapper.GetEnvironmentVariablesForSecrets(config.SecretVariables[:])
        secEnv = append(secEnv, credentials...)
        fullEnv := append(env, secEnv...)

        if contains(env, "TF_VAR_env=prod") && (os.Args[1] == "apply" ) {
                reader := bufio.NewReader(os.Stdin)
                fmt.Print("You are provisioning PROD. Type PROD to continue... ")
                input, _ := reader.ReadString('\n')
                if strings.TrimRight(input, "\n") != "PROD" {
                        os.Exit(0)
                }
        }

        start := time.Now()
        fmt.Println("Started terraform operation at:", start)
        wrapper.RunCmds(config.Commands[:])
	wrapper.ExecuteCommand("terraform", terraformArgs, fullEnv)
	stop := time.Now()
	fmt.Println("Started terraform operation at:", start)
	fmt.Println("Finished terraform operation at:", stop)
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

func contains(slice []string, item string) bool {
    set := make(map[string]struct{}, len(slice))
    for _, s := range slice {
        set[s] = struct{}{}
    }

    _, ok := set[item] 
    return ok
}
