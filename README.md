# Goutput

Simple CLI tool for formated STDOUT/STDERR handling.

## Getting started

First you must clone this repository and the following command:

```shell
go build -o gout
```

Then in the same directory run your builded binary:

```shell
./gout -d -m "Some message: '{}' and '{}'" -p "1" -p "2"
```

Finally you will get:

`[DEBUG]: Some message: '1' and '2'`

## Developing

To develop on this project you will need golang version `go1.15.2` or superior.

To test your code:

```shell
go test github.com/pascencio/goutput/...
```

## TODO

- [ ] Add continous integration
- [ ] Add release generation
- [ ] Add date support 