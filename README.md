# Cloud Tools

This repo contains and describes the tools you are expected to have installed and set up before embarking on the adventure that is working with infrastructure or deploying to the cloud.

The `terragrunt-wrapper` is a wrapper for Terragrunt, which in turn wraps Terraform. It's purpose is to:

- read values for secret variables from `pass`
- set up environment variables
- run custom commands (for example to dynamically generate more variables)
- run terraform passing along any arguments

Configure using a file named `cloud-config.yml`. See this file from `mtl` for an example: https://github.com/nsbno/mtl/blob/master/terraform/prod/cloud-config.yml

The rest of this file explains how to install and set everything up.


## Install `pass`

See https://github.com/nsbno/.password-store for information on setting this up.


## Install Terraform and Terragrunt

```
brew install terraform
```

Download Terragrunt from https://github.com/gruntwork-io/terragrunt/releases

_(select the released file `terragrunt_darwin_amd64` from the newest release that has this package)_

```bash
cd ~/Downloads
chmod 0755 ~/Downloads/terragrunt_darwin_amd64
mv ~/Downloads/terragrunt_darwin_amd64 /usr/local/bin/
ln -nsf /usr/local/bin/terragrunt_darwin_amd64 /usr/local/bin/terragrunt

```


## Install Go

```bash
brew install go
```

Add the following to your `.bashrc`

```bash
export GOPATH=<path-to-your-sourcecode>/go
export GOBIN=$GOPATH/bin
PATH=$GOBIN:$PATH
export PATH
```

> Note: `path-to-your-sourcecode/go` should point to an empty folder you create to store your go code in.


## Set up developer environment

```bash
source ~/.bashrc
mkdir -p $GOPATH/{bin,pkg,src/github.com/nsbno,vendor}
go get github.com/nsbno/cloud-tools # Ignore the warning message
cd $GOPATH/src/github.com/nsbno/cloud-tools
./deps.sh
./make.sh
```


## Install envchain to set aws credentials

```bash
brew install envchain
envchain --set aws AWS_ACCESS_KEY_ID AWS_SECRET_ACCESS_KEY AWS_DEFAULT_REGION
```

> Note: AWS_DEFAULT_REGION = eu-central-1

See: https://github.com/sorah/envchain


## Install additional tools

```
brew install s3cmd jq
sudo easy_install pip
pip install awscli awsebcli ansible
```

> Note: It is better to install `ansible` with `pip` rather than with `brew`. This way `ansible` is available to Python when running certain scripts for getting dynamic inventory when provisioning.

## Test it!

Run `envchain aws terragrunt-wrapper plan` in a terraform base directory to check if it works.


---
The tools has borrowed a lot from: https://github.com/digipost/cloud-tools
