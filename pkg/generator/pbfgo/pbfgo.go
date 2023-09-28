package pbfgo

import (
	"bytes"
	"strings"
	"text/template"

	"github.com/xh3b4sd/tracer"
	"github.com/xh3b4sd/workflow/pkg/version"
)

type Config struct {
	Command            string
	FilePath           string
	GithubOrganization string
	GithubRepository   string
	VersionGolang      string
	VersionProtoc      string
}

type PbfGo struct {
	command            string
	filePath           string
	githubOrganization string
	githubRepository   string
	versionGolang      string
	versionProtoc      string
}

func New(config Config) (*PbfGo, error) {
	if config.Command == "" {
		return nil, tracer.Maskf(invalidConfigError, "%T.Command must not be empty", config)
	}
	if config.FilePath == "" {
		return nil, tracer.Maskf(invalidConfigError, "%T.FilePath must not be empty", config)
	}
	if config.GithubOrganization == "" {
		return nil, tracer.Maskf(invalidConfigError, "%T.GithubOrganization must not be empty", config)
	}
	if config.GithubRepository == "" {
		return nil, tracer.Maskf(invalidConfigError, "%T.GithubRepository must not be empty", config)
	}
	if config.VersionGolang == "" {
		return nil, tracer.Maskf(invalidConfigError, "%T.VersionGolang must not be empty", config)
	}
	if config.VersionProtoc == "" {
		return nil, tracer.Maskf(invalidConfigError, "%T.VersionProtoc must not be empty", config)
	}

	p := &PbfGo{
		command:            config.Command,
		filePath:           config.FilePath,
		githubOrganization: config.GithubOrganization,
		githubRepository:   config.GithubRepository,
		versionGolang:      config.VersionGolang,
		versionProtoc:      config.VersionProtoc,
	}

	return p, nil
}

func (p *PbfGo) Workflow() ([]byte, error) {
	b, err := p.render(templateWorkflow)
	if err != nil {
		return nil, tracer.Mask(err)
	}

	return b, nil
}

func (p *PbfGo) data() interface{} {
	type Github struct {
		Organization string
		Repository   string
	}

	type Data struct {
		Command string
		Github  Github
		Version version.Version
	}

	return Data{
		Command: p.command,
		Github: Github{
			Organization: p.githubOrganization,
			Repository:   p.githubRepository,
		},
		Version: version.Version{
			Checkout: version.Checkout,
			Golang:   p.versionGolang,
			Protoc:   p.versionProtoc,
			SetupGo:  version.SetupGo,
		},
	}
}

func (p *PbfGo) render(t string) ([]byte, error) {
	f := template.FuncMap{
		"ToUpper": func(s string) string {
			n := s

			n = strings.ToUpper(n)
			n = strings.ReplaceAll(n, "-", "")

			return n
		},
	}

	s, err := template.New(p.filePath).Funcs(f).Parse(t)
	if err != nil {
		return nil, tracer.Mask(err)
	}

	var b bytes.Buffer
	err = s.ExecuteTemplate(&b, p.filePath, p.data())
	if err != nil {
		return nil, tracer.Mask(err)
	}

	return b.Bytes(), nil
}
