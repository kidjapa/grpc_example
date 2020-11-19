#!/usr/bin/env bash

#!/usr/bin/env bash

DST_DIR="."
SRC_DIR="./proto"

for filename in $(find ./proto -name '*.proto'); do
  [ -e "$filename" ] || continue

  echo "========================================="
  echo "Generating protofile for: " $filename
  echo "Generating the file: "$(basename "$filename" .proto).pb.go
  echo "========================================="

  protoc --go_out=$DST_DIR --proto_path=proto --go_opt=M$filename=services \
   --go_opt=paths=import --go-grpc_out=. \
   $filename

done