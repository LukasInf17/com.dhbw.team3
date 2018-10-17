package actions

import (
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
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

var contentHashes map[string][]string

// SRIHandler adds support for Subresource integrity
func SRIHandler(next buffalo.Handler) buffalo.Handler {
	contentHashes := map[string][]string{}
	return func(c buffalo.Context) error {
		jsonstring := r.AssetsBox.Bytes("assets/manifest.json")
		var m map[string]string
		json.Unmarshal(jsonstring, &m)
		for k, v := range m {
			if strings.Contains(k, ".css") || strings.Contains(k, ".js") {
				sha384 := sha512.New384()
				sha5 := sha512.New()

				sha384.Write(r.AssetsBox.Bytes("assets/" + v))
				sha5.Write(r.AssetsBox.Bytes("assets/" + v))
				hash := sha384.Sum(nil)
				hash2 := sha5.Sum(nil)
				if strings.Contains(k, ".css") {
					contentHashes["style"] = append(contentHashes["style"], "sha512-"+base64.StdEncoding.EncodeToString(hash2))
				} else {
					contentHashes["script"] = append(contentHashes["script"], "sha512-"+base64.StdEncoding.EncodeToString(hash2))
				}
				k1 := strings.Replace(k, ".", "_", -1)
				c.Set(k1, "sha384-"+base64.StdEncoding.EncodeToString(hash)+" sha512-"+base64.StdEncoding.EncodeToString(hash2))
			}
		}
		return next(c)
	}
}

// SetSecurityHeaders sets security headers
func SetSecurityHeaders(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		scriptstring := ""
		stylestring := ""
		for k, v := range contentHashes {
			for _, vl := range v {
				switch k {
				case "style":
					stylestring = stylestring + "'" + vl + "' "
				case "script":
					scriptstring = scriptstring + "'" + vl + "' "
				}
			}
		}
		c.Response().Header().Add("Content-Security-Policy", "default-src 'none'; script-src 'strict-dynamic' "+scriptstring+"'self'; img-src 'self'; style-src 'self' "+stylestring+"; form-action 'self'; frame-ancestors 'none'; object-src 'none'; base-uri 'none';")
		return next(c)
	}
}
