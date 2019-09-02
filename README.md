# go-loadbalancer-example

Writing the NodeJS loadbalancer in golang to learn the language. So far it has helped me see changes I can apply in the NodeJS implementation.

First go thru is learning go's idomatic way of doing things. Next is to use packages (ex, viper) for the env vars.

Compile
```bash
go install $(echo $GOPATH)/src/github.com/billymfl/go-loadbalancer-example
```

Run
```bash
KEY=<KEY> $(echo $GOPATH)/bin/go-loadbalancer-example
```
