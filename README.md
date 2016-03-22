# go-mobile-framework
Example framework for the Go mobile talk

## Installation
$ go get golang.org/x/mobile/cmd/gomobile
$ gomobile init

## Run Tests
$ go test -v ./...

## Build iOS
$ gomobile bind -target ios -o mobilesdk.framework

## Build Android
$ gomobile bind -target android -o mobilesdk.aar
