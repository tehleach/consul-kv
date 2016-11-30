# consul-kv

Backup, restore, and set kvs from consul.

## Usage

### `consul-kv set <consul address> <key> <value>`

Set key to value at given consul address.

### `consul-kv backup`

Backup kvs from consul.

#### Options

* `--from <value>, -f <value>` (default 'localhost:8500') - consul address to retrieve keys from.
* `--name <value>, -n <value>` (default 'data.json') - filename to save kvs to.
* `--prefix <value>, -p <value>` - prefix to pull keys with. If not supplied, gets all keys.

### `consul-kv restore`

Restore kvs to consul, from either JSON or a different consul.

#### Options

* `--from <value>, -f <value>` (default 'localhost:8500') - consul address to retrieve keys from.
* `--name <value>, -n <value>` (default 'data.json') - filename to retrieve keys from. Takes priority over --from.
* `--to <value>, -t <value>` (default 'localhost:8500') - filename to restore keys to.
* `--prefix <value>, -p <value>` - prefix to restore keys with. If not supplied, restores all keys.

restore keys

## Dependencies

[Golang 1.7](https://golang.org/dl/)

## Install

To install, use `go get`:

```bash
$ go get -d github.com/tehleach/consul-kv
```

## Testing

```bash
$ go test ./...
```

## Author

[Kyle Leach](https://github.com/tehleach)
