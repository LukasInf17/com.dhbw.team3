package actions

import (
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"log"
	"strings"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/packr"
)

var r *render.Engine
var assetsBox = packr.NewBox("../public")

func init() {
	r = render.New(render.Options{
		// HTML layout to be used for all HTML requests:
		HTMLLayout: "application.html",

		// Box containing all of the templates:
		TemplatesBox: packr.NewBox("../templates"),
		AssetsBox:    assetsBox,

		// Add template helpers here:
		Helpers: render.Helpers{
			// uncomment for non-Bootstrap form helpers:
			// "form":     plush.FormHelper,
			// "form_for": plush.FormForHelper,
		},
	})
}

// SRIHandler adds support for Subresource integrity
func SRIHandler(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		jsonstring := r.AssetsBox.Bytes("assets/manifest.json")
		var m map[string][]interface{}
		json.Unmarshal(jsonstring, m)
		for k, v := range m {
			if strings.Contains(k, ".css") || strings.Contains(k, ".js") {
				filename := v[0].(string)
				file := r.AssetsBox.Bytes("assets/" + filename)
				sha384 := sha512.New384()
				sha384.Write(file)
				hash := sha384.Sum(nil)
				k1 := strings.Replace(k, ".", "_", -1)
				c.Set(k1, "sha384-"+base64.StdEncoding.EncodeToString(hash))
				log.Println("Set " + k1 + ": " + "sha384-" + base64.StdEncoding.EncodeToString(hash))
			}
		}
		return next(c)
	}
}
