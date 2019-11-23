# MySQL Toolkit

MySQL Client Toolkit helps you to handle basic MySQL commands in bash.

## Usage

`mysql-toolkit` or `mysql-toolkit --help` will show usage and a list of commands:

```
MySQL Client Toolkit helps you to handle basic MySQL commands in bash.

Usage:
  mysql-toolkit [flags]
  mysql-toolkit [command]

Available Commands:
  help        Help about any command
  version     Toolkit version detail

Flags:
  -c, --config string   config file (default is $HOME/.mysql-toolkit.yaml)
  -d, --db string       database name to select
  -h, --help            help for mysql-toolkit
  -p, --pass string     password for MySQL server (default "password")
  -s, --server string   MySQL server host name (default "127.0.0.1")
  -u, --user string     username for MySQL server (default "root")

Use "mysql-toolkit [command] --help" for more information about a command.
```

## Configuration

There are many different ways to configure `mysql-toolkit` CLI:
- Use Environemnt Variables (use `MYSQL_TOOLKIT` prefix)
- Store Configuration in File (locate `mysql-toolkit.yaml` in `$PWD` or `$HOME`)
- Pass Configutations via Flags (checkout usage `mysql-toolkit --help`)

## Development

```bash
# Run the application with CLI help
go run main.go --help

# Check code quality
make check

# Build for local
make build

# Release for all platforms
make release
```
