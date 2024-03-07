package nestjs

import (
	"encoding/json"
	"log"
	"os"
	"path"
)

type NestCliConfig struct {
	Schema          string              `json:"$schema"`
	Collection      string              `json:"collection"`
	SourceRoot      string              `json:"sourceRoot"`
	CompilerOptions *CompilerOptions    `json:"compilerOptions"`
	Monorepo        bool                `json:"monorepo"`
	Root            string              `json:"root"`
	Projects        map[string]*Project `json:"projects"`
}

type CompilerOptions struct {
	DeleteOutDir bool   `json:"deleteOutDir"`
	Webpack      bool   `json:"webpack"`
	TsConfigPath string `json:"tsConfigPath"`
}

type NestTranspiler string

const (
	webpack NestTranspiler = "webpack"
	swc     NestTranspiler = "swc"
	tsc     NestTranspiler = "tsc"
)

type Project struct {
	Name         string
	Rel          string
	TsConfig     *TsConfig
	TsConfigRel  string
	TsConfigPath string

	// Default generated by nest cli
	Type            string `json:"type"`
	Root            string `json:"root"`
	EntryFile       string `json:"entryFile"`
	SourceRoot      string `json:"sourceRoot"`
	CompilerOptions struct {
		TsConfigPath string `json:"tsConfigPath"`
	} `json:"compilerOptions"`
}

func (p *Project) ParseTsConfig(cwd string) *TsConfig {
	var tsConfig *TsConfig

	if p.TsConfigRel == "" {
		return tsConfig
	}

	tsConfigPath := path.Join(cwd, p.TsConfigRel)
	data, err := os.ReadFile(tsConfigPath)
	if err != nil {
		log.Printf("Parse tsconfig for project failed: %v, err: %v \n", p.Name, err)
		return tsConfig
	}

	err = json.Unmarshal(data, &tsConfig)

	if err != nil {
		log.Printf("Parse tsconfig then unmarshal failed for project: %v, err: %v \n", p.Name, err)
		return tsConfig
	}

	p.TsConfig = tsConfig

	return tsConfig
}