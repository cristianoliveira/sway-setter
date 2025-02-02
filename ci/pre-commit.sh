#!/usr/bin/env bash

#
# An example hook script to verify what is about to be committed.
# Called by "git commit" with no arguments.  The hook should
# exit with non-zero status after issuing an appropriate message if
# it wants to stop the commit.
#
# To enable this hook, rename this file to "pre-commit".

# Check if pre-commit is "enabled" if not then create a symlink to this file
if [ ! -f .git/hooks/pre-commit ]; then
  answer="n"
  read -p "Do you want to install pre-commit hook? [y/n]" answer

  if [ "$answer" != "${answer#[Yy]}" ] ;then
    ln -s $PWD/ci/pre-commit.sh $PWD/.git/hooks/pre-commit
  else
    echo "Skipping pre-commit hook installation"
    exit 0
  fi
fi

make fmt
make test
make test-e2e
