# goconfig

This is a simple library that provides a wrapper around [gcfg]() that gives you
easy access to your configuration throughout your application, _and_ provides
facilities either for a default config file or a file path that you pass in when
you run your application.

## Installation

```
go get github.com/sirsean/goconfig
```

## Implement your own config package

```
package config

import (
	"github.com/sirsean/goconfig"
)

type Config struct {
	ThirdPartyApi struct {
		Key string
	}

	Host struct {
		Port string
		Path string
	}
}

var C Config

func init() {
	cl := goconfig.NewConfigLoader("/etc/myapp/myapp.gfcg")
	cl.Load(&C)
}
```

## Use it in your app

```
package main

import (
    "github.com/yourname/myapp/config"
    "log"
)

func main() {
    log.Printf("serving on port: %v", config.C.Host.Port)
}
```

## Override the default filename

```
./myapp special-config.gcfg
```
