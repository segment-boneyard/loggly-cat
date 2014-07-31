
# loggly-cat

  Stream logs to loggly via stdio.

## Usage

```

  Usage:
    loggly-cat --token t [--tag t]...
    loggly-cat -h | --help
    loggly-cat --version

  Options:
    -t, --token t     loggly api token
    -T, --tag t       loggly tag(s)
    -h, --help        output help information
    -v, --version     output version

```

## Installation

```
$ go get github.com/segmentio/loggly-cat
```

## Example

```
$ myapp | loggly-cat -t my-token-here -T myapp
```

# License

 MIT