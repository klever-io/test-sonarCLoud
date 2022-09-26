# Logging

A gRPC Server that uses a lib from **kloud-sdk** for **logging**

## Library

[klever-sdk/logging](https://github.com/klever-io/kloud-sdk/tree/main/logging)

Feel free to open pull requests or ask for changes

## Considerations

* For gRPC Server and Client **ALWAYS** use **ZAP** interceptors from [go-grpc-middleware](https://github.com/grpc-ecosystem/go-grpc-middleware/tree/master/logging/zap)

```
serverOpts := make([]grpc.ServerOption, 0)
serverOpts = append(serverOpts,
  grpc.ChainUnaryInterceptor(grpc_middleware.ChainUnaryServer(
    grpc_zap.UnaryServerInterceptor(logger),
  )),
  grpc.ChainStreamInterceptor(grpc_middleware.ChainStreamServer(
    grpc_zap.StreamServerInterceptor(logger),
  )),
)
```

* default **INFO** logs at the end of each called method are inject by logging interceptor

```
2022-05-30T02:37:08.060284016-03:00     INFO    zap/options.go:212   finished unary call with code OK        {"grpc.start_time": "2022-05-30T02:37:08-03:00", "system": "grpc", "span.kind": "server", "grpc.service": "helloworld.Greeter", "grpc.method": "SayHello", "grpc.code": "OK", "grpc.time_ms": 0.11}
```

* You will be also able to Extract a **LOGGER** from context:
```
logger := ctxzap.Extract(ctx)
logger.Info("SayHello INFO message", zap.String("name", in.Name))
```

* The **LIBRARY** will **AUTOMATICALLY** check if application is running on cloud or locally and will build the correct config based on that

* You can add fields to the log message and QUERY using **LOGGING** tool from GCP


## Log Level

The default log level is **ERROR** but you can change that setting **LOG_LEVEL** env variable. Supported values are:
* debug **(prints debug, info, warn and error messages)**
* info  **(prints info, warn and error messages)**
* warn  **(prints warn and error messages)**
* error **(prints ONLY error messages)**

## Running this sample

Cloud:
```
# Go to your favorite kubernetes cluster and run:
make deploy
```

Locally:
```
go run main.go
```