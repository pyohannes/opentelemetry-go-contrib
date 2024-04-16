// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package host

import (
	"context"
	"os"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.24.0"
)

func Test_DetectLinux(t *testing.T) {
	detector := NewResourceDetector()

	hostResource, err := detector.Detect(context.Background())

	assert.True(t, err == nil)

	machineId, _ := getHostId()
	hostName, _ := os.Hostname()

	expectedResource := resource.NewWithAttributes(semconv.SchemaURL, []attribute.KeyValue{
		semconv.HostArchKey.String(runtime.GOARCH),
		semconv.HostName(hostName),
		semconv.HostID(machineId),
	}...)

	assert.Equal(t, expectedResource, hostResource)
}

func Test_DetectLinux_WithOptIns(t *testing.T) {
	detector := NewResourceDetector(
		WithIPAddresses(),
		WithMACAddresses(),
	)

	hostResource, err := detector.Detect(context.Background())

	assert.True(t, err == nil)

	machineId, _ := getHostId()
	hostName, _ := os.Hostname()

	expectedResource := resource.NewWithAttributes(semconv.SchemaURL, []attribute.KeyValue{
		semconv.HostArchKey.String(runtime.GOARCH),
		semconv.HostName(hostName),
		semconv.HostID(machineId),
		semconv.HostIP(getIPAddresses()...),
		semconv.HostMac(getMACAddresses()...),
	}...)

	assert.Equal(t, expectedResource, hostResource)
}
