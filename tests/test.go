package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/projectdiscovery/gologger"
	"net/http"
	"os"
	"path/filepath"
	"strings"
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
	http.HandleFunc("/secret/", secretJettyDirectory)
	http.HandleFunc("/re", redirectToHandler)
	http.HandleFunc("/cors", corsHandler)
	http.HandleFunc("/libs", libsHandler)

	exposeFolder("tests/static/js/", "/assets/js/", true)
	exposeFolder("tests/static/composer/", "/", true)
	exposeFolder("tests/static/", "/", false)
	exposeFolder("tests/static/git/", "/.git/", true)

	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		gologger.Fatal().Msgf("Could not start test server: %s", err.Error())
	}
}

func exposeFolder(folderToExpose string, exposurePath string, recurse bool) {
	err := filepath.Walk(folderToExpose, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			if !recurse && folderToExpose != path {
				return filepath.SkipDir
			}
			return nil
		}

		exposureFile := exposurePath + strings.Replace(path, folderToExpose, "", 1)
		http.HandleFunc(exposureFile, func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("X-Content-Type-Options", "nosniff")
			http.ServeFile(w, r, path)
		})

		return nil
	})
	if err != nil {
		gologger.Fatal().Msgf("Could not add javascript files: %s", err.Error())
	}
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
		return
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
		"/secret/",
		"/cors",
		"/libs",
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

func robotsTxtHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	_, _ = w.Write([]byte("User-agent: *\nDisallow:\n"))
}

func humansTxtHandler(w http.ResponseWriter, _ *http.Request) {
	content := `
/* TEAM */
	Chef: John Doe
	Contact: jdoe [at] example.com
	From: City, State, Country
`
	w.Header().Set("Content-Type", "text/plain")
	_, _ = w.Write([]byte(content))
}

func sitemapHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/xml")
	_, _ = w.Write([]byte(`<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap-image/1.1">
    <url>
        <loc>/icons/sitemap.png</loc>
    </url>
</urlset>`))
}

func securityTxtHandler(w http.ResponseWriter, _ *http.Request) {
	content := `# Welcome to Example Security Page
Contact: security@example.com
Contact: security[at]example.com`
	w.Header().Set("Content-Type", "text/plain")
	_, _ = w.Write([]byte(content))
}

func nginxHeaderHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Server", "nginx/1.33.7")
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Hello, World! #Nginx"))
}

func apacheHeaderHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Server", "Apache/2.4.41")
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Hello, World! #Apache"))
}

func phpHeaderHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Server", "Apache/2.4.41 (Debian) PHP/7.4.0")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Hello, World! #PHP"))
}

func emptyIconsFolder(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("server", "Jetty(12.0.17.v20201231)")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`<a href="robots.txt">old robots.txt file</a>`))
}

func sitemapFakeFileHandler(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte{})
}

func nonGenericRobotsTxtHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("X-Powered-By", "Unknown/21.4.21")
	_, _ = w.Write([]byte("User-agent: *\nDisallow: /admin/"))
}

func inlineCommentHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("X-Powered-By", "PHP/7.4.0")
	http.SetCookie(w, &http.Cookie{Name: "PHPSESSID", Value: "session_id_here"})
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("<!-- my secret password is: toto123 -->"))
}

func multilinesCommentHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("<link rel=\"icon\" href=\"/secret.ico\" type=\"image/x-icon\">" +
		"<!-- \n\n\n\nmy secret password is:\n\n\n\n toto123\n\n-->"))
}

func simpleFormHandler(w http.ResponseWriter, _ *http.Request) {
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
	http.ServeFile(w, r, "tests/static/search.html")
}

type HiddenSpyInput struct {
	Input string `json:"input"`
}

func ngHiddenSpyHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var inputData HiddenSpyInput
		err := json.NewDecoder(r.Body).Decode(&inputData)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		_, _ = w.Write([]byte(fmt.Sprintf("You typed: %s", inputData.Input)))

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

func cookieFormHandler(w http.ResponseWriter, _ *http.Request) {
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
		http.SetCookie(w, &http.Cookie{Name: "user", Value: encodedUsername})
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
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

func jsEventLinkHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("X-Entrypoint", "/empty_page/1234/")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`
		<button onclick="location.href='/empty_page/1'">Go to Empty Page 1</button>
	`))
}

func jsEventLinkIdHandler(w http.ResponseWriter, _ *http.Request) {
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

func jsExternalLinkIdHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`
		<button id="redirectButton">Go to Empty Page 4</button>
		<script src="/link_id.js"></script>
	`))
}

func secretJettyDirectory(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/secret/" {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`"404 page not found"`))
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<link href="jetty-dir.css" rel="stylesheet" />
<title>Directory: /secret/</title>
</head>
<body>
<h1 class="title">Directory: /secret/</h1>
<table class="listing">
<thead><tr><th class="name"><a href="?C=N&O=D">Name&nbsp; &#8679;</a></th><th class="lastmodified"><a href="?C=M&O=A">Last Modified</a></th><th class="size"><a href="?C=S&O=A">Size</a></th></tr></thead>
<tbody>
<tr><td class="name"><a href="/secret/dir.empty">dir.empty&nbsp;</a></td><td class="lastmodified">1 janv. 2009 00:00:00&nbsp;</td><td class="size">1337 bytes&nbsp;</td></tr>
</tbody>
</table>
</body></html>
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

func corsHandler(w http.ResponseWriter, r *http.Request) {
	origin := r.Header.Get("Origin")
	if origin == "" {
		origin = fmt.Sprintf("http://%s", r.Host)
	}

	w.Header().Set("Access-Control-Allow-Origin", origin)
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	if r.Method == http.MethodOptions {
		return
	}

	w.Header().Set("Content-Type", "text/html")
	_, _ = fmt.Fprintf(w, `
		<!DOCTYPE html>
		<html>
		<head>
			<title>CORS Example</title>
		</head>
		<body>
			<p>Your Origin: %s</p>
			<p>Your Cookies: %s</p>
		</body>
		</html>
	`, origin, r.Header.Get("Cookie"))
}

func libsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Content-Security-Policy", "default-src 'none'")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`<!DOCTYPE html>
<html>
<head>
	<title>Random Libs</title>

	<!-- version in path -->
    <script src="https://cdnjs.cloudflare.com/ajax/libs/moment.js/2.29.1/moment.min.js"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/aos/2.3.4/aos.js"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/d3/7.8.1/d3.min.js"></script>
	<script src="https://cdnjs.cloudflare.com/ajax/libs/lodash.js/4.17.21/lodash.min.js"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/2.11.6/umd/popper.min.js"></script>
    <script src="https://code.jquery.com/ui/1.12.1/jquery-ui.min.js"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/jquery-validate/1.19.3/jquery.validate.min.js"></script>
	<script src="https://cdnjs.cloudflare.com/ajax/libs/gsap/3.11.1/gsap.min.js"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/leaflet/1.7.1/leaflet.min.js"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/socket.io/4.3.2/socket.io.min.js"></script>

	<!-- version in name -->
    <script src="https://code.jquery.com/jquery-3.6.0.min.js"></script>

	<!-- version in copyright -->
    <script src="https://cdn.jsdelivr.net/npm/chart.js"></script>

	<!-- @version (+redirect) -->
	<script src="https://unpkg.com/react@18/umd/react.production.min.js"></script>
    <script src="https://unpkg.com/react-dom@18/umd/react-dom.production.min.js"></script>
    <script src="https://cdn.jsdelivr.net/npm/vue@2.6.14/dist/vue.min.js"></script>
    <script src="https://cdn.jsdelivr.net/npm/alpinejs@2.8.2/dist/alpine.min.js"></script>
    <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0-alpha1/dist/js/bootstrap.min.js"></script>

	<!-- no version -->
    <script src="https://cdn.jsdelivr.net/npm/axios/dist/axios.min.js"></script>

	<!-- local -->
	<script src="/assets/js/angular.js"></script>
	<script src="/assets/js/bootstrap.js"></script>
	<script src="/assets/js/jquery.js"></script>
	<script src="/assets/js/jquery-1.14.0.js"></script>
	<script src="/assets/js/jquery-migrate.js"></script>
	<script src="/assets/js/jquery-ui.js"></script>
	<script src="/assets/js/jszip.js"></script>
	<script src="/assets/js/leaflet.js"></script>
	<script src="/assets/js/bundle.js"></script>	
</head>
</html>`))
}
