package pbfts

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
	VersionNode        string
	VersionProtoc      string
}

type PbfTs struct {
	command            string
	filePath           string
	githubOrganization string
	githubRepository   string
	versionNode        string
	versionProtoc      string
}

func New(config Config) (*PbfTs, error) {
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
	if config.VersionNode == "" {
		return nil, tracer.Maskf(invalidConfigError, "%T.VersionNode must not be empty", config)
	}
	if config.VersionProtoc == "" {
		return nil, tracer.Maskf(invalidConfigError, "%T.VersionProtoc must not be empty", config)
	}

	p := &PbfTs{
		command:            config.Command,
		filePath:           config.FilePath,
		githubOrganization: config.GithubOrganization,
		githubRepository:   config.GithubRepository,
		versionNode:        config.VersionNode,
		versionProtoc:      config.VersionProtoc,
	}

	return p, nil
}

func (p *PbfTs) Workflow() ([]byte, error) {
	b, err := p.render(templateWorkflow)
	if err != nil {
		return nil, tracer.Mask(err)
	}

	return b, nil
}

func (p *PbfTs) data() interface{} {
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
			Checkout:  version.Checkout,
			Node:      p.versionNode,
			Protoc:    p.versionProtoc,
			SetupNode: version.SetupNode,
		},
	}
}

func (p *PbfTs) render(t string) ([]byte, error) {
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
