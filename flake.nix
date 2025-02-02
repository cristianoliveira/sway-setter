{
  description = "sway-setter cli";
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs";
    utils.url = "github:numtide/flake-utils";
  };
  outputs = { self, nixpkgs, utils }: 
    utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; };
      in {
        devShells.default = pkgs.mkShell {
          packages = with pkgs; [
            go

            # File watcher
            funzzy
            yq
          ];
        };

        packages = {
          default = pkgs.callPackage ./nix/package.nix { inherit pkgs; };
          source = pkgs.callPackage ./nix/package-source.nix { inherit pkgs; };
        };
    });
}
