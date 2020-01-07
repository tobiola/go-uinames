# go-uinames

[![GoDoc](https://godoc.org/github.com/tobiola/go-uinames?status.svg)](https://godoc.org/github.com/tobiola/go-uinames)
[![Build Status](https://travis-ci.com/tobiola/go-uinames.svg?branch=master)](https://travis-ci.com/tobiola/go-uinames)
[![Go Report Card](https://goreportcard.com/badge/github.com/tobiola/go-uinames)](https://goreportcard.com/report/github.com/tobiola/go-uinames)

go-uinames is a Go wrapper for [https://uinames.com](https://uinames.com)

## Getting Started

````Go
// Generate a random Name struct
name, _ := GetName(nil)

// Generate an array of 5 NameExtra structs from the region United States
names, _ := GetNames(&Options{Region: "United States", Amount: 5})
````
