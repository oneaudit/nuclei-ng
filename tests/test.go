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
	http.HandleFunc("/aspNetErrorPage", aspNetErrorPage)
	http.HandleFunc("/django/", djangoHandler)
	http.HandleFunc("/?token=in_the_url", func(w http.ResponseWriter, _ *http.Request) { forbidden(w) })

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

		if info.IsDir() && !recurse && folderToExpose != path {
			return filepath.SkipDir
		}

		exposureFile := exposurePath + strings.Replace(path, folderToExpose, "", 1)
		if !info.IsDir() {
			http.HandleFunc(exposureFile, func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("X-Content-Type-Options", "nosniff")
				http.ServeFile(w, r, path)
			})
		} else if exposureFile != "/" {
			http.HandleFunc(exposureFile, func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != exposureFile {
					notFound(w)
					return
				}
				forbidden(w)
			})
		}

		return nil
	})
	if err != nil {
		gologger.Fatal().Msgf("Could not add javascript files: %s", err.Error())
	}
}

func notFound(w http.ResponseWriter) {
	// If no other handler matches, serve 404
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	_, _ = w.Write([]byte(`<html>
<head><title>404 Not Found</title></head>
<body><h1>Not Found</h1>
<p>The requested URL was not found on the server. If you entered the URL manually please check your spelling and try again.</p>
</body>
</html>
`))
}

func aspNetErrorPage(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Server", "Microsoft-IIS/10.0")
	w.Header().Set("X-AspNet-Version", "18.0.23273.0")
	w.Header().Set("X-Powered-By", "ASP.NET")
	w.WriteHeader(http.StatusForbidden)
	_, _ = w.Write([]byte(`<!DOCTYPE html>
<html>
<head>
    <title>Configuration Error</title>
    <meta name="viewport" content="width=device-width" />
</head>
<body bgcolor="white">
<span>
    <H1>Server Error in '/' Application.<hr width=100% size=1 color=silver></H1>
    <h2> <i>Configuration Error</i> </h2></span>
    <font face="Arial, Helvetica, Geneva, SunSans-Regular, sans-serif ">
    <b> Description: </b>An error occurred during the processing of a configuration file required to service this request.
        Please review the specific error details below and modify your configuration file appropriately.
    <br><br>

    <b> Parser Error Message: </b>An error occurred loading a configuration file: Failed to start monitoring changes to
        '\\SecretShare\Website\Admin' because access is denied.<br><br>

    <b>Source Error:</b> <br><br>

    <table width=100% bgcolor="#ffffcc">
        <tr>
            <td>
                <code>

                    An application error occurred on the server. The current custom error settings for this application prevent the details of the application error from being viewed remotely (for security reasons). It could, however, be viewed by browsers running on the local server machine.                      </code>

            </td>
        </tr>
    </table>

    <br>

    <b> Source File: </b> \\SecretShare\Website\Admin\web.config<b> &nbsp;&nbsp; Line: </b> 0
    <br><br>

    <br><div class="expandable" onclick="OnToggleTOCLevel1('additionalConfigErrorInfo')">Click here to show additional error information:</div>
    <div id="additionalConfigErrorInfo" style="display: none;">
        <br>            <b> Exception Details: </b>System.Web.HttpException: Failed to start monitoring changes to '\\SecretShare\Website\Admin' because access is denied.<br><br>

        <b>Source Error:</b> <br><br>

        <table width=100% bgcolor="#ffffcc">
            <tr>
                <td>
                    <code>

                        An unhandled exception was generated during the execution of the current web request. Information regarding the origin and location of the exception can be identified using the exception stack trace below.                      </code>

                </td>
            </tr>
        </table>

        <br>

        <b>Stack Trace:</b> <br><br>

        <table width=100% bgcolor="#ffffcc">
            <tr>
                <td>
                    <code><pre>

[HttpException (0x00000000): Failed to start monitoring changes to &#39;\\SecretShare\Website\Admin&#39; because access is denied.]
   System.Web.FileChangesMonitor.FindDirectoryMonitor(String dir, Boolean addIfNotFound, Boolean throwOnError) +555
   [...SNIP...]
</pre>                      </code>

                </td>
            </tr>
        </table>

        <br>



    </div>

    <script type="text/javascript">
        function OnToggleTOCLevel1(level2ID)
        {
            var elemLevel2 = document.getElementById(level2ID);
            if (elemLevel2.style.display == 'none')
            {
                elemLevel2.style.display = '';
            }
            else {
                elemLevel2.style.display = 'none';
            }
        }
    </script>
    <hr width=100% size=1 color=silver>
    <b>Version Information:</b>&nbsp;Microsoft .NET Framework Version:4.7.03056; ASP.NET Version:18.0.23273.0

</font>

</body>
</html>
<!--
[HttpException]: Failed to start monitoring changes to &#39;\\SecretShare\Website\Admin&#39; because access is denied.
   at ...
[ConfigurationErrorsException]: An error occurred loading a configuration file: Failed to start monitoring changes to &#39;\\SecretShare\Website\Admin&#39; because access is denied. (\\SecretShare\Website\Admin\web.config)
   at ...
--><!--
This error page might contain sensitive information because ASP.NET is configured to show verbose error messages using &lt;customErrors mode="Off"/&gt;. Consider using &lt;customErrors mode="On"/&gt; or &lt;customErrors mode="RemoteOnly"/&gt; in production environments.-->
`))
}

func forbidden(w http.ResponseWriter) {
	// If no other handler matches, serve 404
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusForbidden)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		notFound(w)
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
		"/aspNetErrorPage",
		"/django/",
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
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Hello, World! #Nginx"))
}

func apacheHeaderHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Server", "Apache/2.4.41")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Hello, World! #Apache"))
}

func phpHeaderHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Server", "Apache/2.4.41 (Debian) PHP/7.4.0")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Hello, World! #PHP"))
}

func emptyIconsFolder(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
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
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("<!-- my secret password is: toto123 -->"))
}

func multilinesCommentHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("<link rel=\"icon\" href=\"/secret.ico\" type=\"image/x-icon\">" +
		"<!-- \n\n\n\nmy secret password is:\n\n\n\n toto123\n\n-->"))
}

func simpleFormHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
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
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
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
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
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
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
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
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(fmt.Sprintf("Logged in as: %s. <a href='/empty_page/1337'>Logout</a>", username)))
	case "OPTIONS":
		w.Header().Set("Allow", "HEAD, POST, OPTIONS")
		w.WriteHeader(http.StatusOK)

	case "HEAD":
		w.WriteHeader(http.StatusOK)

	default:
		// Handle non-POST methods
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
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
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("X-Entrypoint", "/empty_page/1234/")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`
		<button onclick="location.href='/empty_page/1'">Go to Empty Page 1</button>
	`))
}

func jsEventLinkIdHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
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
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`
		<button id="redirectButton">Go to Empty Page 4</button>
		<script src="/link_id.js"></script>
	`))
}

func secretJettyDirectory(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/secret/" {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`"404 page not found"`))
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
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
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
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

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
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

func djangoHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/django/":
		if strings.HasPrefix(r.Host, "localhost") {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			_, _ = w.Write([]byte(`<!DOCTYPE html>
<html lang="en">
<head>
  <title>DisallowedHost 
       at /</title> 
</head>
<body>
  <pre class="exception_value">Invalid HTTP_HOST header: . You may need to add  to ALLOWED_HOSTS.</pre>
  <table class="meta">
    <tr><th scope="row">Django Version:</th><td>5.0.0</td></tr>
    <tr><th scope="row">Python Version:</th><td>3.09.0</td></tr>
  </table>

<div id="traceback">
  <h2>Traceback <span class="commands"><a href="#" onclick="return switchPastebinFriendly(this);">Switch to copy-and-paste view</a></span></h2>
  <div id="browserTraceback">
    <ul class="traceback">
Django Version: 5.0.0
Python Version: 3.09.0
Installed Applications:
[&#x27;django.contrib.admin&#x27;,
 &#x27;django.contrib.auth&#x27;,
 &#x27;django.contrib.contenttypes&#x27;,
 &#x27;django.contrib.sessions&#x27;,
 &#x27;django.contrib.messages&#x27;,
 &#x27;django.contrib.staticfiles&#x27;,
 &#x27;rest_framework&#x27;,
 &#x27;drf_yasg&#x27;,
 &#x27;myapp&#x27;]

        <tr>
          <td>OS</td>
          <td class="code"><pre>&#x27;Windows_NT&#x27;</pre></td>
        </tr>
      
        <tr>
          <td>ALLOWED_HOSTS</td>
          <td class="code"><pre>[&#x27;django.sec2&#x27;]</pre></td>
        </tr>

        <tr>
          <td>DATABASES</td>
          <td class="code"><pre>{&#x27;default&#x27;: {&#x27;ATOMIC_REQUESTS&#x27;: False,
             &#x27;ENGINE&#x27;: &#x27;django.db.backends.sqlite3&#x27;,
		     </pre></td>
        </tr>
</div>
<p>
You’re seeing this error because you have <code>DEBUG = True</code> in your
Django settings file. Change that to <code>False</code>, and Django will
display a standard page generated by the handler for this status code.
</p>
</body>
</html>`))
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = w.Write([]byte(`<input type="hidden" name="csrfmiddlewaretoken" value="">
<a href="/django/api/">API</a>
<a href="/django/api/swagger/">Swagger</a>
<a href="/django/api/redoc/">Redoc</a>
<a href="/django/install/">Install</a>
<a href="/django/admin/login/">Admin</a>`))
	case "/django/api/":
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = w.Write([]byte(`<!DOCTYPE html>
<html>
<head>
<title>Hello World - Django REST framework</title>
</head>
<body>
<a class='navbar-brand' rel="nofollow" href='https://www.django-rest-framework.org/'>Django REST framework</a>
<!-- SNIP -->
<script type="application/json" id="drf_csrf">
        {
          "csrfHeaderName": "X-CSRFTOKEN",
          "csrfToken": ""
        }
      </script>
      <fake src="/static/rest_framework/js/default.js"></fake>
  </body>
</html>`))
	case "/django/api/redoc/":
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = w.Write([]byte(`<!DOCTYPE html>
<html>
<head>
    <title>My API</title>
	<fake rel="icon" type="image/png" href="/static/drf-yasg/redoc/redoc-logo.png"/>
    <fake rel="stylesheet" type="text/css" href="/static/drf-yasg/style.css"/>
</head>
<body>
<div id="redoc-placeholder"></div>
<script id="redoc-settings" type="application/json">{"lazyRendering": false, "hideHostname": false, "expandResponses": "all", "pathInMiddlePanel": false, "nativeScrollbars": false, "requiredPropsFirst": false, "fetchSchemaWithQuery": true}</script>
<fake src="/static/drf-yasg/insQ.min.js"></fake>
<fake src="/static/drf-yasg/redoc-init.js"></fake>
<fake src="/static/drf-yasg/redoc/redoc.min.js"></fake>
</body>
</html>`))
	case "/django/api/swagger/":
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = w.Write([]byte(`<!DOCTYPE html>
<html>
<head>
    <title>My API</title>
	<fake rel="icon" type="image/png" href="/static/drf-yasg/swagger-ui-dist/favicon-32x32.png"/>
	<fake rel="stylesheet" type="text/css" href="/static/drf-yasg/style.css"/>
	<fake rel="stylesheet" type="text/css" href="/static/drf-yasg/swagger-ui-dist/swagger-ui.css">
</head>
<body class="swagger-body">
<div id="swagger-ui"></div>
<script id="swagger-settings" type="application/json">{"docExpansion": "list", "deepLinking": false, "showExtensions": true, "defaultModelRendering": "model", "defaultModelExpandDepth": 3, "defaultModelsExpandDepth": 3, "showCommonExtensions": true, "supportedSubmitMethods": ["get", "put", "post", "delete", "options", "head", "patch", "trace"], "displayOperationId": true, "persistAuth": false, "refetchWithAuth": false, "refetchOnLogout": false, "fetchSchemaWithQuery": true, "csrfCookie": "csrftoken", "csrfHeader": "X-CSRFTOKEN"}</script>
<fake id="oauth2-config" type="application/json">{}</fake>
<fake src="/static/drf-yasg/swagger-ui-dist/swagger-ui-bundle.js"></fake>
<fake src="/static/drf-yasg/swagger-ui-dist/swagger-ui-standalone-preset.js"></fake>
<fake src="/static/drf-yasg/insQ.min.js"></fake>
<fake src="/static/drf-yasg/immutable.min.js"></fake>
<fake src="/static/drf-yasg/swagger-ui-init.js"></fake>
</body>
</html>`))
	case "/django/admin/login/":
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = w.Write([]byte(`<!DOCTYPE html>
<html>
<head>
<title>Log in | Django site admin</title>
</head>
<body>
<div id="site-name">
<a href="/admin/">Django administration</a>
<form action="/admin/login/?next=/admin/" method="post" id="login-form">
<input type="hidden" name="csrfmiddlewaretoken" value="">
<label for="id_username" class="required">Username:</label>
<input type="text" name="username" autofocus autocapitalize="none" autocomplete="username" maxlength="150" required id="id_username">
<label for="id_password" class="required">Password:</label> <input type="password" name="password" autocomplete="current-password" required id="id_password">
<input type="hidden" name="next" value="/admin/">
<input type="submit" value="Log in">
</form>
</div>
</body>
</html>`))
	default:
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = w.Write([]byte(`<!DOCTYPE html>
<html lang="en">
<head>
  <title>Page not found at </title>
</head>
<body>
<main id="info">
      <p>
      Using the URLconf defined in <code>myproject.urls</code>,
      Django tried these URL patterns, in this order:
      </p>
      <ol>
          <li>
              <code>
                admin/
              </code>
          </li>
          <li>
              <code>
                api/redoc/
                [name='redoc']
              </code>
          </li>
      </ol>
      <p>
          The current path, <code></code>,
        didn't match any of these.
      </p>
  </main>
  <footer id="explanation">
    <p>
      You're seeing this error because you have <code>DEBUG = True</code> in
      your Django settings file. Change that to <code>False</code>, and Django
      will display a standard 404 page.
    </p>
  </footer>
</body>
</html>
`))
	}
}

func libsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
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
