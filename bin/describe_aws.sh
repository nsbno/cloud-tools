#!/usr/bin/env bash

the_env=$1

if [[ -z $the_env ]]; then
	echo "Usage: describe_aws ENV"
	exit 1
fi

##
## Helper functions for printing and formatting stuff
##

function print_label() {
	printf "$1:\t\e[3mfetching\e[0m\b\b\b\b\b\b\b\b"
}

function print_heading() {
	printf "\n\e[1m$1\e[0m\n\n"
}

function print_value() {
	read value
	echo -en "          \b\b\b\b\b\b\b\b\b\b"
	echo -e "$value"
}

function join_lines() {
	paste -sd "," -
}

##
## Helper functions for getting stuff from AWS and parsing responses
##

function describe_instance() {
	while read instanceId; do
		envchain aws aws ec2 describe-instances --instance-ids $instanceId
	done
}

function parse_private_ip() {
	jq ".Reservations[].Instances[].NetworkInterfaces[].PrivateIpAddress" | tr -d '"'
}

function parse_public_ip() {
	jq ".Reservations[].Instances[].NetworkInterfaces[].Association.PublicIp" | tr -d '"'
}

function parse_instance_id() {
	jq ".InstanceHealthList[].InstanceId" | tr -d '"'
}

##
## Main script code below
##

print_heading "Base infrastructure"

print_label "Bastion server, private IP"
envchain aws aws ec2 describe-instances --filters "Name=tag-value, Values=${the_env}_bastion" \
 | parse_private_ip \
 | print_value

print_label "Bastion server, public IP"
envchain aws aws ec2 describe-instances --filters "Name=tag-value, Values=${the_env}_bastion" \
 | parse_public_ip \
 | print_value

print_label "Managment server, private IP"
envchain aws aws ec2 describe-instances --filters "Name=tag-value, Values=${the_env}_mgmtserver" \
 | parse_private_ip \
 | print_value


print_heading "Chat"

print_label "EB instances, private IPs"
envchain aws aws elasticbeanstalk describe-instances-health --environment-name "nsb-chat-${the_env}" \
 | parse_instance_id \
 | describe_instance \
 | parse_private_ip \
 | join_lines \
 | print_value
