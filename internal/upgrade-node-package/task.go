package upgrade_node_package

import "log"

func Task(info PackageInfo) error {
	log.Printf("Upgrading %s/%s", info.Org, info.Name)
	distro, err := getPackageDistro(info)
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
