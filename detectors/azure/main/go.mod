module main

go 1.21.1

replace go.opentelemetry.io/contrib/detectors/azure/vm => ../vm

require go.opentelemetry.io/contrib/detectors/azure/vm v0.0.0-00010101000000-000000000000

require (
	github.com/go-logr/logr v1.4.1 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	go.opentelemetry.io/otel v1.25.0 // indirect
	go.opentelemetry.io/otel/metric v1.25.0 // indirect
	go.opentelemetry.io/otel/sdk v1.25.0 // indirect
	go.opentelemetry.io/otel/trace v1.25.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
)
