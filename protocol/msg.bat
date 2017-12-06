protogen -i:LandlordsProto.proto -o:LandlordsProto.cs
protoc --go_out=. LandlordsProto.proto
protoc --plugin=protoc-gen-lua="..\pluginLua\protoc-gen-lua.bat" --lua_out=. LandlordsProto.proto
pause