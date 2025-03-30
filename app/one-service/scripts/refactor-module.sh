#!/usr/bin/env sh
# source from: https://stackoverflow.com/questions/60165440/how-do-i-refactor-module-name-in-go 

export CUR="CHANGE THIS" # example: github.com/user/old-lame-name
export NEW="CHANGE THIS" # example: github.com/user/new-super-cool-name
go mod edit -module ${NEW}
find . -type f -name '*.go' -exec perl -pi -e 's/$ENV{CUR}/$ENV{NEW}/g' {} \;

