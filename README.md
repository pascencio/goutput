# Goutput

[![Build Status](https://travis-ci.org/pascencio/goutput.svg?branch=master)](https://travis-ci.org/pascencio/goutput)

Simple CLI tool for formated STDOUT/STDERR handling.

- [Goutput](#goutput)
  - [Installation](#installation)
  - [Getting started](#getting-started)
    - [From sources](#from-sources)
  - [From binaries](#from-binaries)
  - [Developing](#developing)
  - [TODO](#todo)
    - [Release 0.0.1](#release-001)
    - [Release 0.0.2](#release-002)
    - [Release 0.4.0](#release-040)

## Installation

Execute the following command to intall the latest version.

```shell
curl -L -o gout.zip https://github.com/pascencio/goutput/releases/download/0.3.2/gout-0.3.2.zip && \
unzip gout.zip && \
sudo mv gout /usr/bin && \
sudo chmod +x /usr/bin/gout && \
gout --version
```

Then you will see a output message like this:

```shell
Goutput version 0.3.2
```

## Getting started

### From sources

You cant build from sources cloning this repository and runing the following command:

```shell
go build -o gout
```

Then in the same directory run your builded binary:

```shell
./gout -d -m "Some message: '{}' and '{}'" -p "1" -p "2"
```

Finally you will get:

`[2020-10-10 22:58:40.164] - [DEBUG]: Some message: '1' and '2'`

## From binaries

Download release versión, then unzip it. Finally move the binary to `/usr/bin` directory.

To test if all works, run this:

```shell
gout -d -m "Some message: '{}' and '{}'" -p "1" -p "2"
```

## Developing

To develop on this project you will need golang version `go1.15.2` or superior.

To test your code:

```shell
go test github.com/pascencio/goutput/...
```

## TODO

### Release 0.0.1

- [x] First estable version

### Release 0.0.2

- [x] Add continous integration
- [x] Add release generation
- [x] Add date support
- [x] Add output level management based on environment variable

### Release 0.4.0

- [ ] Add standard input integration
- [ ] Testing for other linux distributions
- [ ] Add argument descriptión
- [ ] Create a logo :stuck_out_tongue_winking_eye:
