// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build darwin
// +build darwin

package host

import (
	"context"
	"os/exec"
	"strings"
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

	machineId, err := exec.Command("ioreg -rd1 -c IOPlatformExpertDevice").Output()

	assert.True(t, err == nil)

	expectedResource := resource.NewWithAttributes(semconv.SchemaURL, []attribute.KeyValue{
		semconv.HostID(strings.Trim(string(machineId), "\n")),
	}...)

	assert.Equal(t, expectedResource, hostResource)
}
