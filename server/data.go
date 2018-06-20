package server

import "github.com/kimrgrey/go-create-react-app/webpack"

// ViewData contains data for the view
type ViewData struct {
	assetsMapper webpack.AssetsMapper
}

// NewViewData creates new data for the view
func NewViewData(buildPath string) (ViewData, error) {
	assetsMapper, err := webpack.NewAssetsMapper(buildPath)
	if err != nil {
		return ViewData{}, err
	}

	return ViewData{assetsMapper: assetsMapper}, nil
}

// Webpack maps file name to path
func (d ViewData) Webpack(file string) string {
	return d.assetsMapper(file)
}
