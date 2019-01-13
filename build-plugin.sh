# build project as plugin instead of executable or binary
rm echo-module.so
go build -buildmode=plugin -o echo-module.so
