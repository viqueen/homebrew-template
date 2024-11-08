package upgrade_node_package

func Task(info PackageInfo) error {
	distro, err := loadPackageDistro(info)
	if err != nil {
		return err
	}
	err = downloadLatestArtifact(distro)
	if err != nil {
		return err
	}
	err = upgradeTapFormula(distro)
	if err != nil {
		return err
	}
	return nil
}
