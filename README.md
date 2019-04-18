# grpc-example

## Command to generate code

```console
$ protoc --proto_path ./echo/proto --go_out=plugins=grpc:./echo/proto/ ./echo/proto/echo.proto
```
