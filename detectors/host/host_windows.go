// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// +build windows

package host // import "go.opentelemetry.io/contrib/detectors/host"

import (
	"context"
	"x/sys/windows/registry"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.24.0"
)

// Comment
type resourceDetector struct {
}

func NewResourceDetector() resource.Detector {
	return &resourceDetector{}
}

func (detector *resourceDetector) Detect(ctx context.Context) (*resource.Resource, error) {
	key, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Cryptography`, registry.QUERY_VALUE)
	if err != nil {
		return nil, err
	}

	defer key.Close()

	machineId, _, err := key.GetStringValue("MachineGuid")
	if err != nil {
		return nil, err
	}

	attributes := []attribute.KeyValue{
		semconv.HostID(machineId),
	}

	return resource.NewWithAttributes(semconv.SchemaURL, attributes...), nil
}
