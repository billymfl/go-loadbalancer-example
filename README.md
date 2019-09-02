# go-loadbalancer-example

Writing the Node.JS loadbalancer in golang to learn the language. So far it has helped me see changes I can apply in the Node.JS implementation.

First go thru is learning go's idomatic way of doing things. Next is to use packages (ex, viper) for the env vars.

Compile
```bash
go install $(echo $GOPATH)/src/github.com/billymfl/go-loadbalancer-example
```

Run
```bash
KEY=<KEY> $(echo $GOPATH)/bin/go-loadbalancer-example
```

Build docker image
```bash
docker build -t loadbalancer-go .
```

Start an instance of the app
```bash
docker run --rm -p 80:80 -e "KEY=<KEY>" loadbalancer-go:latest
```

See [Usage - Registering a server](https://github.com/billymfl/loadbalancer/blob/master/README.md#registering) for the Node.JS version for how to use the loadbalancer.

