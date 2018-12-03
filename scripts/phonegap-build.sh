#!/bin/bash
# Usage phonegap-build.sh output_file_name

if [ -z $PHONEGAP_BUILD_APP_ID ]; then
    echo "PHONEGAP_APP_ID environment variable not found"
    exit 1
fi

if [ -z $PHONEGAP_BUILD_SECRET ]; then
    echo "PHONEGAP_SECRET environment variable not found"
    exit 1
fi

if [ -z $1 ]; then
    echo "Missing filename environment variable not found"
    exit 1
fi

OUTPUT_NAME=$1

echo "Compressing..."
zip -r --quiet $OUTPUT_NAME.zip *


echo "Uploading to PhoneGap Build..."
UPLOAD=$(curl -X PUT --progress-bar -F file=@$OUTPUT_NAME.zip https://build.phonegap.com/api/v1/apps/$PHONEGAP_BUILD_APP_ID?auth_token=$PHONEGAP_BUILD_SECRET)

rm $OUTPUT_NAME.zip

printf "Building"
until $(curl --output /dev/null --silent --head --fail https://build.phonegap.com/api/v1/apps/$PHONEGAP_BUILD_APP_ID/android/?auth_token=${PHONEGAP_BUILD_SECRET}); do
    printf '.'
    sleep 1
done

echo
echo "Downloading file"

FILE_URL=$(curl -s https://build.phonegap.com/api/v1/apps/$PHONEGAP_BUILD_APP_ID/android/?auth_token=${PHONEGAP_BUILD_SECRET})
FILE_URL=$(node -e "console.log(${FILE_URL}.location)")

curl --progress-bar -X GET $FILE_URL > $OUTPUT_NAME.apk