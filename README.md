# MySQL Toolkit

MySQL Toolkit is a light-weight CLI to help you run queries via the command line.

## Usage

`mysql-toolkit --help` will show usage and a list of commands:

```
MySQL Toolkit is a light-weight CLI to help you run queries via the command line.

Usage:
  mysql-toolkit [flags]
  mysql-toolkit [command]

Available Commands:
  debug       Debug internal information
  help        Help about any command
  query       Execute SQL Query
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

### Example

```bash
# Get list of all databases
mysql-toolkit query "SHOW DATABASES" | jq -r ".[] | .Database"
```

## Configuration

There are many different ways to configure `mysql-toolkit` CLI:
- Use Environment Variables (use `MYSQL_TOOLKIT` prefix)
- Store Configuration in File (locate `mysql-toolkit.yaml` in `$PWD` or `$HOME`)
- Pass Configurations via Flags (checkout usage `mysql-toolkit --help`)

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

# Clean bin folder
make clean
```
