package extensions

import (
	"strings"
)

var DefaultDenylist = []string{".3g2", ".3gp", ".7z", ".apk", ".arj", ".avi", ".axd", ".bmp", ".csv", ".deb", ".dll", ".doc", ".drv", ".eot", ".exe", ".flv", ".gif", ".gifv", ".gz", ".h264", ".ico", ".iso", ".jar", ".jpeg", ".jpg", ".lock", ".m4a", ".m4v", ".map", ".mkv", ".mov", ".mp3", ".mp4", ".mpeg", ".mpg", ".msi", ".ogg", ".ogm", ".ogv", ".otf", ".pdf", ".pkg", ".png", ".ppt", ".psd", ".rar", ".rm", ".rpm", ".svg", ".swf", ".sys", ".tar.gz", ".tar", ".tif", ".tiff", ".ttf", ".txt", ".vob", ".wav", ".webm", ".webp", ".wmv", ".woff", ".woff2", ".xcf", ".xls", ".xlsx", ".zip"}

func IsHTMLFile(extension string) bool {
	// We are using the DefaultDenylist of Project Discovery
	// To exclude uninteresting extensions
	// This reduces the number of extensions marked as "html"
	// (we are using a blacklist to avoid missing unhandled pages)
	if extension == "" {
		return true
	}

	for _, htmlExt := range DefaultDenylist {
		if strings.EqualFold(extension, htmlExt) {
			return true
		}
	}
	return false
}
