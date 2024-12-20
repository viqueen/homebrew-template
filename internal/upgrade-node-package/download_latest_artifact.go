package upgrade_node_package

import (
	"fmt"
	"homebrew/internal/helpers"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func downloadLatestArtifact(distro *packageDistro) error {
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current working directory: %w", err)
	}
	buildPath := filepath.Join(cwd, "build", distro.Name)
	if err = os.MkdirAll(buildPath, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create build directory: %w", err)
	}

	artifactPath := filepath.Join(buildPath, "artifact.tgz")

	// Start downloading the tarball
	resp, err := http.Get(distro.TarBall)
	if err != nil {
		return fmt.Errorf("failed to download artifact: %w", err)
	}
	defer helpers.GracefulClose(resp.Body, helpers.LogError)

	// Create the file to save the downloaded artifact
	file, err := os.Create(artifactPath)
	if err != nil {
		return fmt.Errorf("failed to create artifact file: %w", err)
	}
	defer helpers.GracefulClose(file, helpers.LogError)

	// Write the response body to the file
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		// Remove the artifact if copying fails
		_ = os.Remove(artifactPath)
		return fmt.Errorf("failed to save artifact: %w", err)
	}

	return nil
}
