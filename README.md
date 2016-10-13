# cloud-tools

## Required software

```
brew install terraform go
```

## Required environment variables

```
mkdir <path-to-sourcecode>/go
```
# Insert into .bashrc 
export GOPATH=<path-to-sourcecode>/go/
export GOBIN=$GOPATH/bin/
PATH=$GOBIN/bin:$PATH
export PATH
```

## Install Terragrunt

Download Terragrunt from https://github.com/gruntwork-io/terragrunt/releases

```
cd ~/Downloads
chmod 0755 ~/Downloads/terragrunt_darwin_amd64
mv ~/Downloads/terragrunt_darwin_amd64 /usr/local/bin/
ln -nsf /usr/local/bin/terragrunt_darwin_amd64 /usr/local/bin/terragrunt
```
Run `terragrunt plan` to check if it works.
