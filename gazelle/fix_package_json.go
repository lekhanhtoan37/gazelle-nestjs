package nestjs

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
	"strings"
)

type packageJSON struct {
	Name string `json:"name"`
	Main string `json:"main"`
}

func newPackageJSON(nameAlias string, main string) *packageJSON {
	if main == "" {
		log.Fatalf("main in pkg: %v is required", nameAlias)
	}

	return &packageJSON{
		Name: nameAlias,
		Main: main,
	}
}

func fixPackageJSON(packageJSONPath string, _packageJSON *packageJSON) error {
	oldPackageJSONData, err := os.ReadFile(packageJSONPath)

	if err == nil {
		var oldPackageJSON *packageJSON
		err = json.Unmarshal(oldPackageJSONData, &oldPackageJSON)
		if err != nil {
			log.Printf(Warn("Parse file package.json: %v", err))
			return err
		}

		if oldPackageJSON != nil && oldPackageJSON.Name == _packageJSON.Name &&
			oldPackageJSON.Main == _packageJSON.Main {
			return nil
		}
	}

	data, err := json.MarshalIndent(_packageJSON, "", "  ")
	if err != nil {
		log.Fatalf("Error marshalling package.json: %v", err)
		return err
	}

	data, err = unescapeUnicodeCharactersInJSON(data)
	if err != nil {
		log.Printf("Unescape unicode characters in JSON: %v", err)
		return err
	}

	if err := os.WriteFile(packageJSONPath, data, 0o666); err != nil {
		log.Fatalf("Write file package.json: %v", err)
		return err
	}

	return nil
}

func updateRootPackageJSON(rootPackageJSONPath string, internalPkg map[string]bool) error {
	data, err := os.ReadFile(rootPackageJSONPath)
	if err != nil {
		log.Printf("Read file package.json: %v\n", err)
		return err
	}

	var packageJSON map[string]interface{}
	if err := json.Unmarshal(data, &packageJSON); err != nil {
		log.Printf("Unmarshal package.json: %v\n", err)
		return err
	}

	if _, ok := packageJSON["dependencies"]; !ok {
		log.Println("Not found dependencies in package.json")
		return nil
	}

	dependencies := packageJSON["dependencies"].(map[string]interface{})
	for pkg := range internalPkg {
		if dependencies[pkg] != nil || strings.HasSuffix(pkg, "*") {
			continue
		}

		dependencies[pkg] = "workspace:*"
	}

	data, err = json.MarshalIndent(packageJSON, "", "  ")
	if err != nil {
		log.Printf("Error marshalling new package.json: %v\n", err)
		return err
	}

	data, err = unescapeUnicodeCharactersInJSON(data)
	if err != nil {
		log.Printf("Unescape unicode characters in JSON: %v", err)
		return err
	}

	if err := os.WriteFile(rootPackageJSONPath, data, 0o666); err != nil {
		log.Printf("Write new content in package.json: %v", err)
		return err
	}

	return nil
}

func unescapeUnicodeCharactersInJSON(_jsonRaw json.RawMessage) (json.RawMessage, error) {
	str, err := strconv.Unquote(strings.Replace(strconv.Quote(string(_jsonRaw)), `\\u`, `\u`, -1))
	if err != nil {
		return nil, err
	}
	return []byte(str), nil
}
