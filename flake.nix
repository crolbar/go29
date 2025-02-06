{
  inputs.nixpkgs.url = "nixpkgs/nixos-unstable";

  outputs = inputs: let
    system = "x86_64-linux";
    pkgs = import inputs.nixpkgs {inherit system;};
  in {
    devShells.x86_64-linux.default = pkgs.mkShell {
      packages = with pkgs; [
        go
        gopls
      ];
      buildInputs = with pkgs; [systemd];
    };
    packages.x86_64-linux.default = pkgs.buildGoModule {
      pname = "go29";
      version = "0.1";
      src = ./.;
      buildInputs = with pkgs; [systemd.dev];
      vendorHash = "sha256-LD5FKaW1e1HFEWc6i0iLw0LG8sFdy5WNsisBoMgN5yA=";
    };
  };
}
