#!/bin/bash

# Setze den Namen des Images und Containers
IMAGE_NAME="kicktipp-backend"
CONTAINER_NAME="kicktipp-backend"

# Docker Image bauen
docker build -t $IMAGE_NAME .

# Prüfe, ob ein Container mit dem Namen bereits läuft, und stoppe/lösche ihn, falls vorhanden
if [ $(docker ps -q -f name=$CONTAINER_NAME) ]; then
    echo "Stopping and removing existing container $CONTAINER_NAME..."
    docker stop $CONTAINER_NAME
    docker rm $CONTAINER_NAME
fi

# Neuen Container starten
docker run -d -p 8080:8080 --name $CONTAINER_NAME $IMAGE_NAME
