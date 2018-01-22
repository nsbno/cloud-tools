package wrapper

import (
	"github.com/nsbno/cloud-tools/config"
	"os"
	"os/exec"
        "io/ioutil"
        "fmt"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

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

func RunCmds(commands []config.Command) {
	for _, command := range commands {
                var argumentsInterface []interface{} = make([]interface{}, len(command.Arguments))
                for i, d := range command.Arguments {
	            argumentsInterface[i] = d
                }
                fmt.Printf("Running %s\n", command.Executable)
                fmt.Printf("With arguments %v\n", argumentsInterface)
                fmt.Println("Writing to %s\n", command.Outputfile)
                out, err := exec.Command(command.Executable, command.Arguments...).Output()
                fmt.Printf("%s\n", out)
		check(err)
                if command.Outputfile == "" {
			fmt.Printf("%s\n", out)
                }
		if out != nil {
			ioutil.WriteFile(command.Outputfile, []byte(out), 0644)
		}
	}
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
