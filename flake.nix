{
  description = "Go development shell";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs";
    go-license-collector.url = "github:LightJack05/go-license-collector";
  };

  outputs = { self, nixpkgs, go-license-collector, ... }:
    let
      systems = [ "x86_64-linux" "aarch64-linux" "x86_64-darwin" "aarch64-darwin" ];

      forAllSystems = nixpkgs.lib.genAttrs systems;
    in
    {
      devShells = forAllSystems (system:
        let
          pkgs = import nixpkgs {
            inherit system;
          };
        in
        {
          default = pkgs.mkShell {
            name = "go-dev-shell";

            packages = with pkgs; [
              go
              gopls
              gotools
              delve
              go-license-collector.packages.${system}.go-license-collector
            ];

            # Environment variables
            GOPRIVATE = "";
            GONOPROXY = "github.com/SnackLog/*";

            shellHook = ''
              echo "Go development shell ready"
              go version
            '';
          };
        }
      );
    };
}
