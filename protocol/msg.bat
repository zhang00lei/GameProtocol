protogen -i:LoginInfo.proto -o:LoginInfo.cs
protoc --go_out=. LoginInfo.proto
protoc --plugin=protoc-gen-lua="..\pluginLua\protoc-gen-lua.bat" --lua_out=. LoginInfo.proto
pause