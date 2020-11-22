# go-monit

A humble clone of the Monit Service Manager

### About
`go-monit` is a humble clone of [Monit](https://en.wikipedia.org/wiki/Monit). At 
the moment it can monitor a single HTTP service defined in `main.go`

_By default a service on [http://localhost:12345](http://localhost:12345) is monitored_

### Usage

#### Build a binary 

**Build a local binary (automatically matches the host OS)**

```shell
make 
```

**Build a linux binary**

```shell
make build.linux
```

#### Run

**Run in the background**

```shell script
./build/go-monit > logfile 2>&1 &
```

**See logs**

```shell script
tail -f logfile
```

### Tentative Future Plans ( TODOs )

- Ship with a `system.d` service
It would be nice to ship with a system.d service that the user can just install

- Declare services in yaml
At the moment, services must be defined in go code, not great. It would be nice 
if `go-monit` could just read `yaml` files in say `/etc/go-monit` and start to 
monitor the services defined there.
