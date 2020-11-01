package grpcgo

import (
	"bytes"
	"strings"
	"text/template"

	"github.com/xh3b4sd/tracer"
)

type Config struct {
	Command            string
	FilePath           string
	GithubCurrent      string
	GithubOrganization string
	GithubRepository   string
	VersionGolang      string
	VersionProtoc      string
}

type GrpcGo struct {
	command            string
	filePath           string
	githubCurrent      string
	githubOrganization string
	githubRepository   string
	versionGolang      string
	versionProtoc      string
}

func New(config Config) (*GrpcGo, error) {
	if config.Command == "" {
		return nil, tracer.Maskf(invalidConfigError, "%T.Command must not be empty", config)
	}
	if config.FilePath == "" {
		return nil, tracer.Maskf(invalidConfigError, "%T.FilePath must not be empty", config)
	}
	if config.GithubCurrent == "" {
		return nil, tracer.Maskf(invalidConfigError, "%T.GithubCurrent must not be empty", config)
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

	g := &GrpcGo{
		command:            config.Command,
		filePath:           config.FilePath,
		githubCurrent:      config.GithubCurrent,
		githubOrganization: config.GithubOrganization,
		githubRepository:   config.GithubRepository,
		versionGolang:      config.VersionGolang,
		versionProtoc:      config.VersionProtoc,
	}

	return g, nil
}

func (g *GrpcGo) Usage() ([]byte, error) {
	b, err := g.render(usageTemplate)
	if err != nil {
		return nil, tracer.Mask(err)
	}

	return b, nil
}

func (g *GrpcGo) Workflow() ([]byte, error) {
	b, err := g.render(workflowTemplate)
	if err != nil {
		return nil, tracer.Mask(err)
	}

	return b, nil
}

func (g *GrpcGo) data() interface{} {
	type Github struct {
		Current      string
		Organization string
		Repository   string
	}

	type Version struct {
		Golang string
		Protoc string
	}

	type Data struct {
		Command string
		Github  Github
		Version Version
	}

	return Data{
		Command: g.command,
		Github: Github{
			Current:      g.githubCurrent,
			Organization: g.githubOrganization,
			Repository:   g.githubRepository,
		},
		Version: Version{
			Golang: g.versionGolang,
			Protoc: g.versionProtoc,
		},
	}
}

func (g *GrpcGo) render(t string) ([]byte, error) {
	f := template.FuncMap{
		"ToUpper": func(s string) string {
			n := s

			n = strings.ToUpper(n)
			n = strings.ReplaceAll(n, "-", "")

			return n
		},
	}

	s, err := template.New(g.filePath).Funcs(f).Parse(t)
	if err != nil {
		return nil, tracer.Mask(err)
	}

	var b bytes.Buffer
	err = s.ExecuteTemplate(&b, g.filePath, g.data())
	if err != nil {
		return nil, tracer.Mask(err)
	}

	return b.Bytes(), nil
}
