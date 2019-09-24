module github.com/biggestT/faas-gateway

go 1.13

require (
	github.com/Microsoft/go-winio v0.4.14 // indirect
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/docker v1.13.1
	github.com/docker/engine-api v0.4.0 // indirect
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/gogo/protobuf v1.3.0 // indirect
	github.com/opencontainers/go-digest v1.0.0-rc1 // indirect
	github.com/opencontainers/image-spec v1.0.1 // indirect
	golang.org/x/net v0.0.0-20190921015927-1a5e07d1ff72 // indirect
)

replace github.com/docker/docker v1.13.1 => github.com/docker/engine v1.4.2-0.20180816081446-320063a2ad06
