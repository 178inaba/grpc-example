# grpc-example

## Command to generate code

### Echo

```console
$ protoc --proto_path ./echo/proto --go_out=plugins=grpc:./echo/proto/ ./echo/proto/echo.proto
```

### Downloader

```console
$ protoc -I ./downloader/proto --go_out=plugins=grpc:./downloader/proto/ ./downloader/proto/file.proto
```

## License

[MIT](LICENSE)

## Author

[178inaba](https://github.com/178inaba)
