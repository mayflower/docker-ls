#!/bin/bash

os_list="linux darwin windows openbsd freebsd netbsd"

arch_linux="386 amd64 arm"
arch_darwin="amd64"
arch_windows="386 amd64"
arch_openbsd="386 amd64 arm"
arch_freebsd="386 amd64 arm"
arch_netbsd="386 amd64 arm"

suffix_windows=".exe"

package_prefix="github.com/mayflower/docker-ls"
packages="cli/docker-ls cli/docker-rm"

echo

go test github.com/mayflower/docker-ls/lib/... || exit 1
go generate github.com/mayflower/docker-ls/lib/... || exit 1

test -d release && rm -fr release
mkdir release
mkdir release/archives

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
            CGO_ENABLED=0 GOOS="$os" GOARM=5 GOARCH="$arch" go build -o "$binary" "$full_package" || exit 1
        done

        echo archiving for $os $arch

        zipfile="release/archives/docker-ls-${os}-${arch}.zip"
        shafile="$zipfile.sha256"

        zip --junk-paths "$zipfile" $target_dir/*
        cat "$zipfile" | shasum -a 256 | awk '{print $1;}' > "$shafile"
    done
done
