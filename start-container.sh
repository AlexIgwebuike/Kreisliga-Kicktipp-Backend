#!/bin/bash

# Set Container and Image name
IMAGE_NAME="kicktipp-backend"
CONTAINER_NAME="kicktipp-backend"

read -p "Did you make any changes in your project? [y/n] " response

# Check User input
if [[ "$response" == "y" || "$response" == "Y" ]]; then
    # Build Docker Image
    echo "Building Docker image..."
    docker build -t $IMAGE_NAME .

    #Check if exists and is still running to either stop Container if needed, or not.
    if [ "$(docker ps -a -q -f name=$CONTAINER_NAME)" ]; then

        # Stop and remove the old container
        echo "Stopping and removing old container..."
        docker stop $CONTAINER_NAME
        docker rm $CONTAINER_NAME
    fi
elif [[ "$response" == "n" || "$response" == "N" ]]; then
    echo "Skipping Docker image build..."
else
    echo "Invalid response. Please enter 'y' or 'n'."
    exit 1
fi

# Check if container exists
if [ $(docker ps -a -q -f name=$CONTAINER_NAME) ]; then
    
    #Check if container is running
    if [ $(docker ps -q -f name=$CONTAINER_NAME) ]; then
        echo "Container $CONTAINER_NAME is already running"
        exit 0
    fi

    echo "Container $CONTAINER_NAME already exists, starting container..."
    docker start $CONTAINER_NAME
    exit 0
fi

# Create a new Docker container mapping on Port 8080
echo "Starting new Container on Port: 8080"
docker run -d -p 8080:8080 --name $CONTAINER_NAME $IMAGE_NAME
