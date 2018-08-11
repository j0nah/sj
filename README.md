#sj

sj is a tiny utility that helps you search for all occurrences of a key in json

## Installation

  Binary releases are available, or run:

```
$ go get github.com/j0nah/sj/cmd/sj
```

## Usage

#### func  Search

```go
res, err := sj.Search(<needle>, <haystack>)
```
`res` is a `*[]string*` containing a list of the values of the first occurrences of the needle in the json subgraph.

## License

MIT