package grpcts

import (
	"bytes"
	"strings"
	"text/template"

	"github.com/xh3b4sd/tracer"
)

type Config struct {
	FilePath           string
	GithubOrganization string
	GithubRepository   string
	VersionGolang      string
	VersionGrpcWeb     string
	VersionProtoc      string
}

type GrpcTs struct {
	filePath           string
	githubOrganization string
	githubRepository   string
	versionGolang      string
	versionGrpcWeb     string
	versionProtoc      string
}

func New(config Config) (*GrpcTs, error) {
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
	if config.VersionGrpcWeb == "" {
		return nil, tracer.Maskf(invalidConfigError, "%T.VersionGrpcWeb must not be empty", config)
	}
	if config.VersionProtoc == "" {
		return nil, tracer.Maskf(invalidConfigError, "%T.VersionProtoc must not be empty", config)
	}

	g := &GrpcTs{
		filePath:           config.FilePath,
		githubOrganization: config.GithubOrganization,
		githubRepository:   config.GithubRepository,
		versionGolang:      config.VersionGolang,
		versionGrpcWeb:     config.VersionGrpcWeb,
		versionProtoc:      config.VersionProtoc,
	}

	return g, nil
}

func (g *GrpcTs) Generate() ([]byte, error) {
	f := template.FuncMap{
		"ToUpper": strings.ToUpper,
	}

	t, err := template.New(g.filePath).Funcs(f).Parse(workflowTemplate)
	if err != nil {
		return nil, tracer.Mask(err)
	}

	var b bytes.Buffer
	err = t.ExecuteTemplate(&b, g.filePath, g.data())
	if err != nil {
		return nil, tracer.Mask(err)
	}

	return b.Bytes(), nil
}

func (g *GrpcTs) data() interface{} {
	type Github struct {
		Organization string
		Repository   string
	}

	type Version struct {
		Golang  string
		GrpcWeb string
		Protoc  string
	}

	type Data struct {
		Github  Github
		Version Version
	}

	return Data{
		Github: Github{
			Organization: g.githubOrganization,
			Repository:   g.githubRepository,
		},
		Version: Version{
			Golang:  g.versionGolang,
			GrpcWeb: g.versionGrpcWeb,
			Protoc:  g.versionProtoc,
		},
	}
}
