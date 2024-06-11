# Build Scripts for Limitr

This folder contains a bash script ([buildAll.sh](buildAll.sh)) to build Limitr for the following platforms:

- Linux (amd64)
- Linux (arm64)
- Linux (arm)
- Windows (amd64)
- Windows (arm64)
- macOS (amd64)
- macOS (arm64)

## Prerequisites

This script requires Go to run. If you don't have Go installed on your system, you can find installation instructions
at [https://golang.org/doc/install](https://golang.org/doc/install)

The script requires Gum to run. If you don't have Gum installed on your system, you can find installation instructions
at [https://github.com/charmbracelet/gum?tab=readme-ov-file#installxation](https://github.com/charmbracelet/gum?tab=readme-ov-file#installation)

## Usage

To build Limitr for the above platofrms, run the following command (where `[version]` is the version number you're building):

```shell
./buildAll.sh [version]
```

All builds will be placed in the `out` directory. You may need to run `chmod +x` on the executables to make them executable.
