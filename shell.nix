{
  pkgs ? (
    let
      inherit (builtins) fetchTree fromJSON readFile;
      inherit ((fromJSON (readFile ./flake.lock)).nodes) nixpkgs gomod2nix;
    in
    import (fetchTree nixpkgs.locked) {
      overlays = [
        (import "${fetchTree gomod2nix.locked}/overlay.nix")
      ];
    }
  ),
  mkGoEnv ? pkgs.mkGoEnv,
  gomod2nix ? pkgs.gomod2nix,
}:

let
  goEnv = mkGoEnv { pwd = ./.; };
in
pkgs.mkShell {
  packages = [
    goEnv
    gomod2nix
  ];
  buildInputs = with pkgs; [
    go
    gcc
    # X11 for Ebiten
    xorg.libX11
    xorg.libXcursor
    xorg.libXrandr
    xorg.libXinerama
    xorg.libXi
    xorg.libXxf86vm
    xorg.libXext
    # Graphics
    libGL
    libglvnd
    mesa
    mesa.drivers
  ];

  shellHook = ''
    # Force X11 backend for Ebiten
    export DISPLAY=:0
    export GDK_BACKEND=x11

    # OpenGL library path
    export LD_LIBRARY_PATH=${
      pkgs.lib.makeLibraryPath [
        pkgs.mesa
        pkgs.libGL
        pkgs.libglvnd
        pkgs.xorg.libX11
      ]
    }:$LD_LIBRARY_PATH

    echo "Using X11 backend (via XWayland)"
  '';

  CGO_ENABLED = "1";
}
