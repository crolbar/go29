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
    packages.x86_64-linux.default = pkgs.buildGoModule rec {
      pname = "go29";
      version = "0.1";
      src = ./.;
      buildInputs = with pkgs; [systemd.dev];
      vendorHash = "sha256-r43FjbKH3g/ZJp43KvvcbPyphZQP/HDmZ11XMS227Bs=";
      installPhase = ''
        runHook preInstall

        mkdir -p $out
        dir="$GOPATH/bin"
        [ -e "$dir" ] && cp -r $dir $out

        mkdir -p $out/share/applications
        cp $src/${pname}.desktop $out/share/applications

        runHook postInstall
      '';
    };
  };
}
