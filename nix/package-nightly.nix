{ pkgs, ... }:
  pkgs.buildGoModule rec {
    # name of our derivation
    name = "sway-setter";
    version = "nightly"; # Branch 

    # sources that will be used for our derivation.
    src = pkgs.fetchFromGitHub {
      owner = "cristianoliveira";
      repo = "sway-setter";
      rev = version;
      sha256 = "sha256-BtZjaeG687UL/GTLxQGTCiD3sckYUMMUNtOoQcEp/NM=";
    };

    vendorHash = "sha256-1pCoyPjokM+Qm+mbNQLCOmVVqwHDUxJ10RyvLV/becQ=";

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
