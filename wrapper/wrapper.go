package wrapper

import (
	"github.com/nsbno/cloud-tools/config"
	"os"
	"os/exec"
)

// GetEnvironmentVariablesForSecrets appends env vars for secrets to environment
func GetEnvironmentVariablesForSecrets(secretVars []config.SecretVariable) []string {
	var environment []string
	for _, secretVar := range secretVars {
		environment = append(environment, secretVar.Name+"="+config.GetPasswordFor(secretVar.Key))
	}
	return environment
}

// GetEnvironmentVariablesForValues appends env vars for values to environment
func GetEnvironmentVariablesForValues(vars []config.Variable) []string {
	var environment []string
	for _, variable := range vars {
		environment = append(environment, variable.Name+"="+variable.Value)
	}
	return environment
}

// ExecuteCommand builds the command line string to be executed
func ExecuteCommand(command string, args []string, environment []string) {

	cmd := exec.Command(command, args...)

	cmd.Env = append(environment, "PATH="+os.Getenv("PATH"))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Start()
	defer cmd.Wait()
}
