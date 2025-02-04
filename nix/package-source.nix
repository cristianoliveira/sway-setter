{ pkgs, ... }:
  pkgs.buildGoModule rec {
    # name of our derivation
    name = "sway-setter";
    version = "source";

    # sources that will be used for our derivation.
    src = ../.;

    vendorHash = "";

    ldflags = [
      "-s" "-w"
      "-X main.VERSION=${version}"
    ];

    meta = with pkgs.lib; {
      description = "sway-setter a tool for importing sway's configuration";
      homepage = "https://github.com/cristianoliveira/sway-setter";
      license = licenses.mit;
      maintainers = with maintainers; [ cristianoliveira ];
    };
  }
