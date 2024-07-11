# buf

Take a [tour](https://buf.build/docs/introduction) using buf v2.

## Usage

### 1.Define your [gRPC](https://grpc.io/docs/) service using protocol buffers

`your_service.proto`:

```protobuf
 syntax = "proto3";
 package your.service.v1;
 option go_package = "github.com/yourorg/yourprotos/gen/go/your/service/v1";

 message StringMessage {
   string value = 1;
 }

 service YourService {
   rpc Echo(StringMessage) returns (StringMessage) {}
 }
```

### 2. Generate gRPC stubs

This step generates the gRPC stubs that you can use to implement the service and consume from clients:

Here's an example `buf.gen.yaml` you can use to generate the stubs with [buf](https://github.com/bufbuild/buf):

```yaml
version: v2
plugins:
  - local: go
    out: gen/go
    opt:
      - paths=source_relative
  - local: go-grpc
    out: gen/go
    opt:
      - paths=source_relative
```

With this file in place, you can generate your files using `buf generate`.

> For a complete example of using `buf generate` to generate protobuf stubs, see
> [the boilerplate repo](https://github.com/johanbrandhorst/grpc-gateway-boilerplate).
> For more information on generating the stubs with buf, see
> [the official documentation](https://docs.buf.build/generate-usage).

Note: Just remember choose right version, I'm using buf v2 currently.

See [buf.local.yaml](buf.local.yaml) and [here](https://buf.build/docs/generate/overview#generating-with-local-plugins).

## Usage with remote plugins

As an alternative to all of the above, you can use buf with remote plugins to manage plugin versions and generation. An example buf.gen.yaml using remote plugin generation looks like this:

```yaml
version: v2
plugins:
  - remote: buf.build/grpc-ecosystem/gateway:v2.20.0
    out: gen/go
    opt: paths=source_relative
  - remote: buf.build/community/google-gnostic-openapi:v0.7.0
    out: gen/go
    opt: paths=source_relative
  - remote: buf.build/bufbuild/validate-go:v1.0.4
    out: gen/go
    opt: paths=source_relative
  - remote: buf.build/grpc/go
    out: gen/go
    opt: paths=source_relative
  - remote: buf.build/protocolbuffers/go
    out: gen/go
    opt: paths=source_relative
```

This requires no local installation of any plugins. Be careful to use the same
version of the generator as the runtime library, i.e. if using `v2.16.2`, run

```shell
$ go get github.com/grpc-ecosystem/grpc-gateway/v2@v2.16.2
```

To get the same version of the runtime in your `go.mod`.

Note that usage of remote plugins is incompatible with usage of external configuration files like [grpc_api_configuration](https://grpc-ecosystem.github.io/grpc-gateway/docs/mapping/grpc_api_configuration/#using-an-external-configuration-file).

See [buf.remote.yaml](buf.remote.yaml) and more [here](https://buf.build/docs/generate/remote-plugins).


## Some tips and tricks for starter

- `buf.yaml`

First, you can run `buf config init` to create [buf.yaml](buf.yaml) file.

It defines a workspace and the configurations for each module within it. It's the primary configuration file, and defines each module's directory, name, lint and breaking configurations, and any files to exclude, along with the workspace's shared dependencies.

```yaml
version: v2
modules:
  - path: proto
lint:
  use:
    - DEFAULT
breaking:
  use:
    - FILE

deps:
  - buf.build/googleapis/googleapis
  - buf.build/kratos/apis
  - buf.build/gnostic/gnostic
```

- `buf.lock`

Second, you can run `buf dep update` to create [buf.lock](buf.lock) file

It contains the workspace's dependency manifest, and represents a single, reproducible build of its dependencies.

- `buf.gen.yaml`

Third, you can run `touch buf.gen.yaml` to create [buf.gen.yaml](buf.gen.yaml) file.

It defines the set of code generation plugins, their options, and the inputs used by the buf generate command to generate code from your Protobuf files. It also allows you to enable and configure managed mode.

