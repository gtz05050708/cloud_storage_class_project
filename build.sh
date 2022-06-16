#!/bin/bash
# shellcheck disable=SC2068

# !!!!!! READ BEFORE YOU RUN THE SCRIPT !!!!!!!!!
#    You need to setup your runtime environment variable,
#    such that the script can look up the correct bin directory
#    and run the binaries we just compiled
# 
# If you are using a Mac, open ~/.bash_profile
# If you are using unix/linux machine, open ~/.bashrc
# 
# 1. Append these lines to the bash config file ^^^
#  export GOPATH=<YOUR WORKING DIRECTORY>/CSE-x24-Project-2/Project-2
#  export PATH=$PATH:$GOPATH/bin
#  export GO111MODULE=off
#
# 2. Run `source ~/.bash_profile` or `source ~/.bashrc` to make it effective
# 
# 3. Voila

# Clean up the existing binaries in current directory
rm -rf ./bin/surfstoreAdminExec ./bin/surfstoreServerExec ./bin/surfstoreDebugExec

# Build and install the necessary binaries for scripts to run
cd src/surfstore/
go install ./...
cd ../..

# Build the docker image for run-client.sh to run
docker build -t surfstore-client .
