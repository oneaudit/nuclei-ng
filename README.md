<h1 align="center">
  Nuclei-ng
</h1>
<h4 align="center">Nuclei-ng is a wrapper of <a href="https://github.com/projectdiscovery/nuclei">nuclei</a> from Project Discovery.</h4>

## Installation ✍️

nuclei-ng requires **Go 1.22.5+** to install successfully.

```console
CGO_ENABLED=1 go install github.com/oneaudit/nuclei-ng/cmd/nuclei-ng@latest
```

## Features ⭐

We have not modified, and do not intend to modify, the Nuclei source code.

* 😎 Summarized output per template
* 🗺️ Tagging per URL extension
* 🐸️ Fetching missing information using `oneaudit:` from templates
* 🚀 [TODO] Export results in formats similar to Nuclei

We’ve added JavaScript additional utilities within the tool:

* 🔑 [TODO] JavaScript secret detection using [jsluice](https://github.com/BishopFox/jsluice)
* 🔍 JavaScript URL-based version detection for out-of-scope URLs
* 🐲 [TODO] JavaScript tampering detection

We’ve added an internal proxy server, which enables us to:

* 🤠 [TODO] Replay responses from a CSV file
* 🤖 Cache responses marked with `(proxy)`
* 🧪 Do not send requests marked with `(request)`

This wrapper is designed to be used with DAST templates. A free public example can be found [here](https://github.com/oneaudit/nuclei-dast-templates).

## Usage 📚

```
nuclei-ng -config config.yaml -i oppa_openapi.yaml
```

This will display the following results on the test instance.

```
                     __     _
   ____  __  _______/ /__  (_)
  / __ \/ / / / ___/ / _ \/ /
 / / / / /_/ / /__/ /  __/ /
/_/ /_/\__,_/\___/_/\___/_/ ng  v1.3.0

		github.com/oneaudit

[INF] Running nuclei with tags: [html] against 46 targets
[INF] Temporary file created: /tmp/swagger.yaml1269739532
[cdn-detect:cdnjs] [http] [info] Found on 1 URLs [/libs]
[cdn-detect:jquery-cdn] [http] [info] Found on 1 URLs [/libs]
[cdn-detect:jsdelivr] [http] [info] Found on 1 URLs [/libs]
[cdn-detect:unpkg] [http] [info] Found on 1 URLs [/libs]
[directory-listing] [http] [info] Found on 4 URLs [/, /secret/]
[form-detect] [http] [info] Found on 1 URLs ["#"] [/simple-form]
[form-detect] [http] [info] Found on 1 URLs ["/ng_hidden_login"] [/cookie-form]
[html-comments-detect] [javascript] [info] Found on 1 URLs ["<!-- @version (+redirect) -->"] [/libs]
[html-comments-detect] [javascript] [info] Found on 1 URLs ["<!-- \\n\\n\\n\\nmy secret password is:\\n\\n\\n\\n toto123\\n\\n-->"] [/comment-long]
[html-comments-detect] [javascript] [info] Found on 1 URLs ["<!-- local -->"] [/libs]
[html-comments-detect] [javascript] [info] Found on 1 URLs ["<!-- my secret password is: toto123 -->"] [/comment]
[html-comments-detect] [javascript] [info] Found on 1 URLs ["<!-- no version -->"] [/libs]
[html-comments-detect] [javascript] [info] Found on 1 URLs ["<!-- version in copyright -->"] [/libs]
[html-comments-detect] [javascript] [info] Found on 1 URLs ["<!-- version in name -->"] [/libs]
[html-comments-detect] [javascript] [info] Found on 1 URLs ["<!-- version in path -->"] [/libs]
[html-comments-detect] [javascript] [info] Found on 1 URLs ["<!--\\nThis error page might contain sensitive information because ASP.NET is configured to show verbose error messages using &lt;customErrors mode="Off"/&gt;. Consider using &lt;customErrors mode="On"/&gt; or &lt;customErrors mode="RemoteOnly"/&gt; in production environments.-->"] [/aspNetErrorPage]
[html-comments-detect] [javascript] [info] Found on 1 URLs ["<!--\\n[HttpException]: Failed to start monitoring changes to &#39;\\\\SecretShare\\Website\\Admin&#39; because access is denied.\\n   at ...\\n[ConfigurationErrorsException]: An error occurred loading a configuration file: Failed to start monitoring changes to &#39;\\\\SecretShare\\Website\\Admin&#39; because access is denied. (\\\\SecretShare\\Website\\Admin\\web.config)\\n   at ...\\n-->"] [/aspNetErrorPage]

[INF] Running nuclei with tags: [generic] against 64 targets
[INF] Temporary file created: /tmp/swagger.yaml747212024
[composer:composer.json] [http] [low] Found on 2 URLs [/composer.json, /vendor/composer/installed.json]
[composer:composer.phar] [http] [medium] Found on 1 URLs [/composer.phar]
[composer:vendor] [http] [medium] Found on 2 URLs [/vendor/autoload.php, /vendor/composer/installed.json]
[cookie-detect] [http] [info] Found on 1 URLs ["PHPSESSID"] [/comment]
[cookie-detect] [http] [info] Found on 1 URLs ["user"] [/ng_hidden_login]
[cookies-without-httponly] [javascript] [info] Found on 1 URLs ["PHPSESSID"] [/comment]
[cookies-without-httponly] [javascript] [info] Found on 1 URLs ["user"] [/ng_hidden_login]
[cookies-without-secure] [javascript] [info] Found on 1 URLs ["PHPSESSID"] [/comment]
[cookies-without-secure] [javascript] [info] Found on 1 URLs ["user"] [/ng_hidden_login]
[cors-detect:checked] [http] [info] Found on 1 URLs [header:User-Agent] [GET] [/cors]
[cors-detect:exploited] [http] [info] Found on 1 URLs [header:Origin] [GET] [/cors]
[email-detect] [http] [info] Found on 2 URLs ["barbushin@gmail.com"] [/composer.lock, /vendor/composer/installed.json]
[email-detect] [http] [info] Found on 1 URLs ["jdoe [at] example.com"] [/humans.txt]
[email-detect] [http] [info] Found on 2 URLs ["jdoe@example.com"] [/.git/config, /.git/logs/HEAD]
[email-detect] [http] [info] Found on 1 URLs ["security@example.com"] [/.well-known/security.txt]
[email-detect] [http] [info] Found on 1 URLs ["security[at]example.com"] [/.well-known/security.txt]
[email-detect] [http] [info] Found on 1 URLs ["stuart [at] stuartk.com"] [/assets/js/jszip.js]
[favicon-detect:Vue] [http] [info] Found on 1 URLs [/favicon.ico]
[favicon-new:md5] [http] [info] Found on 1 URLs ["b7f5b488d0b802ed63ea4ffefbbb1d6d"] [/secret.ico]
[favicon-new:mmh3] [http] [info] Found on 1 URLs ["1823185746"] [/secret.ico]
[file-inclusion:detect] [javascript] [low] Found on 1 URLs ["redirect"] [/re]
[git-config:config] [http] [medium] Found on 3 URLs [/.git/HEAD, /.git/config, /.git/index]
[git-config:folder] [http] [medium] Found on 1 URLs [/.git/]
[git-config:ignore] [http] [medium] Found on 1 URLs [/.gitignore]
[git-config:logs] [http] [medium] Found on 1 URLs [/.git/logs/HEAD]
[http-missing-security-headers:content-security-policy] [http] [info] Found on 63 URLs [/, /.git/, /.git/HEAD, ...]
[http-missing-security-headers:permissions-policy] [http] [info] Found on 64 URLs [/, /.git/, /.git/HEAD, ...]
[http-missing-security-headers:referrer-policy] [http] [info] Found on 64 URLs [/, /.git/, /.git/HEAD, ...]
[http-missing-security-headers:strict-transport-security] [http] [info] Found on 64 URLs [/, /.git/, /.git/HEAD, ...]
[http-missing-security-headers:x-content-type-options] [http] [info] Found on 42 URLs [/, /.git/, /.git/logs/, ...]
[http-suspicious-request-headers] [javascript] [info] Found on 1 URLs ["X-Api-Key: MYS3cr374P1K3y"] [/ng_hidden_spy]
[http-suspicious-response-headers] [javascript] [info] Found on 1 URLs ["X-Entrypoint: /empty_page/1234/"] [/js-event-link]
[known-403-page:apache] [http] [info] Found on 3 URLs [/.git/, /assets/js/, /vendor]
[known-403-page:aspnet] [http] [info] Found on 1 URLs [/aspNetErrorPage]
[known-404-page:custom] [http] [info] Found on 7 URLs [/.git/logs/, /.well-known/, /PAGENOTFOUND, ...]
[open-redirect:detect] [javascript] [low] Found on 1 URLs ["redirect"] [/re]
[open-redirect:exploit] [http] [low] Found on 1 URLs [query:redirect] [GET] [/]
[options-method-generic] [http] [info] Found on 2 URLs [/ng_hidden_login, /ng_hidden_spy]
[options-method-non-generic] [javascript] [info] Found on 2 URLs ["POST"] [/ng_hidden_login]
[robots-txt-generic] [http] [info] Found on 2 URLs [/robots.txt]
[robots-txt-non-generic] [http] [info] Found on 1 URLs ["b77a70c632c0be0cea6bbce3dff8a2317392303b0730327589b3485c4b892dcd"] [/icons/robots.txt]
[server-errors:aspnet] [http] [info] Found on 1 URLs [/aspNetErrorPage]
[sitemap-detect] [http] [info] Found on 2 URLs [/sitemap.xml]
[stacktrace:aspnet] [http] [low] Found on 1 URLs [/aspNetErrorPage]
[stacktrace:generic] [http] [low] Found on 1 URLs [/aspNetErrorPage]
[tech-detect:apache] [http] [info] Found on 2 URLs [/apache-v, /php-v]
[tech-detect:aspnet] [http] [info] Found on 1 URLs [/aspNetErrorPage]
[tech-detect:jetty] [http] [info] Found on 5 URLs [/icons/, /icons/.robots.txt.swp, /icons/robots.bak, ...]
[tech-detect:jsdelivr] [http] [info] Found on 1 URLs [/libs]
[tech-detect:ms-iis] [http] [info] Found on 1 URLs [/aspNetErrorPage]
[tech-detect:nginx] [http] [info] Found on 1 URLs [/nginx-v]
[tech-detect:php] [http] [info] Found on 1 URLs [/comment]
[tech-detect:python] [http] [info] Found on 2 URLs [/]
[tech-version-new] [javascript] [info] Found on 1 URLs ["Unknown/21.4.21"] [/icons/robots.txt]
[tech-version-new] [javascript] [info] Found on 1 URLs ["Version:18.0.23273.0"] [/aspNetErrorPage]
[tech-version-new] [javascript] [info] Found on 1 URLs ["Version:4.7.03056"] [/aspNetErrorPage]
[tech-version-new] [javascript] [info] Found on 1 URLs ["empty_page/1234"] [/js-event-link]
[tech-version:apache] [http] [info] Found on 2 URLs ["2.4.41"] [/apache-v, /php-v]
[tech-version:aspnet] [http] [info] Found on 1 URLs ["18.0.23273.0"] [/aspNetErrorPage]
[tech-version:aspnet_framework] [http] [info] Found on 1 URLs ["4.7.03056"] [/aspNetErrorPage]
[tech-version:jetty] [http] [info] Found on 5 URLs ["12.0.17.v20201231"] [/icons/, /icons/.robots.txt.swp, /icons/robots.bak, ...]
[tech-version:msiis] [http] [info] Found on 1 URLs ["10.0"] [/aspNetErrorPage]
[tech-version:nginx] [http] [info] Found on 1 URLs ["1.33.7"] [/nginx-v]
[tech-version:os] [http] [info] Found on 1 URLs ["Debian"] [/php-v]
[tech-version:payara-server] [http] [info] Found on 1 URLs ["2.2023.2"] [/comment-long]
[tech-version:php] [http] [info] Found on 2 URLs ["7.4.0"] [/comment, /php-v]
[tech-version:python] [http] [info] Found on 2 URLs ["3.10.2"] [/]
[tech-version:werkzeug] [http] [info] Found on 2 URLs ["1.5.7"] [/]
[well-known-detect] [http] [info] Found on 2 URLs [/.well-known/security.txt, /humans.txt]

[INF] Running nuclei with tags: [javascript] against 10 targets
[INF] Temporary file created: /tmp/swagger.yaml3600235712
[javascript-library] [javascript] [info] Found on 1 URLs ["angular.js==1.8.3"] [/assets/js/angular.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["bootstrap.js==5.3.3"] [/assets/js/bootstrap.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["bundle.js==unknown"] [/assets/js/bundle.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["jquery-1.14.0.js==1.14.0"] [/assets/js/jquery-1.14.0.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["jquery-migrate.js==3.5.2"] [/assets/js/jquery-migrate.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["jquery-ui.js==1.14.1"] [/assets/js/jquery-ui.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["jquery.js==3.7.1"] [/assets/js/jquery.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["jszip.js==3.10.1"] [/assets/js/jszip.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["leaflet.js==unknown"] [/assets/js/leaflet.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["link_id.js==unknown"] [/link_id.js]
[source-map-js:name] [http] [low] Found on 1 URLs ["bundle.js.map"] [/assets/js/bundle.js]
[source-map-js:file] [http] [low] Found on 1 URLs [/assets/js/bundle.js.map]

[INF] Running nuclei with tags: [jsext] against 18 targets
[javascript-library] [code] [info] Found on 1 URLs ["alpine.min.js==2.8.2"] [/libs]
[javascript-library] [code] [info] Found on 1 URLs ["aos.js==2.3.4"] [/libs]
[javascript-library] [code] [info] Found on 1 URLs ["axios.min.js==unknown"] [/libs]
[javascript-library] [code] [info] Found on 1 URLs ["bootstrap.min.js==unknown"] [/libs]
[javascript-library] [code] [info] Found on 1 URLs ["chart.js==unknown"] [/libs]
[javascript-library] [code] [info] Found on 1 URLs ["d3.min.js==7.8.1"] [/libs]
[javascript-library] [code] [info] Found on 1 URLs ["gsap.min.js==3.11.1"] [/libs]
[javascript-library] [code] [info] Found on 1 URLs ["jquery-3.6.0.min.js==3.6.0"] [/libs]
[javascript-library] [code] [info] Found on 1 URLs ["jquery-ui.min.js==1.12.1"] [/libs]
[javascript-library] [code] [info] Found on 1 URLs ["jquery.validate.min.js==1.19.3"] [/libs]
[javascript-library] [code] [info] Found on 1 URLs ["leaflet.min.js==1.7.1"] [/libs]
[javascript-library] [code] [info] Found on 1 URLs ["lodash.min.js==4.17.21"] [/libs]
[javascript-library] [code] [info] Found on 1 URLs ["moment.min.js==2.29.1"] [/libs]
[javascript-library] [code] [info] Found on 1 URLs ["popper.min.js==2.11.6"] [/libs]
[javascript-library] [code] [info] Found on 1 URLs ["react-dom.production.min.js==unknown"] [/libs]
[javascript-library] [code] [info] Found on 1 URLs ["react.production.min.js==unknown"] [/libs]
[javascript-library] [code] [info] Found on 1 URLs ["socket.io.min.js==4.3.2"] [/libs]
[javascript-library] [code] [info] Found on 1 URLs ["vue.min.js==2.6.14"] [/libs]
```

This will display the following results on the test WordPress instance.

```
                     __     _
   ____  __  _______/ /__  (_)
  / __ \/ / / / ___/ / _ \/ /
 / / / / /_/ / /__/ /  __/ /
/_/ /_/\__,_/\___/_/\___/_/ ng  v1.3.0

		github.com/oneaudit

[INF] Running nuclei with tags: [html] against 53 targets
[INF] Temporary file created: /tmp/swagger.yaml2790815807

[INF] Running nuclei with tags: [generic] against 90 targets
[INF] Temporary file created: /tmp/swagger.yaml4026525154
[cgi-bin:folder] [http] [medium] Found on 1 URLs [/cgi-bin/]
[cookie-detect] [http] [info] Found on 4 URLs ["wordpress_test_cookie"] [/mysecretlogin/]
[cookie-detect] [http] [info] Found on 1 URLs ["wp_lang"] [/mysecretlogin/]
[cookies-without-httponly] [javascript] [info] Found on 2 URLs ["wordpress_test_cookie"] [/mysecretlogin/]
[cookies-without-secure] [javascript] [info] Found on 2 URLs ["wordpress_test_cookie"] [/mysecretlogin/]
[cors-detect:checked] [http] [info] Found on 1 URLs [header:User-Agent] [GET] [/wp-admin/admin-ajax.php]
[http-missing-security-headers:content-security-policy] [http] [info] Found on 74 URLs [/, /.wp-signup.swp, //, ...]
[http-missing-security-headers:permissions-policy] [http] [info] Found on 74 URLs [/, /.wp-signup.swp, //, ...]
[http-missing-security-headers:referrer-policy] [http] [info] Found on 74 URLs [/, /.wp-signup.swp, //, ...]
[http-missing-security-headers:strict-transport-security] [http] [info] Found on 74 URLs [/, /.wp-signup.swp, //, ...]
[http-missing-security-headers:x-content-type-options] [http] [info] Found on 74 URLs [/, /.wp-signup.swp, //, ...]
[known-403-page:apache] [http] [info] Found on 8 URLs [/*, /:, /wp-admin/js/, ...]
[known-404-page:apache] [http] [info] Found on 1 URLs [/wp-json/]
[known-404-page:wordpress] [http] [info] Found on 5 URLs [///mysecretlogin/, //mysecretlogin/, /404/, /comments/]
[known-500-page:wordpress] [http] [info] Found on 1 URLs [/wp-config-sample.php]
[license-file] [http] [info] Found on 4 URLs [/LICENSE, /LICENSE.txt, /license, /license.txt]
[open-redirect:detect] [javascript] [low] Found on 1 URLs ["dir"] [//wp-admin/load-styles.php]
[options-method-generic] [http] [info] Found on 21 URLs [/LICENSE, /wp-admin/css/, /wp-admin/css/forms.min.css, ...]
[options-method-non-generic] [javascript] [info] Found on 18 URLs ["POST"] [/wp-content/themes/twentytwentyfive/assets/fonts/fira-code/fira-code]
[options-method-non-generic] [javascript] [info] Found on 17 URLs ["TRACE"] [/wp-content/themes/twentytwentyfive/assets/fonts/fira-code/fira-code]
[os-detect:windows] [http] [info] Found on 20 URLs [//mysecretlogin/, //wp-admin/load-styles.php, //wp-admin/maint/repair.php, ...]
[robots-txt-non-generic] [http] [info] Found on 2 URLs ["673fe3c92956f2618d0d11e0a9f0be1023e8a63ca3fbe977816ced034e767534"] [/robots.txt]
[server-errors:php] [http] [info] Found on 2 URLs [/wp-settings, /wp-settings.php]
[stacktrace:generic] [http] [low] Found on 1 URLs [/wp-content/debug.log]
[stacktrace:php] [http] [low] Found on 1 URLs [/wp-settings.php]
[tech-detect:apache] [http] [info] Found on 44 URLs [/, //, /////, ...]
[tech-detect:gravatar] [http] [info] Found on 1 URLs [/2025/02/14/bonjour-tout-le-monde/]
[tech-detect:php] [http] [info] Found on 33 URLs [/, //, /////, ...]
[tech-version:apache] [http] [info] Found on 62 URLs ["2.4.59"] [/, //, ////, ...]
[tech-version:cms] [http] [info] Found on 13 URLs ["WordPress 6.7.2"] [/, //, ////, ...]
[tech-version:mod_fcgid] [http] [info] Found on 62 URLs ["2.3.10"] [/, //, ////, ...]
[tech-version:php] [http] [info] Found on 62 URLs ["8.3.6"] [/, //, ////, ...]

[INF] Running nuclei with tags: [wordpress] against 90 targets
[INF] Temporary file created: /tmp/swagger.yaml869050629
[tech-detect:mysql] [http] [info] Found on 79 URLs [/, //, ///, ...]
[tech-detect:php] [http] [info] Found on 79 URLs [/, //, ///, ...]
[tech-detect:wordpress] [http] [info] Found on 79 URLs [/, //, ///, ...]
[wordpress-config] [http] [info] Found on 1 URLs [/wp-config.bak]
[wordpress-debug-log] [http] [low] Found on 1 URLs ["3987 bytes"] [/wp-content/debug.log]
[wordpress-login:login] [http] [info] Found on 1 URLs [/mysecretlogin/]
[wordpress-login:register] [http] [info] Found on 1 URLs [/mysecretlogin/]
[wordpress-plugins:wordpress-seo] [http] [info] Found on 17 URLs ["24.5"] [/, //, ////, ...]
[wordpress-repair] [http] [low] Found on 1 URLs [/wp-admin/maint/repair.php]
[wordpress-themes-detect] [http] [info] Found on 8 URLs ["twentytwentyfive"] [/wp-content/themes/twentytwentyfive/, /wp-content/themes/twentytwentyfive/assets/, /wp-content/themes/twentytwentyfive/assets/fonts/, ...]
[wordpress-users:author] [http] [info] Found on 1 URLs ["testwp"] [/]
[wordpress-users:rdf] [http] [info] Found on 12 URLs ["testwp"] [//rdf/, /2025/02/14/bonjour-tout-le-monde/feed/rdf/, /2025/02/14/feed/rdf/, ...]
[wordpress-users:usernames] [http] [low] Found on 1 URLs ["testwp"] [/]
[wordpress-users:yoast] [http] [info] Found on 2 URLs ["testwp"] [/author-sitemap.xml]
[wordpress-version:by_css] [http] [info] Found on 22 URLs ["6.7.2"] [/, //, ////, ...]
[wordpress-version:by_generator] [http] [info] Found on 23 URLs ["6.7.2"] [/, //, ////, ...]
[wordpress-version:by_js] [http] [info] Found on 17 URLs ["6.7.2"] [/, //, ////, ...]
[wordpress-wpjson:routes] [http] [info] Found on 1 URLs ["50 routes"] [/]
[wordpress-xmlrpc:listmethods] [http] [info] Found on 3 URLs ["80 methods"] [/xmlrpc.php, /xmlrpc.php/xmlrpc.php, /xmlrpc/xmlrpc.php]
[wordpress-xmlrpc] [http] [info] Found on 4 URLs [//xmlrpc.php, /xmlrpc, /xmlrpc.php]
[wordpress-xmlrpc:pingback] [http] [medium] Found on 4 URLs [/xmlrpc.php]

[INF] Running nuclei with tags: [javascript] against 2 targets
[INF] Temporary file created: /tmp/swagger.yaml1590506295
[javascript-library] [javascript] [info] Found on 1 URLs ["password-strength-meter.min.js==6.7.2"] [/wp-admin/js/password-strength-meter.min.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["user-profile.min.js==6.7.2"] [/wp-admin/js/user-profile.min.js]

[INF] Running nuclei with tags: [jsext] against 9 targets
[javascript-library] [code] [info] Found on 1 URLs ["a11y.min.js==unknown"] [/mysecretlogin/]
[javascript-library] [code] [info] Found on 1 URLs ["dom-ready.min.js==unknown"] [/mysecretlogin/]
[javascript-library] [code] [info] Found on 1 URLs ["i18n.min.js==unknown"] [/mysecretlogin/]
[javascript-library] [code] [info] Found on 1 URLs ["load-scripts.php==unknown"] [/mysecretlogin/]
[javascript-library] [code] [info] Found on 1 URLs ["password-strength-meter.min.js==unknown"] [/mysecretlogin/]
[javascript-library] [code] [info] Found on 1 URLs ["underscore.min.js==unknown"] [/mysecretlogin/]
[javascript-library] [code] [info] Found on 1 URLs ["user-profile.min.js==unknown"] [/mysecretlogin/]
[javascript-library] [code] [info] Found on 7 URLs ["view.min.js==unknown"] [/, /2025/, /2025/02/, ...]
[javascript-library] [code] [info] Found on 1 URLs ["wp-util.min.js==unknown"] [/mysecretlogin/]
```

This will display the following results on the test GLPI instance.

```
                     __     _
   ____  __  _______/ /__  (_)
  / __ \/ / / / ___/ / _ \/ /
 / / / / /_/ / /__/ /  __/ /
/_/ /_/\__,_/\___/_/\___/_/ ng  v1.3.0

		github.com/oneaudit

[INF] Running nuclei with tags: [glpi] against 241 targets
[INF] Temporary file created: /tmp/swagger.yaml1521432459
[glpi-config:file] [http] [low] Found on 2 URLs [/status.php]
[glpi-directory-listing] [http] [high] Found on 1 URLs [/files/_sessions/]
[glpi-login] [http] [medium] Found on 2 URLs [/, /public/]
[glpi-login:default-credentials] [http] [high] Found on 1 URLs [pass="glpi",user="glpi"] [/front/login.php]
[glpi-version:by_link] [http] [info] Found on 2 URLs ["10.0.10"] [/version/]
[glpi-version:by_uri] [http] [info] Found on 1 URLs ["10.0.10"] [/version/10.0.10]
[tech-detect:glpi] [http] [info] Found on 22 URLs [/, //index.php, /Index, ...]
[tech-detect:mysql] [http] [info] Found on 22 URLs [/, //index.php, /Index, ...]
[tech-detect:php] [http] [info] Found on 22 URLs [/, //index.php, /Index, ...]
[tech-detect:scss] [http] [info] Found on 22 URLs [/, //index.php, /Index, ...]
[tech-detect:twig] [http] [info] Found on 22 URLs [/, //index.php, /Index, ...]

[INF] Running nuclei with tags: [html] against 76 targets
[INF] Temporary file created: /tmp/swagger.yaml1887317488
[directory-listing] [http] [info] Found on 35 URLs [/bin/, /css/, /css/includes/, ...]
[form-detect] [http] [info] Found on 7 URLs ["/front/login.php"] [/, //index.php, /Index, ...]
[form-detect] [http] [info] Found on 1 URLs ["/public/front/login.php"] [/public/]

[INF] Running nuclei with tags: [generic] against 241 targets
[INF] Temporary file created: /tmp/swagger.yaml3295598526
[cgi-bin:folder] [http] [medium] Found on 1 URLs [/cgi-bin/]
[composer:vendor] [http] [medium] Found on 1 URLs [/vendor/]
[cookie-detect] [http] [info] Found on 22 URLs ["glpi_8c1bab8cdfb8e857ff0e6290f00f8ed5"] [/, ///front/locale.php, //front/locale.php, ...]
[cookies-without-httponly] [javascript] [info] Found on 22 URLs ["glpi_8c1bab8cdfb8e857ff0e6290f00f8ed5"] [/index.php?error=3]
[cookies-without-secure] [javascript] [info] Found on 22 URLs ["glpi_8c1bab8cdfb8e857ff0e6290f00f8ed5"] [/status/application/json.bak]
[css-source-file:scss] [http] [info] Found on 63 URLs [/css/includes/_base.scss, /css/includes/_fonts.scss, /css/includes/_global-variables.scss, ...]
[email-detect] [http] [info] Found on 2 URLs ["wowohoo@qq.com"] [//public/lib/fuzzy.min.js, /public/lib/fuzzy.min.js]
[favicon-detect:glpi] [http] [info] Found on 1 URLs [/pics/favicon.ico]
[http-missing-security-headers:content-security-policy] [http] [info] Found on 242 URLs [/, ///front/locale.php, //front/locale.php, ...]
[http-missing-security-headers:permissions-policy] [http] [info] Found on 242 URLs [/, ///front/locale.php, //front/locale.php, ...]
[http-missing-security-headers:referrer-policy] [http] [info] Found on 242 URLs [/, ///front/locale.php, //front/locale.php, ...]
[http-missing-security-headers:strict-transport-security] [http] [info] Found on 242 URLs [/, ///front/locale.php, //front/locale.php, ...]
[http-missing-security-headers:x-content-type-options] [http] [info] Found on 242 URLs [/, ///front/locale.php, //front/locale.php, ...]
[ide-config:jetbrains] [http] [low] Found on 1 URLs [/.idea/workspace.xml]
[known-403-page:apache] [http] [info] Found on 7 URLs [/*, /:, /js/"+CFG_GLPI.root_doc+e.original.url+", ...]
[known-404-page:apache] [http] [info] Found on 3 URLs [/PAGENOTFOUND, /PAGENOTFOUND/, /public/front/login.php]
[license-file] [http] [info] Found on 2 URLs [/LICENSE, /license]
[log-file] [http] [low] Found on 1 URLs [/files/_log/php-errors.log]
[open-redirect:detect] [javascript] [low] Found on 2 URLs ["domain"] [/front/locale.php]
[options-method-generic] [http] [info] Found on 209 URLs [//js/common.min.js, //js/fileupload.min.js, //js/flatpickr_buttons_plugin.min.js, ...]
[options-method-non-generic] [javascript] [info] Found on 209 URLs ["POST"] [/js/impact.js]
[options-method-non-generic] [javascript] [info] Found on 209 URLs ["TRACE"] [/js/impact.js]
[os-detect:windows] [http] [info] Found on 244 URLs [/, ///front/locale.php, //front/locale.php, ...]
[readme-file] [http] [info] Found on 1 URLs [/README.md]
[server-errors:php] [http] [info] Found on 4 URLs [//front/locale.php, //public/lib/base.min.js, /files/_log/php-errors.log, /public/lib/base.min.js]
[stacktrace:php] [http] [low] Found on 10 URLs [///front/locale.php, /status, /status.php, ...]
[tech-detect:apache] [http] [info] Found on 244 URLs [/, ///front/locale.php, //front/locale.php, ...]
[tech-detect:glpi] [http] [info] Found on 9 URLs [/, //index.php, /Index, ...]
[tech-detect:php] [http] [info] Found on 30 URLs [/, ///front/locale.php, //front/locale.php, ...]
[tech-version:apache] [http] [info] Found on 244 URLs ["2.4.59"] [/, ///front/locale.php, //front/locale.php, ...]
[tech-version:mod_fcgid] [http] [info] Found on 244 URLs ["2.3.10"] [/, ///front/locale.php, //front/locale.php, ...]
[tech-version:php] [http] [info] Found on 244 URLs ["8.3.6"] [/, ///front/locale.php, //front/locale.php, ...]

[INF] Running nuclei with tags: [javascript] against 64 targets
[INF] Temporary file created: /tmp/swagger.yaml2318201103
...SNIP...
[source-map-js] [http] [low] Found on 8 URLs [//public/lib/leaflet.min.js, //public/lib/photoswipe.min.js, /js/planning.js, ...]
```

This will display the following results on the test Django instance.

```
                     __     _
   ____  __  _______/ /__  (_)
  / __ \/ / / / ___/ / _ \/ /
 / / / / /_/ / /__/ /  __/ /
/_/ /_/\__,_/\___/_/\___/_/ ng  v1.3.0

		github.com/oneaudit

[INF] Running nuclei with tags: [django] against 7 targets
[INF] Temporary file created: /tmp/swagger.yaml4107642740
[django-admin-panel] [http] [info] Found on 2 URLs [/django/admin/login/]
[django-debug:database] [http] [medium] Found on 1 URLs ["sqlite3"] [/django/admin/login/]
[django-debug:django-version] [http] [medium] Found on 1 URLs ["5.0.0"] [/django/admin/login/]
[django-debug:endpoints] [http] [medium] Found on 7 URLs ["admin/"] [/django/PAGENOTFOUND/, /django/admin/, /django/admin/login/PAGENOTFOUND/, ...]
[django-debug:endpoints] [http] [medium] Found on 7 URLs ["api/redoc/"] [/django/PAGENOTFOUND/, /django/admin/, /django/admin/login/PAGENOTFOUND/, ...]
[django-debug:python-version] [http] [medium] Found on 1 URLs ["3.09.0"] [/django/admin/login/]
[django-debug:variables] [http] [medium] Found on 1 URLs ["3 variables"] [/django/admin/login/]
[tech-detect:django] [http] [info] Found on 8 URLs [/django/, /django/admin/, /django/admin/login/, ...]
[tech-detect:python] [http] [info] Found on 8 URLs [/django/, /django/admin/, /django/admin/login/, ...]

[INF] Running nuclei with tags: [generic] against 7 targets
[INF] Temporary file created: /tmp/swagger.yaml465723831
[default-pages:django-install] [http] [low] Found on 1 URLs [/django/]
[file-inclusion:detect] [javascript] [low] Found on 1 URLs ["next"] [/django/admin/login/]
[http-missing-security-headers:content-security-policy] [http] [info] Found on 8 URLs [/django/, /django/admin/, /django/admin/login/, ...]
[http-missing-security-headers:permissions-policy] [http] [info] Found on 8 URLs [/django/, /django/admin/, /django/admin/login/, ...]
[http-missing-security-headers:referrer-policy] [http] [info] Found on 8 URLs [/django/, /django/admin/, /django/admin/login/, ...]
[http-missing-security-headers:strict-transport-security] [http] [info] Found on 8 URLs [/django/, /django/admin/, /django/admin/login/, ...]
[http-missing-security-headers:x-content-type-options] [http] [info] Found on 8 URLs [/django/, /django/admin/, /django/admin/login/, ...]
[known-404-page:django] [http] [info] Found on 4 URLs [/django/PAGENOTFOUND, /django/PAGENOTFOUND/, /django/admin/, /django/install/]
[known-500-page:django] [http] [info] Found on 1 URLs [/django/admin/login/]
[server-errors:django] [http] [info] Found on 3 URLs [/django/admin/, /django/admin/login/, /django/install/]
[tech-detect:django] [http] [info] Found on 1 URLs [/django/admin/login/]

[INF] Running nuclei with tags: [html] against 7 targets
[INF] Temporary file created: /tmp/swagger.yaml833361855
[form-detect] [http] [info] Found on 1 URLs ["/django/admin/login/?next=/admin/"] [/django/admin/login/]
[html-comments-detect] [javascript] [info] Found on 1 URLs ["<!-- SNIP -->"] [/django/api/]
```