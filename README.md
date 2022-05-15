# Request-Dumper

Do you need to take a look how a request looks like? Then you need any kind of server application which receives the request and presents it to you.
There for I created the "Request-Dumper".

## Run it locally

```shell
$ rd --help
The request-dumper is a very fast webserver that dumps all incoming requests into his log

Usage:
  request-dumper [flags]

Flags:
  -h, --help       help for request-dumper
  -p, --port int   port on which the application listens for new requests (default 3000)
```

## Run it in a container

```shell
docker run rd gcri.io/florianrusch/request-dumper/rd:1.0.0
```

## Run it in kubernetes

```shell
kubectl run request-dumper gcri.io/florianrusch/request-dumper/rd:1.0.0
kubectl pod request-dumper --port=3000 --name request-dumper
```

Cleanup:

```shell
kubectl delete pod request-dumper
kubectl delete svc request-dumper
```

## ToDo's

- [ ] Optimize/improve CI pipeline
	- <https://github.dev/cli/cli>
	- <https://github.dev/traefik/traefik>
	- <https://github.com/goreleaser/goreleaser-action>
- [ ] Make output format customizable (plain text, json)
- [ ] Implement one-request mode (app exits automatically after the request was received)
- [ ] Create a homebrew formular

## License

request-dumper is free and unencumbered public domain software. For more information, see <https://unlicense.org/> or the accompanying [LICENSE](/LICENSE) file.
