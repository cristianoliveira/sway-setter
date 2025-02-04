#!/usr/bin/env bash

function bump_version() {
  echo "Bumping version to $VERSION in $NIX_FILE"

  sed -i 's/sha256-.*=//g' "$NIX_FILE"
  ## nix build and pipe the error to a build.log file
  rm -f build.log

  nix build ."#$FLAKE_PKG_NAME" 2> build.log

  SHA256=$(grep "got:" build.log | grep -o "sha256-.*=" | cut -d'-' -f2)

  echo "git hash SHA256: $SHA256"
  sed -i "s# hash = \".*\";# hash = \"sha256-$SHA256\";#" "$NIX_FILE"
  nix build ."#$FLAKE_PKG_NAME" 2> build.log

  SHA256=$(grep "got:" build.log | grep -o "sha256-.*=" | cut -d'-' -f2)
  echo "cargo hash SHA256: $SHA256"
  sed -i "s#cargoHash = \".*\";#cargoHash = \"sha256-$SHA256\";#" "$NIX_FILE"

  echo "Building nix derivation"
  nix build ."#$FLAKE_PKG_NAME"

  rm -f build.log

  git add "$NIX_FILE"
  git commit -m "chore(nix): bump $FLAKE_PKG_NAME"
}

for nix_file in $(find nix -name 'package-*.nix'); do
  NIX_FILE="$nix_file"
  FLAKE_PKG_NAME=$(basename "$nix_file" | cut -d'-' -f2 | cut -d'.' -f1)
  bump_version
done
