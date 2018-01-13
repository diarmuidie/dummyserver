#!/usr/bin/env bash

package=$1
version=$2
if [[ -z "$package" ]]; then
  echo "usage: $0 <package-name> [<version>]"
  exit 1
fi

if [ -z "$version" ]; then
    version='development-build'
fi

platforms=(
  "darwin/386"
  "darwin/amd64"
  "linux/386"
  "linux/amd64"
  "linux/arm64"
  "netbsd/386"
  "netbsd/amd64"
  "windows/amd64"
  "windows/386"
)

mkdir -p bin

for platform in "${platforms[@]}"
do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}

    output_name='bin/'$package'-'$version'-'$GOOS'-'$GOARCH
    if [ $GOOS = "windows" ]; then
        output_name+='.exe'
    fi

    echo $output_name

    env GOOS=$GOOS GOARCH=$GOARCH go build -ldflags "-X main.version=$version" -o $output_name .
    if [ $? -ne 0 ]; then
        echo 'An error has occurred! Aborting the script execution...'
        exit 1
    fi
done
