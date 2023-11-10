#!/bin/bash

VERS=$1

git tag "$VERS"
git push --tag


# PATCH=$1
# MINOR=$2
# MAJOR=$3
# FILE=version
#
# if [ ! -f $FILE ]; then
#     echo "not found version"
#     echo "0.0.0" > $FILE
# fi
#
# VERS=$(echo "$(cat $FILE)" | tr "." "\n")
# echo "length: $(!$VERS[@])"
#
# if [[ $PATCH -eq "true" ]]; then
#     echo "patch up: patch=$VERS[2]"
# fi
#
# if [[ $MINOR -eq "true" ]]; then
#     echo "minor up: minor=$VERS[1]"
# fi
#
# if [[ $MAJOR -eq "true" ]]; then
#     echo "minor up: major=$VERS[0]"
# fi