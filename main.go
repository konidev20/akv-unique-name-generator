package main

import (
	"flag"
	"fmt"

	gonanoid "github.com/matoous/go-nanoid"
)

func main() {
	projectPtr := flag.String("projectName", "", "The name of the project.")
	environmentPtr := flag.String("environmentName", "dev", "The name of the environment. default=\"dev\"")
	regionNamePtr := flag.String("regionName", "centralus", "The name of the region to deploy. default=\"centeralus\"")
	totalSizePtr := flag.Int("size", 24, "The total size of the azure key vault name")
	flag.Parse()

	azureKeyVaultName := GenerateAzureKeyVaultName(*projectPtr, *environmentPtr, *regionNamePtr, *totalSizePtr)
	fmt.Println(azureKeyVaultName)
}

func GenerateAzureKeyVaultName(projectName string, environmentName string, regionName string, totalSize int) string {
	var vaultNamePrefix = projectName + "-" + environmentName + "-" + regionName + "-"
	var vaultNamePrefixSize = len(vaultNamePrefix)
	var guidSize = totalSize - vaultNamePrefixSize
	id := vaultNamePrefix + gonanoid.MustGenerate("abcdefghijklmnopqrstuvwxyz1234567890", guidSize)
	return id
}
