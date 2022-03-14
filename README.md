# Casing

This is a simple CLI tool that changes the _casing_ of the arguments.

The program is a wrapper for `github.com/iancoleman/strcase`.

## Usage

```sh
casing [-m MODE] ...<ARGUMENTS>

```

Available modes:

| Transformation |    Modes     |
| :------------- | :----------: |
| kebab-case     | `k`, `kebab` |
| camel-case     | `c`, `camel` |
| snake-case     | `s`, `snake` |
| lower-camel-case     | `lc`, `lowercamel`, `lowerCamel` |

Transforming the input to kebab-case:

```sh
# Note that the input has to be wrapped in quotes for the program to recognize the argument as a
# single string
$ casing -m kebab "Hello World"
hello-world

$ casing -m k "Hello World"
hello-world

# since kebab is the default mode, you can omit it
$ casing "Hello World"
hello-world
```

Transforming the input to camel-case:

```sh
$ casing -m camel "Hello World"
HelloWorld

$ casing -m c "Hello World"
HelloWorld
```

Transforming multiple arguments:
```sh
$ casing "multiple argument test" "HelloWorld" "snekCase" "This is a simple CLI tool that changes the _casing_ of the arguments"
multiple-argument-testhello-world
snek-case
this-is-a-simple-cli-tool-that-changes-the--casing--of-the-arguments
```

## Installation

- Clone the repo
- Run make (the output will be an executable in a new `build` folder in the repository)
- \[Optionally\] Use the `link.sh` script to link the executable from the previous step to your `$HOME/.local/bin` folder