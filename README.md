# cloud-tools

## Required software

```
brew install terraform go
```

# Insert into .bashrc
export GOPATH=<path-to-sourcecode>/go
export GOBIN=$GOPATH/bin
PATH=$GOBIN/bin:$PATH
export PATH
```

## Set up developer environment

```
source ~/.bashrc
mkdir -p $GOPATH/{bin,pkg,src/github.com/nsbno,vendor}
go get github.com/nsbno/cloud-tools # Ignore the warning message
cd $GOPATH/src/github.com/nsbno/cloud-tools
./deps.sh
./make.sh
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
