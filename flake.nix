{
  inputs.nixpkgs.url = "nixpkgs/nixos-unstable";

  outputs = inputs: let
    pkgs = import inputs.nixpkgs {
      system = "x86_64-linux";
    };
  in {
    devShells.x86_64-linux.default = pkgs.mkShell {
      packages = with pkgs; [        
        go        
        gopls      
      ];      
      buildInputs = with pkgs; [        
        pkg-config        
        systemd
      ];
    };
  };
}
