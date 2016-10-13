package config

import (
	"bytes"
	"io"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

// GetPasswordFor will lookup key in pass - the linux password store
// and ask for GPG password
func GetPasswordFor(key string) string {
	return GetPasswordFromPasswordStoreFor(key, "")
}

// GetPasswordFor will lookup key in pass - the linux password store
// and ask for GPG password
// Will use the password store dir provided
func GetPasswordFromPasswordStoreFor(key string, passwordStorageDirectory string) string {

	cmd := exec.Command("pass", key)

	if passwordStorageDirectory != "" {
		cmd.Env = isolatedPassEnvironment(passwordStorageDirectory)
	}

	// Ask for gpg password if necessary
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	output, _ := cmd.Output()
	return strings.TrimSpace(string(output))

}

// GeneratePasswordFor will generate password of given length in given password storage dir
func GeneratePasswordFor(passwordStorageDirectory string, passName string, passLength int) string {

	cmd := exec.Command("pass", "generate", "-f", "-n", passName, strconv.Itoa(passLength))

	cmd.Env = isolatedPassEnvironment(passwordStorageDirectory)
	// Enable when debugging
	//cmd.Stdout = os.Stdout
	//cmd.Stderr = os.Stderr

	output, _ := cmd.Output()

	outputString := string(output)

	// After removing removing the ctrl bytes, this is all that's left of the colour codes in the string
	prefix := regexp.QuoteMeta("[1m[37mThe generated password for [4m" + passName + "[24m is:[0m[1m[93m")
	suffix := regexp.QuoteMeta("[0m")
	r := regexp.MustCompile(prefix + "(.*)" + suffix)

	return r.FindStringSubmatch(stripCtlAndExtFromBytes(outputString))[1]

}

func stripCtlAndExtFromBytes(str string) string {
	b := make([]byte, len(str))
	var bl int
	for i := 0; i < len(str); i++ {
		c := str[i]
		if c >= 32 && c < 127 {
			b[bl] = c
			bl++
		}
	}
	return string(b[:bl])
}

// InsertPasswordFor will insert password pass store in given storage dir
func InsertPasswordFor(passwordStorageDirectory string, passName string, password string) {

	cmd := exec.Command("pass", "insert", "-e", passName)
	cmd.Env = isolatedPassEnvironment(passwordStorageDirectory)
	stdin, _ := cmd.StdinPipe()
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Start()
	defer cmd.Wait()
	io.Copy(stdin, bytes.NewBufferString(password))
	defer stdin.Close()
}

func DeletePasswordsRecursivelyFor(passwordStorageDirectory string, passName string) {
	cmd := exec.Command("pass", "rm", "-r", "-f", passName)
	cmd.Env = isolatedPassEnvironment(passwordStorageDirectory)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Start()
	defer cmd.Wait()
}

// InitialisePasswordStore will initialise password store with given gpgKeyIds
func InitialisePasswordStore(passwordStorageDirectory string, passName string, gpgKeyIds ...string) {

	args := make([]string, 0)
	args = append(args, "init")
	args = append(args, "-p")
	args = append(args, passName)
	args = append(args, gpgKeyIds...)

	cmd := exec.Command("pass", args...)
	cmd.Env = isolatedPassEnvironment(passwordStorageDirectory)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Start()
	defer cmd.Wait()
}

func isolatedPassEnvironment(passwordStorageDirectory string) []string {
	var environment []string
	environment = append(environment,
		"PATH="+os.Getenv("PATH"),
		"HOME="+os.Getenv("HOME"),
		"LANG="+os.Getenv("LANG"),
		"PASSWORD_STORE_DIR="+passwordStorageDirectory,
		"PASSWORD_STORE_GIT="+passwordStorageDirectory)
	return environment
}
