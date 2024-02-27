{
  description = "";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs";
  };

  outputs = { self, nixpkgs }:
    let 
      system = "x86_64-linux";
      lib = nixpkgs.lib;
      overlays = [ (final: prev: {
        javy = pkgs.stdenv.mkDerivation {
          name = "javy";
          version = "1.4.0";
          src = pkgs.fetchurl {
            url = "https://github.com/bytecodealliance/javy/releases/download/v1.4.0/javy-x86_64-linux-v1.4.0.gz";
            hash = "sha256-NZIzT8BdtgKiE3RbePWEY1E5TWe9mr2LSRhhmxzWzd8=";
          };
          nativeBuildInputs = with pkgs; [ gzip autoPatchelfHook ];
          buildInputs = with pkgs; [ glib ];
          unpackPhase = ''
            cp $src src.gz 
            mkdir $out
            ${pkgs.gzip}/bin/gzip -c -d src.gz > $out/javy
          '';
          installPhase = ''
            runHook preInstall
            install -m755 -D $out/javy $out/bin/javy
            runHook postInstall
          '';
          meta = with lib; {
            platforms = platforms.linux;
          };
        };
      }) ];

      pkgs = import nixpkgs { inherit system overlays; };
    in
      {
        packages = {
        };
        # defaultPackage = example;
        devShell.${system} = pkgs.mkShell {
          packages = with pkgs; [ go gopls javy ];
          shellHook = ''
          '';
        };
      };
}
