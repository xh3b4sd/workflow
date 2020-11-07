package grpcts

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
	VersionGrpcWeb     string
	VersionNode        string
	VersionProtoc      string
}

type GrpcTs struct {
	command            string
	filePath           string
	githubCurrent      string
	githubOrganization string
	githubRepository   string
	versionGolang      string
	versionGrpcWeb     string
	versionNode        string
	versionProtoc      string
}

func New(config Config) (*GrpcTs, error) {
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
	if config.VersionGrpcWeb == "" {
		return nil, tracer.Maskf(invalidConfigError, "%T.VersionGrpcWeb must not be empty", config)
	}
	if config.VersionNode == "" {
		return nil, tracer.Maskf(invalidConfigError, "%T.VersionNode must not be empty", config)
	}
	if config.VersionProtoc == "" {
		return nil, tracer.Maskf(invalidConfigError, "%T.VersionProtoc must not be empty", config)
	}

	g := &GrpcTs{
		command:            config.Command,
		filePath:           config.FilePath,
		githubCurrent:      config.GithubCurrent,
		githubOrganization: config.GithubOrganization,
		githubRepository:   config.GithubRepository,
		versionGolang:      config.VersionGolang,
		versionGrpcWeb:     config.VersionGrpcWeb,
		versionNode:        config.VersionNode,
		versionProtoc:      config.VersionProtoc,
	}

	return g, nil
}

func (g *GrpcTs) Usage() ([]byte, error) {
	b, err := g.render(usageTemplate)
	if err != nil {
		return nil, tracer.Mask(err)
	}

	return b, nil
}

func (g *GrpcTs) Workflow() ([]byte, error) {
	b, err := g.render(workflowTemplate)
	if err != nil {
		return nil, tracer.Mask(err)
	}

	return b, nil
}

func (g *GrpcTs) data() interface{} {
	type Github struct {
		Current      string
		Organization string
		Repository   string
	}

	type Version struct {
		Golang  string
		GrpcWeb string
		Node    string
		Protoc  string
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
			Golang:  g.versionGolang,
			GrpcWeb: g.versionGrpcWeb,
			Node:    g.versionNode,
			Protoc:  g.versionProtoc,
		},
	}
}

func (g *GrpcTs) render(t string) ([]byte, error) {
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
