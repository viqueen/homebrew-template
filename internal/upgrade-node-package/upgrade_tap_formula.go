package upgrade_node_package

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

type formulaInfo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Homepage    string `json:"homepage"`
	License     string `json:"license"`
}

type templateParams struct {
	Name        string
	Description string
	Homepage    string
	Url         string
	Sha256      string
	License     string
}

func upgradeTapFormula(distro *packageDistro) error {
	hashes, err := getArtifactHashes()
	if err != nil {
		return fmt.Errorf("failed to get artifact hashes: %w", err)
	}

	if hashes.Sha1 != distro.ShaSum {
		return fmt.Errorf("checksum mismatch: %s !== %s", hashes.Sha1, distro.ShaSum)
	}

	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current working directory: %w", err)
	}

	formulaPath := filepath.Join(cwd, "node-formulas", fmt.Sprintf("%s.json", distro.Name))

	formulaInfoContent, err := os.ReadFile(formulaPath)
	if err != nil {
		return fmt.Errorf("failed to read formula distro file: %w", err)
	}

	var formulaInfo formulaInfo
	if err = json.Unmarshal(formulaInfoContent, &formulaInfo); err != nil {
		return fmt.Errorf("failed to parse formula distro JSON: %w", err)
	}

	rubyFormulaPath := filepath.Join(cwd, fmt.Sprintf("%s.rb", distro.Name))

	tmpl, err := template.New("rubyFormula").Parse(rubyFormulaTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse Ruby formula template: %w", err)
	}

	params := templateParams{
		Name:        formulaInfo.Name,
		Description: formulaInfo.Description,
		Homepage:    formulaInfo.Homepage,
		Url:         distro.TarBall,
		Sha256:      hashes.Sha256,
		License:     formulaInfo.License,
	}
	var buffer bytes.Buffer
	if err = tmpl.Execute(&buffer, params); err != nil {
		return fmt.Errorf("failed to execute Ruby formula template: %w", err)
	}

	if err = os.WriteFile(rubyFormulaPath, buffer.Bytes(), 0644); err != nil {
		return fmt.Errorf("failed to write formula file: %w", err)
	}

	return nil
}

var rubyFormulaTemplate = `require "language/node"

class {{.Name}} < Formula
	desc "{{.Description}}"
	homepage "{{.Homepage}}"
	url "{{.Url}}"
	sha256 "{{.Sha256}}"
	license "{{.License}}"

	depends_on "node"

	def install
    	system "npm", "install", *Language::Node.std_npm_install_args(libexec)
    	bin.install_symlink Dir["#{libexec}/bin/*"]
  	end

  	test do
    	system "false"
  	end
end
`
