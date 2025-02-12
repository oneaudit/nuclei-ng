package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/robots.txt", robotsTxtHandler)
	http.HandleFunc("/humans.txt", humansTxtHandler)
	http.HandleFunc("/sitemap.xml", sitemapHandler)
	http.HandleFunc("/.well-known/security.txt", securityTxtHandler)
	http.HandleFunc("/nginx-v", nginxHeaderHandler)
	http.HandleFunc("/apache-v", apacheHeaderHandler)
	http.HandleFunc("/php-v", phpHeaderHandler)
	http.HandleFunc("/icons/", emptyIconsFolder)
	http.HandleFunc("/icons/sitemap.png", sitemapFakeFileHandler)
	http.HandleFunc("/icons/robots.txt", nonGenericRobotsTxtHandler)
	http.HandleFunc("/comment", inlineCommentHandler)
	http.HandleFunc("/comment-long", multilinesCommentHandler)
	http.HandleFunc("/simple-form", simpleFormHandler)
	http.HandleFunc("/complex-js-form", complexJsFormHandler)
	http.HandleFunc("/ng_hidden_spy", ngHiddenSpyHandler)
	http.HandleFunc("/cookie-form", cookieFormHandler)
	http.HandleFunc("/ng_hidden_login", ngHiddenLoginHandler)
	http.HandleFunc("/js-event-link", jsEventLinkHandler)
	http.HandleFunc("/js-event-link-id", jsEventLinkIdHandler)
	http.HandleFunc("/js-external-link-id", jsExternalLinkIdHandler)
	http.HandleFunc("/empty_page/", emptyPageHandler)
	http.HandleFunc("/re", redirectToHandler)

	_ = http.ListenAndServe(":5000", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		// If no other handler matches, serve 404
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`<html>
<head><title>404 Not Found</title></head>
<body><h1>Not Found</h1>
<p>The requested URL was not found on the server. If you entered the URL manually please check your spelling and try again.</p>
</body>
</html>
`))
	}

	// Render the list of routes, excluding secret ones
	routes := []string{
		"/nginx-v",
		"/apache-v",
		"/php-v",
		"/comment",
		"/comment-long",
		"/simple-form",
		"/complex-js-form",
		"/cookie-form",
		"/js-event-link",
		"/js-event-link-id",
		"/js-external-link-id",
		"/re?redirect=/",
	}
	tmpl, err := template.New("index").Parse(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Index of /</title>
</head>
<body>
    <h1>Index of /</h1>
    <table>
        <tbody>
            <tr><th valign="top"><img src="" alt="[ICO]"></th><th><a href="?C=N;O=D">Name</a></th><th><a href="?C=M;O=A">Last modified</a></th><th><a href="?C=S;O=A">Size</a></th><th><a href="?C=D;O=A">Description</a></th></tr>
            <tr><th colspan="5"><hr></th></tr>
            <tr><td valign="top"><img src="" alt="[PARENTDIR]"></td><td><a href="/">Parent Directory</a>       </td><td>&nbsp;</td><td align="right">  - </td><td>&nbsp;</td></tr>
            {{range .}}
            <tr>
                <td valign="top">
                    <img src="" alt="[IMG]">
                </td>
                <td>
                    <a href="{{.}}">{{.}}</a>
                </td>
                <td align="right">2009-01-01 00:00  </td>
                <td align="right">1337 </td>
                <td>&nbsp;</td>
            </tr>
            {{end}}
            <tr><th colspan="5"><hr></th></tr>
        </tbody>
    </table>
</body>
</html>
`)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Server", "Werkzeug/1.5.7 Python/3.10.2")
	_ = tmpl.Execute(w, routes)
}

func robotsTxtHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	_, _ = w.Write([]byte("User-agent: *\nDisallow:\n"))
}

func humansTxtHandler(w http.ResponseWriter, r *http.Request) {
	content := `
/* TEAM */
	Chef: John Doe
	Contact: jdoe [at] example.com
	From: City, State, Country
`
	w.Header().Set("Content-Type", "text/plain")
	_, _ = w.Write([]byte(content))
}

func sitemapHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/xml")
	_, _ = w.Write([]byte(`<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap-image/1.1">
    <url>
        <loc>/icons/sitemap.png</loc>
    </url>
</urlset>`))
}

func securityTxtHandler(w http.ResponseWriter, r *http.Request) {
	content := `# Welcome to Example Security Page
Contact: security@example.com
Contact: security[at]example.com`
	w.Header().Set("Content-Type", "text/plain")
	_, _ = w.Write([]byte(content))
}

func nginxHeaderHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Server", "nginx/1.33.7")
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Hello, World! #Nginx"))
}

func apacheHeaderHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Server", "Apache/2.4.41")
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Hello, World! #Apache"))
}

func phpHeaderHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("X-Powered-By", "PHP/7.4.0")
	w.Header().Set("Server", "Apache/2.4.41 (Debian) PHP/7.4.0")
	http.SetCookie(w, &http.Cookie{Name: "PHPSESSID", Value: "session_id_here"})
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Hello, World! #PHP"))
}

func emptyIconsFolder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`<a href="robots.txt">old robots.txt file</a>`))
}

func sitemapFakeFileHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte{})
}

func nonGenericRobotsTxtHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	_, _ = w.Write([]byte("User-agent: *\nDisallow: /admin/"))
}

func inlineCommentHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("<!-- my secret password is: toto123 -->"))
}

func multilinesCommentHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("<!-- \n\n\n\nmy secret password is:\n\n\n\n toto123\n\n-->"))
}

func simpleFormHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`
	<form action="#" method="GET">
		<label for="name">Name:</label>
		<input type="text" id="name" name="name" required>
		<input type="text" id="csrf" name="csrf" hidden value="nocsrf_token">
		<button type="submit">Submit</button>
	</form>
	`))
}

func complexJsFormHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/search.html")
}

func ngHiddenSpyHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var input string
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		input = r.FormValue("input")
		_, _ = w.Write([]byte(fmt.Sprintf("You typed: %s", input)))

	case "OPTIONS":
		w.Header().Set("Allow", "HEAD, POST, OPTIONS")
		w.WriteHeader(http.StatusOK)

	case "HEAD":
		w.WriteHeader(http.StatusOK)

	default:
		// Handle non-POST methods
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`<html>
<head><title>404 Not Found</title></head>
<body><h1>Not Found</h1>
<p>The requested URL was not found on the server. If you entered the URL manually please check your spelling and try again.</p>
</body>
</html>
`))
	}
}

func cookieFormHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`
	<form action="/ng_hidden_login" method="POST">
		<label for="username">Username:</label>
		<input type="text" id="username" name="username" required>
		<button type="submit">Login</button>
	</form>
	`))
}

func ngHiddenLoginHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		username := r.FormValue("username")

		if username == "" {
			w.Header().Set("Content-Type", "text/html")
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(`<!doctype html>
<html lang="en">
<title>400 Bad Request</title>
<h1>Bad Request</h1>
<p>The browser (or proxy) sent a request that this server could not understand.</p>
</html>
`))
			return
		}

		encodedUsername := base64.StdEncoding.EncodeToString([]byte(username))
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		http.SetCookie(w, &http.Cookie{Name: "user", Value: encodedUsername})
		_, _ = w.Write([]byte(fmt.Sprintf("Logged in as: %s. <a href='/empty_page/1337'>Logout</a>", username)))

	case "OPTIONS":
		w.Header().Set("Allow", "HEAD, POST, OPTIONS")
		w.WriteHeader(http.StatusOK)

	case "HEAD":
		w.WriteHeader(http.StatusOK)

	default:
		// Handle non-POST methods
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`<html>
<head><title>404 Not Found</title></head>
<body><h1>Not Found</h1>
<p>The requested URL was not found on the server. If you entered the URL manually please check your spelling and try again.</p>
</body>
</html>
`))
	}
}

func jsEventLinkHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`
		<button onclick="location.href='/empty_page/1'">Go to Empty Page 1</button>
	`))
}

func jsEventLinkIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`
		<button id="redirectButton">Go to Empty Page 2</button>
		<script>
			document.getElementById('redirectButton').addEventListener('click', function() {
				location.href = '/empty_page/2';
			});
		</script>
	`))
}

func jsExternalLinkIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`
		<button id="redirectButton">Go to Empty Page 4</button>
		<script src="/static/link_id.js"></script>
	`))
}

func emptyPageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	counterID := r.URL.Path[len("/empty_page/"):]
	_, _ = w.Write([]byte(fmt.Sprintf("Empty page: %s", counterID)))
}

func redirectToHandler(w http.ResponseWriter, r *http.Request) {
	targetURL := r.URL.Query().Get("redirect")
	http.Redirect(w, r, targetURL, http.StatusFound)
}
