package pbfgo

import (
	"bytes"
	"fmt"
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

func New(c Config) *PbfGo {
	if c.Command == "" {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Command must not be empty", c)))
	}
	if c.FilePath == "" {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.FilePath must not be empty", c)))
	}
	if c.GithubOrganization == "" {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.GithubOrganization must not be empty", c)))
	}
	if c.GithubRepository == "" {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.GithubRepository must not be empty", c)))
	}
	if c.VersionGolang == "" {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.VersionGolang must not be empty", c)))
	}
	if c.VersionProtoc == "" {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.VersionProtoc must not be empty", c)))
	}

	return &PbfGo{
		command:            c.Command,
		filePath:           c.FilePath,
		githubOrganization: c.GithubOrganization,
		githubRepository:   c.GithubRepository,
		versionGolang:      c.VersionGolang,
		versionProtoc:      c.VersionProtoc,
	}
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
