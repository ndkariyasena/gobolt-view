package web

import (
	"embed"
)

/* Note: This assumes the final Svelte build is moved or copied to server/web/dist/. */
/* //go:embed dist/* */
var StaticFiles embed.FS
