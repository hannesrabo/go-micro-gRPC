#!/bin/bash

# Make sure you grab the latest version
curl -OL https://github.com/google/protobuf/releases/download/v3.2.0/protoc-3.2.0-linux-x86_64.zip

# Unzip
unzip protoc-3.2.0-linux-x86_64.zip -d protoc3

# Move protoc to /usr/local/bin/
mv protoc3/bin/* /usr/local/bin/

# Move protoc3/include to /usr/local/include/
mv protoc3/include/* /usr/local/include/

ln -s /protoc3/bin/protoc /usr/bin/protoc
