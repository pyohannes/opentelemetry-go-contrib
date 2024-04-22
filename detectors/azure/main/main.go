package main

import (
	"context"
	"fmt"
	"go.opentelemetry.io/contrib/detectors/azure/vm"
)

func main() {
	detector := vm.NewResourceDetector()

	azureResource , err := detector.Detect(context.Background())

	if err != nil {
		fmt.Println("error raised")
		return
	}

	for _, attr := range azureResource.Attributes() {
		fmt.Println("Attribute %s: %s", attr.Key, attr.Value)
	}
}
