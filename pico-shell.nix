{pkgs}:
with pkgs;
  mkShell {
    buildInputs = [
      delve
      go
      golangci-lint
      gopls
      jq
      tinygo
    ];

    shellHook = ''
      tinygotarget=$(tinygo info -json -target=pico)
      export GOROOT=$(jq -r '.goroot' <<< "$tinygotarget")
      export GOOS=$(jq -r '.goos' <<< "$tinygotarget")
      export GOARCH=$(jq -r '.goarch' <<< "$tinygotarget")
      export GOFLAGS=-tags=$(jq -r '.build_tags | join(",")' <<< "$tinygotarget")
      export GOPATH=$(git rev-parse --show-toplevel)/go:"${tinygo}/share/tinygo"
    '';
  }
