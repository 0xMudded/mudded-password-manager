# Mudded Password Manager

A simple terminal-based password manager.

# Requirements

- [Go](https://golang.org/) >= 1.19
- Core dependencies: `gnugpg`, `gpgme>=1.7.0`, `libgpg-error`

# Installation

Install [gnupg](https://www.gnupg.org/)

# Usage

```sh
mpw [COMMAND] [ARG]
```

```
COMMANDS:
    init        Initializes new store (.store) inside user's home directory. User must specify their GPG key ID as the required arg.
    generate    Generates new password. Password is encrypted, stored in password store and copied to clipboard. User must specify the password name as the required arg.
    get         Retrieves password from the store. Password is then decrypted and copied to clipboard. User must specify the password name as the required arg.
    help        Help about any command
```
