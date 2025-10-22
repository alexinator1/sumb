#!/bin/bash

# Скрипт для создания .env файла из .env.dist
if [ ! -f .env ]; then
    echo "Creating back .env file from .env.dist..."
    cp .env.dist .env
    echo "Please edit .env file with your local configuration"
else
    echo ".env file already exists"
fi