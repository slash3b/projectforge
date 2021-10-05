#!/bin/bash
# Code generated by Project Forge, see https://projectforge.dev for details.

set -eo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/../..

TGT=$1
[ "$TGT" ] || TGT="0.0.0"

echo "building gomobile for Android..."
mkdir -p build/dist/mobile_android_arm64
time gomobile bind -o build/dist/mobile_android_arm64/projectforge.aar -target=android github.com/kyleu/projectforge/app/cmd
echo "gomobile for Android completed successfully, building distribution..."
cd "build/dist/mobile_android_arm64"
zip -r "../projectforge_${TGT}_mobile_android_aar.zip" .

echo "creating Android project..."
cd $dir/../..
mkdir -p build/dist/mobile_android_app_arm64
cp -R tools/android/* build/dist/mobile_android_app_arm64

echo "building Android project..."
cd build/dist/mobile_android_app_arm64
mkdir -p ./app/libs
rm -rf ./app/libs/projectforge.aar ./app/libs/projectforge-sources.jar
cp ../mobile_android_arm64/projectforge.aar ./app/libs/
cp ../mobile_android_arm64/projectforge-sources.jar ./app/libs/

gradle assembleDebug
cd app/build/outputs/apk/debug
zip -r "$dir/../../build/dist/projectforge_${TGT}_mobile_android_apk.zip" .
