# Arithmatic API Golang [![Build Status](https://travis-ci.org/Brettm12345/arithmatic.svg?branch=master)](https://travis-ci.org/Brettm12345/arithmatic) [![Go Report Card](https://goreportcard.com/badge/github.com/Brettm12345/arithmatic)](https://goreportcard.com/report/github.com/Brettm12345/arithmatic) [![Coverage Status](https://coveralls.io/repos/github/Brettm12345/arithmatic/badge.svg?branch=master)](https://coveralls.io/github/Brettm12345/arithmatic?branch=master)
A simple arithmatic API with cache support written in Golang

## Building

``` console
$ go get -u github.com/Brettm12345/arithmatic
```

## Usage
I use [httpie](httpie) for the following examples 

### Starting the server

``` console
$ arithmatic
```

### Addition
``` console
$ http localhost:8080/add?x=4&y=5
```

``` json
{
    "Action": "add",
    "Answer": 9,
    "Cached": false,
    "X": 4,
    "Y": 5
}
```

### Subtraction
``` console
$ http localhost:8080/subtract?x=4&y=5
```

``` json
{
    "action": "subtract",
    "answer": -1,
    "cached": false,
    "x": 4,
    "y": 5
}
```

### Multiplication
``` console
$ http localhost:8080/subtract?x=4&y=5
```

``` json
{
    "action": "multiply",
    "answer": 20,
    "cached": false,
    "x": 4,
    "y": 5
}
```

### Division
``` console
$ http localhost:8080/subtract?x=4&y=5
```

``` json
{
    "action": "divide",
    "answer": 0,
    "cached": false,
    "x": 4,
    "y": 5
}
```
