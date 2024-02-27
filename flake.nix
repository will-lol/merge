{
  description = "";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs";
  };

  outputs = { self, nixpkgs }:
      let 
        system = "x86_64-linux";
        overlays = [];
        lib = nixpkgs.lib;

        javyCli = pkgs.stdenv.mkDerivation {
          name = "javy";
          version = "1.4.0";
          src = pkgs.fetchurl {
            url = "https://github.com/bytecodealliance/javy/releases/download/v1.4.0/javy-x86_64-linux-v1.4.0.gz";
            hash = "sha256-SfBLy1Q7OlHf/7Qsv9cVmpHNPGgNVxtBvojBmSzx0vI=";
            postFetch = ''
              cp $out src.gz
              ${pkgs.gzip}/bin/gzip -c -d src.gz > $out
            '';
          };

          nativeBuildInputs = with pkgs; [ gzip autoPatchelfHook ];
          buildInputs = with pkgs; [ openssl glibc gcc-unwrapped ];
          phases = [ "installPhase" ];
          installPhase = ''
            runHook preInstall
            install -m755 -D $src $out/bin/javy
            runHook postInstall
          '';
          meta = with lib; {
            platforms = platforms.linux;
          };
        };
        pkgs = import nixpkgs { inherit system overlays; };
      in
        {
          packages = {
          };
          # defaultPackage = example;
          devShell.${system} = pkgs.mkShell {
            packages = [ pkgs.go pkgs.gopls javyCli ];
            shellHook = ''
            '';
          };
        }
    ;
}
