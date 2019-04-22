# grpc-example

## Command to generate code

### Echo

```console
$ protoc --proto_path ./echo/proto --go_out=plugins=grpc:./echo/proto/ ./echo/proto/echo.proto
```

### Downloader

`-I` is a short option of `--proto_path`.

```console
$ protoc -I ./downloader/proto --go_out=plugins=grpc:./downloader/proto/ ./downloader/proto/file.proto
```

### Uploader

```console
$ protoc -I ./uploader/proto --go_out=plugins=grpc:./uploader/proto/ ./uploader/proto/file.proto
```

### Chat

```console
$ protoc -I ./chat/proto --go_out=plugins=grpc:./chat/proto/ ./chat/proto/chat.proto
```

## License

[MIT](LICENSE)

## Author

[178inaba](https://github.com/178inaba)
