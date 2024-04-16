// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build macos
// +build macos

package host // import "go.opentelemetry.io/contrib/detectors/host"

import (
	"context"
	"os"
	"strings"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.24.0"
)

type resourceDetector struct{}

// NewResourceDetector returns a resource detector that will detect host resources.
func NewResourceDetector() resource.Detector {
	return &resourceDetector{}
}

// Detect detects associated resources when running on a physical host.
func (detector *resourceDetector) Detect(ctx context.Context) (*resource.Resource, error) {
	machineId, err := os.ReadFile("/etc/machine-id")
	if err != nil {
		return nil, err
	}

	attributes := []attribute.KeyValue{
		semconv.HostID(strings.Trim(string(machineId), "\n")),
	}

	return resource.NewWithAttributes(semconv.SchemaURL, attributes...), nil
}
