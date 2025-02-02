{ pkgs, ... }:
  pkgs.buildGoModule rec {
    # name of our derivation
    name = "sway-setter";
    version = "v0.0.2";

    # sources that will be used for our derivation.
    src = pkgs.fetchFromGitHub {
      owner = "cristianoliveira";
      repo = "sway-setter";
      rev = version;
      sha256 = "sha256-TQAcKtIid1+zP6dRixcuH6jZZHko1+ujZH2Kya+Snq8=";
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
