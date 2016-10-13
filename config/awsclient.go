package config

// Retrieves access key id and secret key from the pass
// password store
func AWSCloudClientConfig() (string, string, string) {

	var access string
	var secret string

	cloudConfig := ParseDefaultCloudConfig()

	for _, secVar := range cloudConfig.SecretVariables {

		if secVar.Name == "AWS_ACCESS_KEY_ID" {
			access = GetPasswordFor(secVar.Key)
		}

		if secVar.Name == "AWS_SECRET_ACCESS_KEY" {
			secret = GetPasswordFor(secVar.Key)
		}

	}

	return access, secret

}
