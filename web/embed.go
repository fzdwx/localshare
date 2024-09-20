package web

import (
	"embed"
	_ "embed"
)

//go:embed index.html
//go:embed output.css
//go:embed client.js
//go:embed util.js
var Dist embed.FS
