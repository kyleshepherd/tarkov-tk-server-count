# tarkov-tk-server-count

Simple Go API for getting server count for Tarkov TK


## Development

 - Go 1.11+
 - Dependencies managed with `go mod`

### Setup

These steps will describe how to setup this project for active development. Adjust paths to your desire.

1. Clone the repository: `git clone git@github.com:kyleshepherd/tarkov-tk-server-count.git tarkov-tk-server-count`
2. Build: `make build`
3. 🍻

### Dependencies

Dependencies are managed using `go mod` (introduced in 1.11), their versions
are tracked in `go.mod`.

To add a dependency:
```
go get url/to/origin
```

### Configuration

Configuration can be provided through a toml file, these are loaded
in order from:

- `/etc/tarkov-tk-server-count/tarkov-tk-server-count.toml`
- `$HOME/.config/tarkov-tk-server-count.toml`

Alternatively a config file path can be provided through the
-c/--config CLI flag.

#### Example tarkov-tk-server-count.toml
```toml
[log]
console = true
level = "debug"  # [debug|info|error]
```
