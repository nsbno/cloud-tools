#!/usr/bin/env bash

the_env=$1

if [[ -z $the_env ]]; then
	echo "Usage: describe_aws ENV"
	exit 1
fi

function print_label() {
	printf "$1:\t\e[3mfetching\e[0m\b\b\b\b\b\b\b\b"
}

function print_value() {
	read value
	echo -en "          \b\b\b\b\b\b\b\b\b\b"
	echo -e "$value" | tr -d '"'
}

echo "Information about $the_env VPC:"
echo

print_label "Bastion server, private IP"
envchain aws aws ec2 describe-instances --filters "Name=tag-value, Values=${the_env}_bastion" \
 | jq ".Reservations[].Instances[].NetworkInterfaces[].PrivateIpAddress" \
 | print_value

print_label "Bastion server, public IP"
envchain aws aws ec2 describe-instances --filters "Name=tag-value, Values=${the_env}_bastion" \
 | jq ".Reservations[].Instances[].NetworkInterfaces[].Association.PublicIp" \
 | print_value

print_label "Managment server, private IP"
envchain aws aws ec2 describe-instances --filters "Name=tag-value, Values=${the_env}_mgmtserver" \
 | jq ".Reservations[].Instances[].NetworkInterfaces[].PrivateIpAddress" \
 | print_value
