// Code generated from Pkl module `remychantenay.pkl_playground.internal.config`. DO NOT EDIT.
package config

import (
	"context"

	"github.com/apple/pkl-go/pkl"
)

type Config struct {
	// Something is just a simple integer.
	Something int `pkl:"something"`

	// somethingElse is a simple string.
	SomethingElse string `pkl:"somethingElse"`

	// window contains the window configuration.
	Window *Window `pkl:"window"`
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
