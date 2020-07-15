package generate

import (
	"errors"
	"fmt"
	"github.com/domgolonka/ent2proto/internal/gen"
	"github.com/domgolonka/ent2proto/internal/load"
	"github.com/facebookincubator/ent/schema/field"
	"os"
	"path"
	"path/filepath"
	"text/template"
)

type Option func(*gen.Config) error

// custom implementation for pflag.
type idType field.Type

// Set implements the Set method of the flag.Value interface.
func (t *idType) Set(s string) error {
	switch s {
	case field.TypeInt.String():
		*t = idType(field.TypeInt)
	case field.TypeInt64.String():
		*t = idType(field.TypeInt64)
	case field.TypeUint.String():
		*t = idType(field.TypeUint)
	case field.TypeUint64.String():
		*t = idType(field.TypeUint64)
	case field.TypeString.String():
		*t = idType(field.TypeString)
	default:
		return errors.New("invalid type")
	}
	return nil
}

// Type returns the type representation of the id option for help command.
func (idType) Type() string {
	return fmt.Sprintf("%v", []field.Type{
		field.TypeInt,
		field.TypeInt64,
		field.TypeUint,
		field.TypeUint64,
		field.TypeString,
	})
}

// String returns the default value for the help command.
func (idType) String() string {
	return field.TypeInt.String()
}
// LoadGraph loads the schema package from the given schema path,
// and constructs a *gen.Graph.
func LoadGraph(schemaPath string, cfg *gen.Config) (*gen.Graph, error) {
	spec, err := (&load.LoadConfig{Path: schemaPath}).Load()
	if err != nil {
		return nil, err
	}
	cfg.Schema = spec.PkgPath
	if cfg.Package == "" {
		// default package-path for codegen is one package
		// before the schema package (`<project>/ent/schema`).
		cfg.Package = path.Dir(spec.PkgPath)
	}
	return gen.NewGraph(cfg, spec.Schemas...)
}


func Generate(schemaPath string, cfg *gen.Config, options ...Option) (err error) {
	if cfg.Target == "" {
		abs, err := filepath.Abs(schemaPath)
		if err != nil {
			return err
		}
		// default target-path for codegen is one dir above
		// the schema.
		cfg.Target = filepath.Dir(abs)
	}



	for _, opt := range options {
		if err := opt(cfg); err != nil {
			return err
		}
	}

	undo, err := gen.PrepareEnv(cfg)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = undo()
		}
	}()

	graph, err := LoadGraph(schemaPath, cfg)
	if err != nil {
		return err
	}


	return graph.Gen()
}

// TemplateFiles parses the named files and associates the resulting templates
// with codegen templates.
func TemplateFiles(filenames ...string) Option {
	return templateOption(func(cfg *gen.Config) (err error) {
		cfg.Template, err = cfg.Template.ParseFiles(filenames...)
		return
	})
}

// TemplateGlob parses the template definitions from the files identified
// by the pattern and associates the resulting templates with codegen templates.
func TemplateGlob(pattern string) Option {
	return templateOption(func(cfg *gen.Config) (err error) {
		cfg.Template, err = cfg.Template.ParseGlob(pattern)
		return
	})
}
// TemplateDir parses the template definitions from the files in the directory
// and associates the resulting templates with codegen templates.
func TemplateDir(path string) Option {
	return templateOption(func(cfg *gen.Config) error {
		return filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return fmt.Errorf("load template: %v", err)
			}
			if info.IsDir() {
				return nil
			}
			cfg.Template, err = cfg.Template.ParseFiles(path)
			return err
		})
	})
}

// templateOption ensures the template instantiate
// once for config and execute the given Option.
func templateOption(next Option) Option {
	return func(cfg *gen.Config) (err error) {
		if cfg.Template == nil {
			cfg.Template = template.New("external").Funcs(gen.Funcs)
		}
		return next(cfg)
	}
}