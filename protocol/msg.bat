protogen -i:LoginInfo.proto -o:LoginInfo.cs
protoc --go_out=. LoginInfo.proto
pause