# sway-setter

The missing CLI to use the outputs from sway builtin "getters" commands `swaymsg -t get_*`
Sway-setter allows you to dump and restore objects from sway and swaymsg cli.

Very useful to save output confiugurations after using wdisplays or similar tools.

## Installation

### Requirements

 - Golang 1.23 or later


## Running

```bash
go run ./cmd/sway-setter/main.go < fixtures/workspaces.json
```

## Testing

```bash
make test
```

## License

MIT
