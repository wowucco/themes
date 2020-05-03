package sword

import (
	adminTemplate "github.com/wowucco/go-admin/template"
	"github.com/wowucco/go-admin/template/components"
	"github.com/wowucco/go-admin/template/types"
	"github.com/GoAdminGroup/themes/common"
	"github.com/GoAdminGroup/themes/sword/resource"
	"github.com/gobuffalo/packr/v2"
	"html/template"
	"strings"
)

type Theme struct {
	ThemeName string
	components.Base
	common.BaseTheme
}

var Sword = Theme{
	ThemeName: "sword",
	Base: components.Base{
		Attribute: types.Attribute{
			TemplateList: TemplateList,
		},
	},
}

func init() {
	adminTemplate.Add("sword", &Sword)
}

func Get() *Theme {
	return &Sword
}

func (t *Theme) Name() string {
	return t.ThemeName
}

func (*Theme) GetTmplList() map[string]string {
	return TemplateList
}

func (*Theme) GetTemplate(isPjax bool) (tmpl *template.Template, name string) {
	var err error

	if !isPjax {
		name = "layout"
		tmpl, err = template.New("layout").Funcs(adminTemplate.DefaultFuncMap).
			Parse(TemplateList["layout"] +
				TemplateList["head"] + TemplateList["header"] + TemplateList["sidebar"] +
				TemplateList["footer"] + TemplateList["js"] + TemplateList["menu"] +
				TemplateList["admin_panel"] + TemplateList["content"])
	} else {
		name = "content"
		tmpl, err = template.New("content").Funcs(adminTemplate.DefaultFuncMap).
			Parse(TemplateList["admin_panel"] + TemplateList["content"])
	}

	if err != nil {
		panic(err)
	}

	return
}

func (*Theme) GetAsset(path string) ([]byte, error) {
	path = strings.Replace(path, "/assets/dist", "", -1)
	box := packr.New("sword", "./resource/assets/dist")
	return box.Find(path)
}

func (*Theme) GetAssetList() []string {
	return resource.AssetsList
}
