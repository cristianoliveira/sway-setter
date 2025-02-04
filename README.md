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
sway-setter outputs < ~/.local/state/sway-outputs.json
```

### Integration with Sway

Add the following to your sway config to save and restore your outputs configuration
```
# Load the output configuration on startup
exec_always sway-setter outputs < $HOME/.local/state/sway-outputs.json
# Load the output configuration on demand
set $outputLoadMode "Output config: [s]ave, [r]eload"
mode $outputLoadMode {
    # Mappings
    bindsym s exec swaymsg \
        -t get_outputs > $HOME/.local/state/sway-outputs.json \
        && swaymsg mode "default"
    bindsym r exec sway-setter outputs < $HOME/.local/state/sway-outputs.json \
        && swaymsg mode "default"

    # Exit mode
    bindsym Escape mode "default"
    bindsym Return mode "default"
}
bindsym $mod+Shift+o mode $outputLoadMode
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
sway-setter workspaces < output.json
```

- [ ] Restore outputs 

  - [x] Restore positions
  - [x] Restore modes
  - [ ] Restore scale
  - [x] Restore transform (rotation 90, 180, et)
  - [ ] Restore primary
  - [ ] Restore active

```bash
swaymsg -t get_outputs > output.json
sway-setter outputs < output.json
```

- [ ] Restore containers

```bash
swaymsg -t get_tree > output.json
sway-setter containers < output.json
```

- [x] Move containers to workspaces
- [x] Resize containers 
- [x] floating containers
- [x] floating containers position


```bash
swaymsg "$(sway-setter workspaces -p < data.json)"
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
