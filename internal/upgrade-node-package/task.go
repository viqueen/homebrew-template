package upgrade_node_package

func Task(info PackageInfo) error {
	distro, err := loadPackageDistro(info)
	if err != nil {
		return err
	}
	return downloadLatestArtifact(distro)
}
