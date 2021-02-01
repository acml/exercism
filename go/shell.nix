{ pkgs ? import <nixpkgs> {} }:

with pkgs;

mkShell {
  buildInputs = [
    errcheck
    go
    go2nix
    gocode
    godef
    goimports
    golangci-lint
    golint
    gomodifytags
    gopls
    gore
    gosec
    gotags
    gotests
    gotools                     # guru
  ];
  shellHook = ''
    export FLASK_DEBUG=1
  '';
}
