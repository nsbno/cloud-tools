package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

func ParseDefaultCloudConfig() CloudConfig {
	return ParseConfig("cloud-config.yml")
}

func ParseConfig(filename string) CloudConfig {

	dir, _ := os.Getwd()
	absFilename, _ := filepath.Abs(fmt.Sprintf("%s%c%s", dir, os.PathSeparator, filename))
	yamlFile, err := ioutil.ReadFile(absFilename)

	if err != nil {
		panic(err)
	}

	var config CloudConfig

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}

	return config

}

type SecretVariable struct {
	Name string
	Key  string
}

type Variable struct {
	Name  string
	Value string
}

type CloudConfig struct {
	SecretVariables  []SecretVariable `secret-vars`
	Variables        []Variable       `vars`
	TerraformVersion string           `terraform-version`
}
