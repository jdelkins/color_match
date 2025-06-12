{
  description = "color_match - various tools to help with color ricing";

  inputs = {
    nixpkgs.url = "nixpkgs/nixos-unstable";
    gomod2nix = {
      url = "github:nix-community/gomod2nix";
      inputs.nixpkgs.follows = "nixpkgs";
    };
  };

  outputs = { self, nixpkgs, gomod2nix }: let
    system = "x86_64-linux";
    pkgs = import nixpkgs { inherit system; overlays = [ gomod2nix.overlays.default ]; };
  in {
    packages.${system} = rec {
      color_match = pkgs.buildGoApplication {
        pname = "color_match";
        version = "0.9";
        src = ./.;
        modules = ./gomod2nix.toml;
      };
      default = color_match;
    };

    devShells.${system}.default = pkgs.mkShell {
      packages = with pkgs; [
        go
        gomod2nix.packages.${system}.default
        gopls
      ];
    };
  };
}
