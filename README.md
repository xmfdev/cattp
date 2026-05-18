# cattp - cat HTTP

A minimal static file server written in Go.

## Usage

```
cattp -a <address> -p <port> -r <root>
```

### Flags

| Flag | Long form | Description |
|------|-----------|-------------|
| `-a` | `--address=` | Bind address (e.g. `localhost` or `0.0.0.0`) |
| `-p` | `--port=`    | Port to listen on (e.g. `8080`)              |
| `-r` | `--root=`    | Directory to serve files from                |

All three flags are required. The server will exit if any are missing.

### Example

```
cattp --address=localhost --port=8080 --root=static
```

Serves the contents of the`static` directory at `http://localhost:8080`.

## Building

```
go build -o cattp .
```

## Dependencies

* [targs](https://github.com/xmfdev/targs) — argument parsing
