#!/bin/bash

set -e

MAIN_DIRECTORY="$HOME/.goweather"

rm -rf $MAIN_DIRECTORY

git clone https://github.com/seitin/go-weather-cli/ $MAIN_DIRECTORY

read -p "Add your API key from https://www.weatherapi.com/: " apiKey
read -p "Preferred language: " language

encodedApiKey=$(echo $apiKey | base64)

printf "{\"weatherApiKey\": \"${encodedApiKey}\",\"language\": \"${language}\"}" | jq "." >> $MAIN_DIRECTORY/conf.json

cd $MAIN_DIRECTORY

go build main.go
mkdir ./bin

EXECUTABLE_PATH=$MAIN_DIRECTORY/bin
mv main $EXECUTABLE_PATH/go-weather-cli

printf "To add the executable to you PATH, run the following commands:\n"

cat ./scripts/default
source ./scripts/default
