{
  description = "Sumneko-lua-language-server";

  inputs.flake-utils.url = "github:numtide/flake-utils";

  outputs = { self, nixpkgs, flake-utils }:

    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; };
      in
      {
        devShell = pkgs.mkShell {
          name = "lua-lsp";

          buildInputs = with pkgs; [
            (lua5_4_compat.withPackages(ps: with ps; [ busted ]))
            sumneko-lua-language-server
          ];
        };
      }
    );
}
