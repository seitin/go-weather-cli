#!/bin/bash

set -e

MAIN_DIRECTORY="$HOME/.goweather"

rm -rf $MAIN_DIRECTORY

git clone https://github.com/seitin/go-weather-cli/ $MAIN_DIRECTORY

while [ "$apiKey" = "" ]
do
  read -p "Add your API key from https://www.weatherapi.com/: " apiKey
done

read -p "Preferred language (English is default if not set): " language
# while [ "$automate" -ne "y" ] || [ "$automate" -ne "n" ] || [ "$automate" = "" ]
# do
#   read -p "Do you want to get the location from your IP? (y/n) " automate
# done

# if [ $automate = "n" ];
# then
#   read -p "Which city do you want to track: " city
# fi


encodedApiKey=$(echo $apiKey | base64)

printf "{
\"weatherApiKey\": \"${encodedApiKey}\",
\"language\": \"${language}\",
\"city\": \"${city}\",
\"geolocation\": {
    \"lon\": \"${lon}\",
    \"lat\": \"${lat}\"
  }
}" | jq "." >> $MAIN_DIRECTORY/conf.json

cd $MAIN_DIRECTORY

go build main.go
mkdir ./bin

EXECUTABLE_PATH=$MAIN_DIRECTORY/bin
mv main $EXECUTABLE_PATH/go-weather-cli

printf "To add the executable to you PATH, run the following commands:\n"

cat ./scripts/default
source ./scripts/default
