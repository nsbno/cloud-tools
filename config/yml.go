package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"errors"
)

func ParseDefaultCloudConfig() (CloudConfig, error) {
	return ParseConfig("cloud-config.yml")
}

func ParseConfig(filename string) (CloudConfig, error) {

	dir, _ := os.Getwd()
	absFilename, _ := filepath.Abs(fmt.Sprintf("%s%c%s", dir, os.PathSeparator, filename))
	yamlFile, err := ioutil.ReadFile(absFilename)

	var config CloudConfig

	if err != nil {
		return config, errors.New(fmt.Sprintf("Missing %s. Make sure you run this command from a Terraform environment directory.\n", filename))
	}


	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return config, errors.New(fmt.Sprintf("Unable to parse %s. Make sure it contains valid YAML syntax", filename))
	}

	return config, nil

}

type SecretVariable struct {
	Name string
	Key  string
}

type Variable struct {
	Name  string
	Value string
}

type Command struct {
        Executable string
        Arguments  []string
        Outputfile string
}

type CloudConfig struct {
	SecretVariables  []SecretVariable `secret-vars`
	Variables        []Variable       `vars`
        Commands         []Command        `commands`
	TerraformVersion string           `terraform-version`
}
