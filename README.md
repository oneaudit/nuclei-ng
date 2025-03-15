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
[cookie-detect] [http] [info] Found on 5 URLs ["wordpress_test_cookie"] [/mysecretlogin/]
[cookie-detect] [http] [info] Found on 1 URLs ["wp_lang"] [/mysecretlogin/]
[cookies-without-httponly] [javascript] [info] Found on 2 URLs ["wordpress_test_cookie"] [/mysecretlogin/]
[cookies-without-secure] [javascript] [info] Found on 2 URLs ["wordpress_test_cookie"] [/mysecretlogin/]
[cors-detect:checked] [http] [info] Found on 1 URLs [header:User-Agent] [GET] [/wp-admin/admin-ajax.php]
[http-missing-security-headers:content-security-policy] [http] [info] Found on 90 URLs [/, /.wp-signup.swp, //, ...]
[http-missing-security-headers:permissions-policy] [http] [info] Found on 90 URLs [/, /.wp-signup.swp, //, ...]
[http-missing-security-headers:referrer-policy] [http] [info] Found on 90 URLs [/, /.wp-signup.swp, //, ...]
[http-missing-security-headers:strict-transport-security] [http] [info] Found on 90 URLs [/, /.wp-signup.swp, //, ...]
[http-missing-security-headers:x-content-type-options] [http] [info] Found on 90 URLs [/, /.wp-signup.swp, //, ...]
[known-403-page:apache] [http] [info] Found on 12 URLs [/*, /:, /wp-admin/css/, ...]
[known-404-page:apache] [http] [info] Found on 1 URLs [/wp-json/]
[known-404-page:wordpress] [http] [info] Found on 7 URLs [///mysecretlogin/, //mysecretlogin/, /404/, ...]
[known-500-page:wordpress] [http] [info] Found on 4 URLs [/wp-config-sample.php, /wp-signup, /wp-signup.bak, /wp-signup.old]
[license-file] [http] [info] Found on 4 URLs [/LICENSE, /LICENSE.txt, /license, /license.txt]
[options-method-generic] [http] [info] Found on 30 URLs [//wp-admin/css/install.css, /LICENSE, /LICENSE.txt, ...]
[options-method-non-generic] [javascript] [info] Found on 26 URLs ["POST"] [/wp-content/debug.log]
[options-method-non-generic] [javascript] [info] Found on 24 URLs ["TRACE"] [/wp-content/debug.log]
[robots-txt-non-generic] [http] [info] Found on 2 URLs ["673fe3c92956f2618d0d11e0a9f0be1023e8a63ca3fbe977816ced034e767534"] [/robots.txt]
[server-errors:php] [http] [info] Found on 2 URLs [/wp-settings, /wp-settings.php]
[stacktrace:generic] [http] [low] Found on 1 URLs [/wp-content/debug.log]
[stacktrace:php] [http] [low] Found on 4 URLs [/.wp-signup.swp, /wp-settings, /wp-settings.php, /wp-signup.bak]
[tech-detect:apache] [http] [info] Found on 93 URLs [/, //, ///, ...]
[tech-detect:gravatar] [http] [info] Found on 1 URLs [/2025/02/14/bonjour-tout-le-monde/]
[tech-detect:php] [http] [info] Found on 64 URLs [/, //, ///, ...]
[tech-version:apache] [http] [info] Found on 93 URLs ["2.4.59"] [/, //, ///, ...]
[tech-version:cms] [http] [info] Found on 17 URLs ["WordPress 6.7.2"] [/, //, ////, ...]
[tech-version:mod_fcgid] [http] [info] Found on 93 URLs ["2.3.10"] [/, //, ///, ...]
[tech-version:os] [http] [info] Found on 93 URLs ["Win64"] [/, //, ///, ...]
[tech-version:php] [http] [info] Found on 93 URLs ["8.3.6"] [/, //, ///, ...]

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

[INF] Running nuclei with tags: [wordpress] against 90 targets
[INF] Temporary file created: /tmp/swagger.yaml869050629
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
[wordpress-xmlrpc:pingback] [http] [medium] Found on 3 URLs [/xmlrpc.php]
```

This will display the following results on the test GLPI instance.

```
                     __     _
   ____  __  _______/ /__  (_)
  / __ \/ / / / ___/ / _ \/ /
 / / / / /_/ / /__/ /  __/ /
/_/ /_/\__,_/\___/_/\___/_/ ng  v1.3.0

		github.com/oneaudit

[INF] Running nuclei with tags: [html] against 76 targets
[INF] Temporary file created: /tmp/swagger.yaml1887317488
[directory-listing] [http] [info] Found on 35 URLs [/bin/, /css/, /css/includes/, ...]
[form-detect] [http] [info] Found on 7 URLs ["/front/login.php"] [/, //index.php, /Index, ...]
[form-detect] [http] [info] Found on 1 URLs ["/public/front/login.php"] [/public/]

[INF] Running nuclei with tags: [generic] against 240 targets
[INF] Temporary file created: /tmp/swagger.yaml3295598526
[cookie-detect] [http] [info] Found on 22 URLs ["glpi_8c1bab8cdfb8e857ff0e6290f00f8ed5"] [/, ///front/locale.php, //front/locale.php, ...]
[cookies-without-httponly] [javascript] [info] Found on 22 URLs ["glpi_8c1bab8cdfb8e857ff0e6290f00f8ed5"] [/front/cron.php]
[cookies-without-secure] [javascript] [info] Found on 22 URLs ["glpi_8c1bab8cdfb8e857ff0e6290f00f8ed5"] [/status.php]
[email-detect] [http] [info] Found on 2 URLs ["wowohoo@qq.com"] [//public/lib/fuzzy.min.js, /public/lib/fuzzy.min.js]
[favicon-detect:glpi] [http] [info] Found on 1 URLs [/pics/favicon.ico]
[http-missing-security-headers:content-security-policy] [http] [info] Found on 241 URLs [/, ///front/locale.php, //front/locale.php, ...]
[http-missing-security-headers:permissions-policy] [http] [info] Found on 241 URLs [/, ///front/locale.php, //front/locale.php, ...]
[http-missing-security-headers:referrer-policy] [http] [info] Found on 241 URLs [/, ///front/locale.php, //front/locale.php, ...]
[http-missing-security-headers:strict-transport-security] [http] [info] Found on 241 URLs [/, ///front/locale.php, //front/locale.php, ...]
[http-missing-security-headers:x-content-type-options] [http] [info] Found on 241 URLs [/, ///front/locale.php, //front/locale.php, ...]
[known-403-page:apache] [http] [info] Found on 7 URLs [/*, /:, /js/"+CFG_GLPI.root_doc+e.original.url+", ...]
[known-404-page:apache] [http] [info] Found on 3 URLs [/PAGENOTFOUND, /PAGENOTFOUND/, /public/front/login.php]
[license-file] [http] [info] Found on 2 URLs [/LICENSE, /license]
[open-redirect:detect] [javascript] [low] Found on 2 URLs ["domain"] [/front/locale.php]
[options-method-generic] [http] [info] Found on 208 URLs [//js/common.min.js, //js/fileupload.min.js, //js/flatpickr_buttons_plugin.min.js, ...]
[options-method-non-generic] [javascript] [info] Found on 208 URLs ["POST"] [/public/lib/base.min.js]
[options-method-non-generic] [javascript] [info] Found on 208 URLs ["TRACE"] [/public/lib/base.min.js]
[readme-file] [http] [info] Found on 1 URLs [/README.md]
[server-errors:php] [http] [info] Found on 3 URLs [//front/locale.php, //public/lib/base.min.js, /public/lib/base.min.js]
[stacktrace:php] [http] [low] Found on 10 URLs [///front/locale.php, /status, /status.php, ...]
[tech-detect:apache] [http] [info] Found on 243 URLs [/, ///front/locale.php, //front/locale.php, ...]
[tech-detect:php] [http] [info] Found on 30 URLs [/, ///front/locale.php, //front/locale.php, ...]
[tech-version:apache] [http] [info] Found on 243 URLs ["2.4.59"] [/, ///front/locale.php, //front/locale.php, ...]
[tech-version:mod_fcgid] [http] [info] Found on 243 URLs ["2.3.10"] [/, ///front/locale.php, //front/locale.php, ...]
[tech-version:os] [http] [info] Found on 243 URLs ["Win64"] [/, ///front/locale.php, //front/locale.php, ...]
[tech-version:php] [http] [info] Found on 243 URLs ["8.3.6"] [/, ///front/locale.php, //front/locale.php, ...]

[INF] Running nuclei with tags: [javascript] against 64 targets
[INF] Temporary file created: /tmp/swagger.yaml2318201103
[javascript-library] [javascript] [info] Found on 1 URLs ["ContentTemplatesParameters.js==unknown"] [/js/RichText/ContentTemplatesParameters.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["ContentTemplatesParameters.min.js==unknown"] [/js/RichText/ContentTemplatesParameters.min.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["FaIconSelector.js==unknown"] [/js/Forms/FaIconSelector.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["FaIconSelector.min.js==unknown"] [/js/Forms/FaIconSelector.min.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["ResultsView.min.js==2b8ebd59a8d46b694cc14c813395db6f6b21a878"] [//js/modules/Search/ResultsView.min.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["ResultsView.min.js==unknown"] [/js/modules/Search/ResultsView.min.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["Table.min.js==2b8ebd59a8d46b694cc14c813395db6f6b21a878"] [//js/modules/Search/Table.min.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["Table.min.js==unknown"] [/js/modules/Search/Table.min.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["UserMention.js==unknown"] [/js/RichText/UserMention.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["UserMention.min.js==unknown"] [/js/RichText/UserMention.min.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["base.min.js==1.0.0"] [//public/lib/base.min.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["base.min.js==2b8ebd59a8d46b694cc14c813395db6f6b21a878"] [/public/lib/base.min.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["cable.js==unknown"] [/js/cable.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["cable.min.js==unknown"] [/js/cable.min.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["clipboard.js==unknown"] [/js/clipboard.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["clipboard.min.js==unknown"] [/js/clipboard.min.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["common.js==unknown"] [/js/common.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["common.min.js==2b8ebd59a8d46b694cc14c813395db6f6b21a878"] [/js/common.min.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["common.min.js==unknown"] [//js/common.min.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["dashboard.js==unknown"] [/js/dashboard.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["dashboard.min.js==unknown"] [/js/dashboard.min.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["fileupload.js==unknown"] [/js/fileupload.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["fileupload.min.js==2b8ebd59a8d46b694cc14c813395db6f6b21a878"] [/js/fileupload.min.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["fileupload.min.js==unknown"] [//js/fileupload.min.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["flatpickr.min.js==2b8ebd59a8d46b694cc14c813395db6f6b21a878"] [//public/lib/flatpickr.min.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["flatpickr.min.js==4.6.11"] [/public/lib/flatpickr.min.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["flatpickr_buttons_plugin.js==unknown"] [/js/flatpickr_buttons_plugin.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["flatpickr_buttons_plugin.min.js==2b8ebd59a8d46b694cc14c813395db6f6b21a878"] [/js/flatpickr_buttons_plugin.min.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["flatpickr_buttons_plugin.min.js==unknown"] [//js/flatpickr_buttons_plugin.min.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["fuzzy.min.js==2b8ebd59a8d46b694cc14c813395db6f6b21a878"] [//public/lib/fuzzy.min.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["fuzzy.min.js==3.8.8"] [/public/lib/fuzzy.min.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["fuzzysearch.js==unknown"] [/js/fuzzysearch.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["fuzzysearch.min.js==2b8ebd59a8d46b694cc14c813395db6f6b21a878"] [/js/fuzzysearch.min.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["fuzzysearch.min.js==unknown"] [//js/fuzzysearch.min.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["glpi_dialog.js==unknown"] [/js/glpi_dialog.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["glpi_dialog.min.js==2b8ebd59a8d46b694cc14c813395db6f6b21a878"] [/js/glpi_dialog.min.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["glpi_dialog.min.js==unknown"] [//js/glpi_dialog.min.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["impact.js==unknown"] [/js/impact.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["impact.min.js==unknown"] [/js/impact.min.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["jquery-file-upload.min.js==2b8ebd59a8d46b694cc14c813395db6f6b21a878"] [//public/lib/jquery-file-upload.min.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["jquery-file-upload.min.js==unknown"] [/public/lib/jquery-file-upload.min.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["leaflet.min.js==2b8ebd59a8d46b694cc14c813395db6f6b21a878"] [/public/lib/leaflet.min.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["leaflet.min.js==unknown"] [//public/lib/leaflet.min.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["log_filters.js==unknown"] [/js/log_filters.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["log_filters.min.js==2b8ebd59a8d46b694cc14c813395db6f6b21a878"] [/js/log_filters.min.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["log_filters.min.js==unknown"] [//js/log_filters.min.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["marketplace.js==unknown"] [/js/marketplace.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["marketplace.min.js==unknown"] [/js/marketplace.min.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["misc.js==unknown"] [/js/misc.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["misc.min.js==2b8ebd59a8d46b694cc14c813395db6f6b21a878"] [/js/misc.min.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["misc.min.js==unknown"] [//js/misc.min.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["notifications_ajax.js==unknown"] [/js/notifications_ajax.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["notifications_ajax.min.js==unknown"] [/js/notifications_ajax.min.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["photoswipe.min.js==2b8ebd59a8d46b694cc14c813395db6f6b21a878"] [//public/lib/photoswipe.min.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["photoswipe.min.js==unknown"] [/public/lib/photoswipe.min.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["planning.js==unknown"] [/js/planning.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["planning.min.js==unknown"] [/js/planning.min.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["rack.js==unknown"] [/js/rack.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["rack.min.js==unknown"] [/js/rack.min.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["reservations.js==unknown"] [/js/reservations.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["reservations.min.js==unknown"] [/js/reservations.min.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["webkit_fix.js==unknown"] [/js/webkit_fix.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["webkit_fix.min.js==2b8ebd59a8d46b694cc14c813395db6f6b21a878"] [/js/webkit_fix.min.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["webkit_fix.min.js==unknown"] [//js/webkit_fix.min.js]
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
[server-errors:django] [http] [info] Found on 2 URLs [/django/admin/, /django/install/]
[tech-detect:django] [http] [info] Found on 1 URLs [/django/admin/login/]

[INF] Running nuclei with tags: [html] against 7 targets
[INF] Temporary file created: /tmp/swagger.yaml833361855
[form-detect] [http] [info] Found on 1 URLs ["/django/admin/login/?next=/admin/"] [/django/admin/login/]
[html-comments-detect] [javascript] [info] Found on 1 URLs ["<!-- SNIP -->"] [/django/api/]
```

## Current issues

When access to a route is blocked, but an alternative route works, common checks are not executed on the alternative route unless the route is added to the list.

- [ ] Jetty (Java), XXX (Python), WordPress (MySQL), etc.
- [ ] MIA

```
"Yoast SEO:24.5"
"WordPress Site Editor"
"WordPress Block Editor"
SASS/SCSS
.idea (path exposure, etc.)
.twig
```