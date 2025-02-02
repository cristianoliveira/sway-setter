{ pkgs ? import <nixpkgs> {} }:
{
  default = pkgs.callPackage ./nix/package.nix { inherit pkgs; };
  source = pkgs.callPackage ./nix/package-source.nix { inherit pkgs; };
}
