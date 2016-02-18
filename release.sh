#!/bin/bash

os_list="linux darwin windows"

arch_linux="386 amd64"
arch_darwin="386 amd64"
arch_windows="386 amd64"

suffix_windows=".exe"

package_prefix="git.mayflower.de/vaillant-team/docker-ls"
packages="cli/docker-ls cli/docker-rm"

make install || exit 1
export GOPATH="`pwd`/build"

echo

for os in $os_list; do
    arch_list="arch_$os"
    suffix="suffix_$os"
    suffix="${!suffix}"

    for arch in ${!arch_list}; do
        echo building for $os $arch

        target_dir="release/${os}_${arch}"
        mkdir -p "$target_dir"

        for package in $packages; do
            full_package="$package_prefix/$package"
            binary="$target_dir/${full_package##*/}$suffix"
            CGO_ENABLED=0 GOOS="$os" GOARCH="$arch" go build -installsuffix no_cgo -o "$binary" "$full_package" || exit 1
        done
    done
done
