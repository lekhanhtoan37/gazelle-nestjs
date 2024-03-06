package nestjs

type TsConfig struct {
	Extends         string                  `json:"extends"`
	CompilerOptions TsConfigCompilerOptions `json:"compilerOptions"`
	Include         []string                `json:"include"`
	Exclude         []string                `json:"exclude"`
}

type TsConfigCompilerOptions struct {
	Module                           string              `json:"module"`
	Declaration                      bool                `json:"declaration"`
	RemoveComments                   bool                `json:"removeComments"`
	EmitDecoratorMetadata            bool                `json:"emitDecoratorMetadata"`
	ExperimentalDecorators           bool                `json:"experimentalDecorators"`
	AllowSyntheticDefaultImports     bool                `json:"allowSyntheticDefaultImports"`
	Target                           string              `json:"target"`
	SourceMap                        bool                `json:"sourceMap"`
	OutDir                           string              `json:"outDir"`
	BaseURL                          string              `json:"baseUrl"`
	Incremental                      bool                `json:"incremental"`
	SkipLibCheck                     bool                `json:"skipLibCheck"`
	StrictNullChecks                 bool                `json:"strictNullChecks"`
	NoImplicitAny                    bool                `json:"noImplicitAny"`
	StrictBindCallApply              bool                `json:"strictBindCallApply"`
	ForceConsistentCasingInFileNames bool                `json:"forceConsistentCasingInFileNames"`
	NoFallthroughCasesInSwitch       bool                `json:"noFallthroughCasesInSwitch"`
	Paths                            map[string][]string `json:"paths"`
}
