## placed on the standard path after installing google's PB support
protoc --go_out=./golang teletype.proto

## Download and unzip nanopb sdk into a folder, and point to it with the env var NANOPBSDK
## https://koti.kapsi.fi/jpa/nanopb/
## for example: export NANOPBSDK=~/dev/nano/nanopb
${NANOPBSDK}/generator-bin/protoc -I./ -I${NANOPBSDK}/generator/proto/ --nanopb_out=./clang teletype.proto
