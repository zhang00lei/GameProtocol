protogen -i:LandlordsProto.proto -o:LandlordsProto.cs
protoc --go_out=. LandlordsProto.proto
protoc -I=. --descriptor_set_out=.\LandlordsProto.pb LandlordsProto.proto

pause