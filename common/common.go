package common

import (
	"github.com/wowucco/go-admin/modules/config"
	"html/template"
)

type BaseTheme struct{}

const Version = "v0.0.31"

func (BaseTheme) GetVersion() string {
	return Version
}

func (BaseTheme) GetRequirements() []string {
	return []string{"v1.2.9"}
}

var requireAssetHTML = map[string]string{
	"datatable": "/assets/dist/js/datatable.min.js",
	"form":      "/assets/dist/js/form.min.js",
	"tree":      "/assets/dist/js/tree.min.js",
}

func (BaseTheme) GetAssetImportHTML(exclude ...string) template.HTML {
	res := template.HTML("")
	for name, html := range requireAssetHTML {
		exist := false
		for i := 0; i < len(exclude); i++ {
			if name == exclude[i] {
				exist = true
			}
		}
		if !exist {
			if config.GetAssetUrl() != "" {
				res += template.HTML(`<script src="` + config.GetAssetUrl() + html + `"></script>`)
			} else {
				res += template.HTML(`<script src="` + config.Url(html) + `"></script>`)
			}
		}
	}
	return res
}
