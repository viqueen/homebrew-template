package upgrade_node_package

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"homebrew/internal/helpers"
	"io"
	"log"
	"os"
	"path/filepath"
)

func runChecksum(distro *PackageDistro) error {
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current working directory: %w", err)
	}

	target := filepath.Join(cwd, "build", "artifact.tgz")

	// Open the file
	file, err := os.Open(target)
	if err != nil {
		return fmt.Errorf("failed to open file for checksum: %w", err)
	}
	defer helpers.GracefulClose(file, helpers.LogError)

	// Calculate SHA1 checksum
	hash := sha1.New()
	if _, err = io.Copy(hash, file); err != nil {
		return fmt.Errorf("failed to calculate checksum: %w", err)
	}
	computedChecksum := hex.EncodeToString(hash.Sum(nil))

	log.Printf("provided checksum: %s", distro.ShaSum)
	log.Printf("computed checksum: %s", computedChecksum)

	// Compare with the provided checksum
	if computedChecksum == distro.ShaSum {
		return nil
	}
	return fmt.Errorf("checksum mismatch: %s !== %s", computedChecksum, distro.ShaSum)
}
