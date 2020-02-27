# Cloud Tools

This repo contains and describes the tools you are expected to have installed and set up before embarking on the adventure that is working with infrastructure or deploying to the cloud.

The `terraform-wrapper` is a wrapper for Terraform. It's purpose is to:

- read values for secret variables from `pass`
- set up environment variables
- run custom commands (for example to dynamically generate more variables)
- run terraform passing along any arguments

Configure using a file named `cloud-config.yml`. See this file from `mtl` for an example: https://github.com/nsbno/mtl/blob/master/terraform/prod/cloud-config.yml

The rest of this file explains how to install and set everything up.


## Install Terraform

Use the latest point release of Terraform 12 (change below)

```
brew install tfenv
tfenv install 0.12.20
```

## Install Go

```bash
brew install go
```
Create a new folder `/go` alongside the rest of your Vy repos.
Add the following to your `.bashrc`

```bash
export GOPATH=<path-to-where-you-keep-all-your-vy-repos>/go
export GOBIN=$GOPATH/bin
PATH=$GOBIN:$PATH
export PATH
```


## Set up developer environment

```bash
source ~/.bashrc
mkdir -p $GOPATH/{bin,pkg,src/github.com/nsbno,vendor}
go get github.com/nsbno/cloud-tools # Ignore the warning message
cd $GOPATH/src/github.com/nsbno/cloud-tools
./deps.sh
./make.sh
```


## Install additional tools

```
brew install s3cmd jq
sudo easy_install pip
pip install ansible --user
```

> Note: It is better to install `ansible` with `pip` rather than with `brew`. This way `ansible` is available to Python when running certain scripts for getting dynamic inventory when provisioning.

## Test it!

Run the following commands in a terraform base directory to check if it works.

```
envchain aws terraform-wrapper init
envchain aws terraform-wrapper plan
```

---
The tools has borrowed a lot from: https://github.com/digipost/cloud-tools
