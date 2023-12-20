//go:build linux

// Package mcp300x implements a sensor model supporting mcp300x adc sensor.
package mcp300x

import (
	"context"

	"go.viam.com/rdk/components/sensor"
	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/resource"
)

var logger = logging.NewDebugLogger("mcp300x")

// registering the component model on init is how we make sure the new model is picked up and usable.
func init() {
	resource.RegisterComponent(
		sensor.API,
		resource.DefaultModelFamily.WithModel("mcp300x"),
		resource.Registration[sensor.Sensor, resource.NoNativeConfig]{Constructor: func(
			ctx context.Context,
			deps resource.Dependencies,
			conf resource.Config,
			logger logging.Logger,
		) (sensor.Sensor, error) {
			return newMcp300x(conf.ResourceName()), nil
		}})
}

func newMcp300x(name resource.Name) sensor.Sensor {
	return &mcp300x{
		Named: name.AsNamed(),
	}
}

// mySensor is a sensor device that always returns "hello world".
type mcp300x struct {
	resource.Named
	resource.AlwaysRebuild
	resource.TriviallyCloseable
}

// Readings always returns "hello world".
func (s *mcp300x) Readings(ctx context.Context, _ map[string]interface{}) (map[string]interface{}, error) {
	return map[string]interface{}{"hello": "world"}, nil
}
