# Pomogoro

A simple console based Pomodoro timer written in Go. When run, the application will count down in seconds until the session has finished (by default this is 30 minutes).

## Build

To build the application, execute the `makefile` target `build`

``` bash
make build
```

This will generate an executable file `pomogoro` in the project's root directory.

Alternatively, the application can be executed directly form the source code by running:

``` bash
go run main.go
```

## Usage

To run pomogoro, run the executable:

``` bash
./pomogoro
Starting new session
Time remaining: 29:57
```

The default duration for a session is 30 minutes. This can be modified by passing the `-duration` parameter to the application:

``` bash
./pomogoro -duration 1
Starting new session
Time remaining: 00:59
```