package webpack

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

// AssetsMapper maps asset name to file name
type AssetsMapper func(string) string

// NewAssetsMapper creates assets mapper
func NewAssetsMapper(buildPath string) (AssetsMapper, error) {
	assetsManifestPath := path.Join(buildPath, "asset-manifest.json")

	if _, err := os.Stat(assetsManifestPath); os.IsNotExist(err) {
		return func(file string) string {
			return file
		}, nil
	}

	content, err := ioutil.ReadFile(assetsManifestPath)
	if err != nil {
		return nil, err
	}

	var manifest map[string]string

	if err = json.Unmarshal(content, &manifest); err != nil {
		return nil, err
	}

	return func(file string) string {
		return fmt.Sprintf("/%s/%s", buildPath, manifest[file])
	}, nil
}
