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
	http.HandleFunc("/icons/", sitemapFakeFolderHandler)
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

	http.ListenAndServe(":5000", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Render the list of routes, excluding secret ones
	routes := []string{
		"/robots.txt",
		"/humans.txt",
		"/sitemap.xml",
		"/.well-known/security.txt",
		"/nginx-v",
		"/apache-v",
		"/php-v",
		"/icons/",
		"/icons/sitemap.png",
		"/icons/robots.txt",
		"/ng_hidden_spy",
		"/ng_hidden_login",
		"/empty_page/1337",
		"/re?redirect=/",
	}
	tmpl, err := template.New("index").Parse(`
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>Index</title>
</head>
<body>
	<h1>Routes</h1>
	<ul>{{range .}}
		<li><a href="{{.}}">{{.}}</a></li>{{end}}
	</ul>
</body>
</html>
`)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Server", "Werkzeug/1.5.7 Python/3.10.2")
	tmpl.Execute(w, routes)
}

func robotsTxtHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("User-agent: *\nDisallow:\n"))
}

func humansTxtHandler(w http.ResponseWriter, r *http.Request) {
	content := `
/* TEAM */
	Chef: John Doe
	Contact: jdoe [at] example.com
	From: City, State, Country
`
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(content))
}

func sitemapHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/xml")
	w.Write([]byte(`<?xml version="1.0" encoding="UTF-8"?>
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
	w.Write([]byte(content))
}

func nginxHeaderHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Server", "nginx/1.33.7")
	w.Write([]byte("Hello, World! #Nginx"))
}

func apacheHeaderHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Server", "Apache/2.4.41")
	w.Write([]byte("Hello, World! #Apache"))
}

func phpHeaderHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-Powered-By", "PHP/7.4.0")
	w.Header().Set("Server", "Apache/2.4.41 (Debian) PHP/7.4.0")
	http.SetCookie(w, &http.Cookie{Name: "PHPSESSID", Value: "session_id_here"})
	w.Write([]byte("Hello, World! #PHP"))
}

func sitemapFakeFolderHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`<a href="robots.txt">old robots.txt file</a>`))
}

func sitemapFakeFileHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte{})
}

func nonGenericRobotsTxtHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("User-agent: *\nDisallow: /admin/"))
}

func inlineCommentHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<!-- my secret password is: toto123 -->"))
}

func multilinesCommentHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<!-- \n\n\n\nmy secret password is:\n\n\n\n toto123\n\n-->"))
}

func simpleFormHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`
	<form action="#" method="GET">
		<label for="name">Name:</label>
		<input type="text" id="name" name="name" required>
		<input type="text" id="csrf" name="csrf" hidden value="nocsrf_token">
		<button type="submit">Submit</button>
	</form>
	`))
}

func complexJsFormHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/search.html")
}

func ngHiddenSpyHandler(w http.ResponseWriter, r *http.Request) {
	var input string
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	input = r.FormValue("input")
	w.Write([]byte(fmt.Sprintf("You typed: %s", input)))
}

func cookieFormHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`
	<form action="/ng_hidden_login" method="POST">
		<label for="username">Username:</label>
		<input type="text" id="username" name="username" required>
		<button type="submit">Login</button>
	</form>
	`))
}

func ngHiddenLoginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	encodedUsername := base64.StdEncoding.EncodeToString([]byte(username))
	http.SetCookie(w, &http.Cookie{Name: "user", Value: encodedUsername})
	w.Write([]byte(fmt.Sprintf("Logged in as: %s. <a href='/empty_page/1337'>Logout</a>", username)))
}

func jsEventLinkHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`
		<button onclick="location.href='/empty_page/1'">Go to Empty Page 1</button>
	`))
}

func jsEventLinkIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`
		<button id="redirectButton">Go to Empty Page 2</button>
		<script>
			document.getElementById('redirectButton').addEventListener('click', function() {
				location.href = '/empty_page/2';
			});
		</script>
	`))
}

func jsExternalLinkIdHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`
		<button id="redirectButton">Go to Empty Page 4</button>
		<script src="/static/link_id.js"></script>
	`))
}

func emptyPageHandler(w http.ResponseWriter, r *http.Request) {
	counterID := r.URL.Path[len("/empty_page/"):]
	w.Write([]byte(fmt.Sprintf("Empty page: %s", counterID)))
}

func redirectToHandler(w http.ResponseWriter, r *http.Request) {
	targetURL := r.URL.Query().Get("redirect")
	http.Redirect(w, r, targetURL, http.StatusFound)
}
