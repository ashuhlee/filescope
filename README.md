# Filescope
A command line tool to list and inspect files in a directory.

## Usage
```zsh
filescope <path>
```

## Requirements
- [Go](https://go.dev/dl/) 1.21 or higher
- A [nerd font](https://www.nerdfonts.com/) installed and set as your terminal font

## Installation
1. **Clone this repository**
```zsh
git clone https://github.com/ashuhlee/filescope.git
cd starlit
```
2. **Install the binary**
```zsh
go install ./cmd/
```
3. **Ensure `~/go/bin` is in your PATH**
```zsh
export PATH=$PATH:~/go/bin
```