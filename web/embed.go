package web

import (
	"embed"
	_ "embed"
)

//go:embed index.html
//go:embed output.css
var Dist embed.FS
