protogen -i:login.proto -o:Login.cs
protoc --go_out=. login.proto
protoc --plugin=protoc-gen-lua="..\pluginLua\protoc-gen-lua.bat" --lua_out=. login.proto
pause