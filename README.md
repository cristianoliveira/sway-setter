# sway-setter

The missing CLI to use the outputs from sway builtin "getters": `swaymsg -t get_*`

Sway-setter allows you to use the output of swaymsg to restore state of your workspaces/outputs configuration.

Very useful to restore state after using wdisplays or similar tools, as well your current monitor setup.

**Example of my usecase**
```bash
# adjust your outputs as you please
wdisplays
# Save the current configuration
swaymsg -t get_outputs > ~/.local/state/sway-outputs.json
# Configure to load it on log in
sway-setter -t set_outputs < ~/.local/state/sway-outputs.json
```

## Motivation

I use SwayWM with my notebook and sometimes I have up to 3 monitors in different layouts, and I use wdisplays to configure them. But once I have to restart or log out my system the configuration is lost and I have to reconfigure them. My first attempt was to manually configure them in sway/config, but it was a pain to maintain and I had to change it every time I changed my monitor setup.

I know sway has a way to export the current state of certain objects like `swaymsg -t get_*` and I use those outputs mainly for debugging or scripting. Since the output is always the same why create a custom script? I should be able to load it back! That's the origin of `sway-setter`

## Current and future features

The main goal of sway-setter is to restore current configuration not the current session state (layout, windows, etc). 
The current features are:

- [x] Restore workspaces
  - [x] Restore names
  - [x] Restore workspace per output
  - [x] Restore focused workspace

```bash
swaymsg -t get_workspaces > output.json
sway-setter -t set_workspaces < output.json
```

- [ ] Restore outputs 

  - [x] Restore positions
  - [ ] Restore modes
  - [ ] Restore scale
  - [ ] Restore transform
  - [ ] Restore primary
  - [ ] Restore active

```bash
swaymsg -t get_outputs > output.json
sway-setter -t set_outputs < output.json
```

- [ ] Restore inputt

```bash
swaymsg -t get_inputs > output.json
sway-setter -t set_inputs < output.json
```

- [x] Use as an input for swaymsg, with the flag `--print|-p`

```bash
swaymsg "$(sway-setter -pt set_workspaces < data.json)"
```

If you are looking for session restoring, there are other tools that can do that, such as:

 - [sway-session](https://github.com/gumieri/sway-session)
 - [swayrst](https://github.com/Nama/swayrst)

### More examples

To see more examples check the [snapshots](./e2e/__snapshots__) folder, the snapshots contain the commands expected to be sent to sway.

## Installation

### Nix
  
```bash
nix profile install 'github:cristianoliveira/sway-setter'
```

### Golang

```bash
go install github.com/cristianoliveira/sway-setter
```

### Requirements

 - Golang 1.23 or later

## Running

```bash
go run ./cmd/sway-setter/main.go < fixtures/workspaces.json
```

## Testing

### Unit tests

```bash
make test
```

### E2E tests

```bash
make test-e2e
```

## More commands

For other commands, please check the Makefile

```bash
make help
```

## Contributing

Want to help? Pick a missing feature and open a PR! Any help is welcomed!

## License

MIT
