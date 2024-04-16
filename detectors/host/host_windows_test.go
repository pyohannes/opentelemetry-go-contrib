// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build windows
// +build windows

package host

import (
	"context"
	"testing"
	"golang.org/x/sys/windows/registry"

	"github.com/stretchr/testify/assert"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.24.0"
)

func Test_DetectLinux(t *testing.T) {
	detector := NewResourceDetector()

	hostResource, err := detector.Detect(context.Background())

	assert.True(t, err == nil)

	key, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Cryptography`, registry.QUERY_VALUE)

	assert.True(t, err == nil)

	defer key.Close()

	machineId, _, err := key.GetStringValue("MachineGuid")

	assert.True(t, err == nil)

	expectedResource := resource.NewWithAttributes(semconv.SchemaURL, []attribute.KeyValue{
		semconv.HostID(machineId),
	}...)

	assert.Equal(t, expectedResource, hostResource)
}
