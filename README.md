# go-autumn-logging-log

Plain golang logging plugin for go-autumn-logging.

## About go-autumn

A collection of libraries for [enterprise microservices](https://github.com/StephanHCB/go-mailer-service/blob/master/README.md) in golang that
- is heavily inspired by Spring Boot / Spring Cloud
- is very opinionated
- names modules by what they do
- unlike Spring Boot avoids certain types of auto-magical behaviour
- is not a library monolith, that is every part only depends on the api parts of the other components
  at most, and the api parts do not add any dependencies.  

Fall is my favourite season, so I'm calling it go-autumn.

## About go-autumn-logging

An interface for pluggable logging support with context integration, 
see [go-autumn-logging](https://github.com/StephanHCB/go-autumn-logging).

## About go-autumn-logging-log

Implementation that pulls in logging using the plain golang log standard library.

Very minimal and not very versatile, but good for command line tools etc.

### Limitations

* Only partially context aware (because plain golang standard library logging has no support for it)
* Does not provide a context injection middleware (same reason)
* Not optimized for performance or throughput

### Features

* minimal dependency footprint (just the golang standard library)
* supports additional fields and errors
* supports log levels
* supports logging requestIds, if `aulogging.RequestIdRetriever` has been assigned

## Usage

### Library Authors

Do NOT import this. You are doing it WRONG. You'll take away the application author's choice of 
logging libraries, which is the whole point of using go-autumn-logging in the first place.

Use [go-autumn-logging](https://github.com/StephanHCB/go-autumn-logging)!

### Application Authors

You're all set with this dependency.

### How To Use

Use call chaining style:

```
import "github.com/StephanHCB/go-autumn-logging"

func someFunction(ctx context.Context) {
    aulogging.Logger.Ctx(ctx).Warn().Print("something bad has happened")
}
```

or if you prefer (does the same thing):

```
import "github.com/StephanHCB/go-autumn-logging"

func someFunction(ctx context.Context) {
    aulogging.Warn(ctx, "something bad has happened")
}
```
