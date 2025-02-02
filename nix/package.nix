{ pkgs, ... }:
  pkgs.buildGoModule rec {
    # name of our derivation
    name = "sway-setter";
    version = "v0.0.1";

    # sources that will be used for our derivation.
    src = pkgs.fetchFromGitHub {
      owner = "cristianoliveira";
      repo = "sway-setter";
      rev = version;
      sha256 = "";
    };

    modSha256 = "";

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
