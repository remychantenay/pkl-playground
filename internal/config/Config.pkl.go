// Code generated from Pkl module `remychantenay.pkl_playground.internal.config`. DO NOT EDIT.
package config

import (
	"context"

	"github.com/apple/pkl-go/pkl"
)

type Config struct {
	// SomeInteger is just a simple integer.
	SomeInteger int `pkl:"someInteger"`

	// SomeString is a simple string.
	SomeString string `pkl:"someString"`

	// window contains the window configuration.
	Window *Window `pkl:"window"`

	// Tabs contains the list of tabs.
	Tabs any `pkl:"tabs"`

	// Components is a map of components.
	Components any `pkl:"components"`
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a Config
func LoadFromPath(ctx context.Context, path string) (ret *Config, err error) {
	evaluator, err := pkl.NewEvaluator(ctx, pkl.PreconfiguredOptions)
	if err != nil {
		return nil, err
	}
	defer func() {
		cerr := evaluator.Close()
		if err == nil {
			err = cerr
		}
	}()
	ret, err = Load(ctx, evaluator, pkl.FileSource(path))
	return ret, err
}

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a Config
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (*Config, error) {
	var ret Config
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
