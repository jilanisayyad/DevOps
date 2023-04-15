#!/bin/bash

# Function to display error message and exit
function display_error() {
  echo "Error: $1"
  exit 1
}

# Function to validate user input for yes/no options
function validate_yes_no_input() {
  local input
  while true; do
    read -p "$1 (y/n): " input
    case $input in
    [Yy]*)
      return 0
      ;;
    [Nn]*)
      return 1
      ;;
    *)
      echo "Invalid input. Please enter y or n."
      ;;
    esac
  done
}

# Function to create the pod
function create_pod() {
  local pod_name=$1
  local image=$2
  local namespace=$3
  local concurrent=$4
  local labels=$5
  local annotations=$6
  local volumes=$7
  local secrets=$8
  local configmaps=$9
  local nodename=${10}
  local nodeaffinities=${11}
  local port=${12}
  local command=${13}
  local envvars=${14}
  local livenessprobe=${15}
  local readinessprobe=${16}
  local startupprobe=${17}

  # Create basic pod configuration
  local pod_config="apiVersion: v1
kind: Pod
metadata:
  name: $pod_name
  namespace: $namespace
"

  # Add labels and annotations to metadata
  if [ -n "$labels" ]; then
    pod_config+="  labels:
$labels"
  fi

  if [ -n "$annotations" ]; then
    pod_config+="  annotations:
$annotations"
  fi

  # Add node name and node affinities
  if [ -n "$nodename" ]; then
    pod_config+="  nodeName: $nodename"
  fi

  if [ -n "$nodeaffinities" ]; then
    pod_config+="  affinity:
    $nodeaffinities"
  fi

  # Add volumes, secrets, and configmaps
  if [ -n "$volumes" ]; then
    pod_config+="  volumes:
$volumes"
  fi

  if [ -n "$secrets" ]; then
    pod_config+="  secrets:
$secrets"
  fi

  if [ -n "$configmaps" ]; then
    pod_config+="  configMaps:
$configmaps"
  fi

  # Add container and image
  pod_config+="spec:
  containers:
  - name: $pod_name
    image: $image"

  # Add command and args
  if [ -n "$command" ]; then
    pod_config+="    command: $command"
  fi

  # Add environment variables
  if [ -n "$envvars" ]; then
    pod_config+="    env:
$envvars"
  fi

  # Add ports
  if [ -n "$port" ]; then
    pod_config+="    ports:
$port"
  fi

  # Add liveness probe
  if [ -n "$livenessprobe" ]; then
    pod_config+="    livenessProbe:
$livenessprobe"
  fi

  # Add readiness probe
  if [ -n "$readinessprobe" ]; then
    pod_config+="    readinessProbe:
$readinessprobe"
  fi

  # Add startup probe
    if [ -n "$startupprobe" ]; then
    pod_config+="    startupProbe:
$startupprobe"
  fi

  # Add concurrent option
  if [ "$concurrent" == "true" ]; then
    pod_config+="
  restartPolicy: Always
  "
  else
    pod_config+="
  restartPolicy: OnFailure
  "
  fi

  # Create the pod
  echo "$pod_config" | kubectl create -f -

  echo "Pod $pod_name created successfully!"
}

# Main script

# Get pod name
read -p "Enter pod name: " pod_name

# Get image
read -p "Enter image: " image

# Get namespace
read -p "Enter namespace (default: default): " namespace
namespace=${namespace:-default}

# Get concurrent option
validate_yes_no_input "Should the pod restart always (concurrent)?" && concurrent="true" || concurrent="false"

# Get labels
read -p "Enter labels in YAML format (optional): " labels

# Get annotations
read -p "Enter annotations in YAML format (optional): " annotations

# Get volumes
read -p "Enter volumes in YAML format (optional): " volumes

# Get secrets
read -p "Enter secrets in YAML format (optional): " secrets

# Get configmaps
read -p "Enter configmaps in YAML format (optional): " configmaps

# Get node name
read -p "Enter node name (optional): " nodename

# Get node affinities
read -p "Enter node affinities in YAML format (optional): " nodeaffinities

# Get ports
read -p "Enter ports in YAML format (optional): " port

# Get command and args
read -p "Enter command and args in YAML format (optional): " command

# Get environment variables
read -p "Enter environment variables in YAML format (optional): " envvars

# Get liveness probe
read -p "Enter liveness probe in YAML format (optional): " livenessprobe

# Get readiness probe
read -p "Enter readiness probe in YAML format (optional): " readinessprobe

# Get startup probe
read -p "Enter startup probe in YAML format (optional): " startupprobe

# Call the create_pod function with user inputs
create_pod "$pod_name" "$image" "$namespace" "$concurrent" "$labels" "$annotations" "$volumes" "$secrets" "$configmaps" "$nodename" "$nodeaffinities" "$port" "$command" "$envvars" "$livenessprobe" "$readinessprobe" "$startupprobe"

