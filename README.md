# Simple Cheat Engine

Just a simple cheat engine for linux written by golang.

## Installation

Just clone this repository and build go binary with `go build` command.

## Usage

This binary needs superuser privilege for every read/write memory process.

```shell
# ./progname -pid <pid> -maps

# ./progname -pid <pid> -read <address>

# ./progname -pid <pid> -read <address> -size <read_size>

# ./progname -pid <pid> -write <address> -bit <bit_size>

# ./progname -pid <pid> -writestr <address> -str <string>
```
