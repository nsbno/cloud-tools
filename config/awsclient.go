package config

// Retrieves secrets from the pass password store
func AWSCloudClientConfig() (string, string) {

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
