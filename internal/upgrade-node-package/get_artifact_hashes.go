package upgrade_node_package

import (
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"homebrew/internal/helpers"
	"io"
	"os"
	"path/filepath"
)

type artifactHashes struct {
	Sha1   string
	Sha256 string
}

func getArtifactHashes(distro *packageDistro) (*artifactHashes, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("failed to get current working directory: %w", err)
	}

	artifactPath := filepath.Join(cwd, "build", distro.Name, "artifact.tgz")

	file, err := os.Open(artifactPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file for checksum: %w", err)
	}
	defer helpers.GracefulClose(file, helpers.LogError)

	sha1Hash := sha1.New()
	if _, err = io.Copy(sha1Hash, file); err != nil {
		return nil, fmt.Errorf("failed to calculate checksum: %w", err)
	}
	sha1Checksum := hex.EncodeToString(sha1Hash.Sum(nil))

	if _, err = file.Seek(0, io.SeekStart); err != nil {
		return nil, fmt.Errorf("failed to seek to beginning of file: %w", err)
	}

	sha256Hash := sha256.New()
	if _, err = io.Copy(sha256Hash, file); err != nil {
		return nil, fmt.Errorf("failed to calculate checksum: %w", err)
	}
	sha256Checksum := hex.EncodeToString(sha256Hash.Sum(nil))

	return &artifactHashes{
		Sha1:   sha1Checksum,
		Sha256: sha256Checksum,
	}, nil
}
