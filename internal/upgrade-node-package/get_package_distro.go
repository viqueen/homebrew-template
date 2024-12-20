package upgrade_node_package

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

type packageSignature struct {
	KeyId     string `json:"keyid"`
	Signature string `json:"sig"`
}

type packageDistro struct {
	Name         string             `json:"name"`
	Integrity    string             `json:"integrity"`
	ShaSum       string             `json:"shasum"`
	TarBall      string             `json:"tarball"`
	FileCount    int                `json:"fileCount"`
	UnpackedSize int                `json:"unpackedSize"`
	Signatures   []packageSignature `json:"signatures"`
}

type PackageInfo struct {
	Org  string
	Name string
}

func getPackageDistro(info PackageInfo) (*packageDistro, error) {
	cmd := exec.Command("npm", "show", fmt.Sprintf("%s/%s", info.Org, info.Name), "dist", "--json")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to get package distro: %w", err)
	}
	var distro packageDistro
	if err = json.Unmarshal(output, &distro); err != nil {
		return nil, fmt.Errorf("failed to unmarshal package distro: %w", err)
	}
	distro.Name = info.Name
	return &distro, nil
}
