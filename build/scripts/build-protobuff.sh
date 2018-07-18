
echo "[GRPC]: Building protobuff file for GO: ${pwd}/shared/protobuff.proto"
protoc -I shared/ shared/protobuff.proto --go_out=plugins=grpc:shared