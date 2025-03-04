<h1 align="center">
  Nuclei-ng
</h1>
<h4 align="center">Nuclei-ng is a wrapper of <a href="https://github.com/projectdiscovery/nuclei">nuclei</a> from Project Discovery.</h4>

## Installation ✍️

nuclei-ng requires **Go 1.22.5+** to install successfully.

```console
CGO_ENABLED=1 go install github.com/oneaudit/nuclei-ng/cmd/nuclei-ng@latest
```

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
/_/ /_/\__,_/\___/_/\___/_/ ng  v1.1.2

		github.com/oneaudit

[INF] Running nuclei with tags: [html] against 48 targets
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

[INF] Running nuclei with tags: [generic] against 68 targets
[INF] Temporary file created: /tmp/swagger.yaml747212024
[composer-config:composer.json] [http] [low] Found on 2 URLs [/composer.json, /vendor/composer/installed.json]
[composer-vendor] [http] [low] Found on 2 URLs [/vendor/autoload.php, /vendor/composer/installed.json]
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
[git-config:config] [http] [medium] Found on 3 URLs [/.git/HEAD, /.git/config, /.git/index]
[git-config:folder] [http] [medium] Found on 1 URLs [/.git/]
[git-config:ignore] [http] [medium] Found on 1 URLs [/.gitignore]
[git-config:logs] [http] [medium] Found on 1 URLs [/.git/logs/HEAD]
[http-missing-security-headers:content-security-policy] [http] [info] Found on 67 URLs [/, /.git/, /.git/HEAD, ...]
[http-missing-security-headers:permissions-policy] [http] [info] Found on 68 URLs [/, /.git/, /.git/HEAD, ...]
[http-missing-security-headers:referrer-policy] [http] [info] Found on 68 URLs [/, /.git/, /.git/HEAD, ...]
[http-missing-security-headers:strict-transport-security] [http] [info] Found on 68 URLs [/, /.git/, /.git/HEAD, ...]
[http-missing-security-headers:x-content-type-options] [http] [info] Found on 46 URLs [/, /.git/, /.git/logs/, ...]
[http-suspicious-request-headers] [javascript] [info] Found on 1 URLs ["X-Api-Key: MYS3cr374P1K3y"] [/ng_hidden_spy]
[http-suspicious-response-headers] [javascript] [info] Found on 1 URLs ["X-Entrypoint: /empty_page/1234/"] [/js-event-link]
[known-403-page:aspnet] [http] [info] Found on 1 URLs [/aspNetErrorPage]
[known-403-page:custom] [http] [info] Found on 3 URLs [/.git/, /assets/js/, /vendor]
[known-404-page:custom] [http] [info] Found on 8 URLs [/.git/logs/, /.well-known/, /PAGENOTFOUND, ...]
[open-redirect-detect:detect] [javascript] [low] Found on 1 URLs ["redirect"] [/re]
[open-redirect-detect:exploit] [http] [low] Found on 1 URLs [query:redirect] [GET] [/]
[options-method-generic] [http] [info] Found on 2 URLs [/ng_hidden_login, /ng_hidden_spy]
[options-method-non-generic] [javascript] [info] Found on 2 URLs ["POST"] [/ng_hidden_spy]
[php-errors] [http] [info] Found on 1 URLs ["Stack Trace:"] [/aspNetErrorPage]
[robots-txt-generic] [http] [info] Found on 1 URLs [/robots.txt]
[robots-txt-non-generic] [http] [info] Found on 1 URLs ["b77a70c632c0be0cea6bbce3dff8a2317392303b0730327589b3485c4b892dcd"] [/icons/robots.txt]
[sitemap-detect] [http] [info] Found on 1 URLs [/sitemap.xml]
[tech-detect:apache] [http] [info] Found on 2 URLs [/apache-v, /php-v]
[tech-detect:jetty] [http] [info] Found on 5 URLs [/icons/, /icons/.robots.txt.swp, /icons/robots.bak, ...]
[tech-detect:jsdelivr] [http] [info] Found on 1 URLs [/libs]
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
[tech-version:nginx] [http] [info] Found on 1 URLs ["1.33.7"] [/nginx-v]
[tech-version:os] [http] [info] Found on 1 URLs ["Debian"] [/php-v]
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
/_/ /_/\__,_/\___/_/\___/_/ ng  v1.1.2

		github.com/oneaudit

[INF] Running nuclei with tags: [html] against 10 targets
[INF] Temporary file created: /tmp/swagger.yaml2790815807
[form-detect] [http] [info] Found on 20 URLs ["http://wp.sec2/"] [/.git/, /.git/HEAD, /.git/config, ...]
[form-detect] [http] [info] Found on 2 URLs ["http://wp.sec2/wp-comments-post.php"] [/2025/02/14/bonjour-tout-le-monde/]
[html-comments-detect] [javascript] [info] Found on 1 URLs ["<!-- #respond -->"] [/2025/02/14/bonjour-tout-le-monde/]
[html-comments-detect] [javascript] [info] Found on 22 URLs ["<!-- / Yoast SEO plugin. -->"] [/docs/]
[html-comments-detect] [javascript] [info] Found on 22 URLs ["<!-- This site is optimized with the Yoast SEO plugin v24.5 - https://yoast.com/wordpress/plugins/seo/ -->"] [/docs/]
[html-comments-detect] [javascript] [info] Found on 1 URLs ["<!-- generator="WordPress/6.7.2" -->"] [/wp-links-opml]

[INF] Running nuclei with tags: [generic] against 22 targets
[INF] Temporary file created: /tmp/swagger.yaml4026525154
[cookie-detect] [http] [info] Found on 1 URLs ["wordpress_test_cookie"] [/mysecretlogin/]
[cookies-without-httponly] [javascript] [info] Found on 1 URLs ["wordpress_test_cookie"] [/mysecretlogin/]
[cookies-without-secure] [javascript] [info] Found on 1 URLs ["wordpress_test_cookie"] [/mysecretlogin/]
[http-missing-security-headers:content-security-policy] [http] [info] Found on 21 URLs [/, /2025/02/14/bonjour-tout-le-monde/, /404/, ...]
[http-missing-security-headers:permissions-policy] [http] [info] Found on 21 URLs [/, /2025/02/14/bonjour-tout-le-monde/, /404/, ...]
[http-missing-security-headers:referrer-policy] [http] [info] Found on 21 URLs [/, /2025/02/14/bonjour-tout-le-monde/, /404/, ...]
[http-missing-security-headers:strict-transport-security] [http] [info] Found on 21 URLs [/, /2025/02/14/bonjour-tout-le-monde/, /404/, ...]
[http-missing-security-headers:x-content-type-options] [http] [info] Found on 21 URLs [/, /2025/02/14/bonjour-tout-le-monde/, /404/, ...]
[known-403-page:apache] [http] [info] Found on 2 URLs [/:, /wp-content/themes/twentytwentyfive/]
[known-404-page:apache] [http] [info] Found on 1 URLs [/wp-json/]
[known-404-page:wordpress] [http] [info] Found on 3 URLs [/404/, /PAGENOTFOUND, /PAGENOTFOUND/]
[known-500-page:wordpress] [http] [info] Found on 1 URLs [/wp-config-sample.php]
[license-file] [http] [info] Found on 1 URLs [/LICENSE]
[options-method-generic] [http] [info] Found on 6 URLs [/LICENSE, /readme.html, /wp-config.bak, ...]
[options-method-non-generic] [javascript] [info] Found on 6 URLs ["POST"] [/readme.html]
[options-method-non-generic] [javascript] [info] Found on 5 URLs ["TRACE"] [/readme.html]
[php-errors] [http] [info] Found on 1 URLs ["Fatal error"] [/wp-settings.php]
[php-errors] [http] [info] Found on 1 URLs ["Stack trace:"] [/wp-content/debug.log]
[robots-txt-non-generic] [http] [info] Found on 1 URLs ["673fe3c92956f2618d0d11e0a9f0be1023e8a63ca3fbe977816ced034e767534"] [/robots.txt]
[stacktrace:php] [http] [low] Found on 1 URLs [/wp-settings.php]
[tech-detect:apache] [http] [info] Found on 21 URLs [/, /2025/02/14/bonjour-tout-le-monde/, /404/, ...]
[tech-detect:gravatar] [http] [info] Found on 1 URLs [/2025/02/14/bonjour-tout-le-monde/]
[tech-detect:php] [http] [info] Found on 15 URLs [/, /2025/02/14/bonjour-tout-le-monde/, /404/, ...]
[tech-version:apache] [http] [info] Found on 21 URLs ["2.4.59"] [/, /2025/02/14/bonjour-tout-le-monde/, /404/, ...]
[tech-version:cms] [http] [info] Found on 3 URLs ["WordPress 6.7.2"] [/, /2025/02/14/bonjour-tout-le-monde/, /404/]
[tech-version:mod_fcgid] [http] [info] Found on 21 URLs ["2.3.10"] [/, /2025/02/14/bonjour-tout-le-monde/, /404/, ...]
[tech-version:os] [http] [info] Found on 21 URLs ["Win64"] [/, /2025/02/14/bonjour-tout-le-monde/, /404/, ...]
[tech-version:php] [http] [info] Found on 21 URLs ["8.3.6"] [/, /2025/02/14/bonjour-tout-le-monde/, /404/, ...]

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

[INF] Running nuclei with tags: [javascript] against 2 targets
[INF] Temporary file created: /tmp/swagger.yaml159050629
[javascript-library] [javascript] [info] Found on 1 URLs ["password-strength-meter.min.js==6.7.2"] [/wp-admin/js/password-strength-meter.min.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["user-profile.min.js==6.7.2"] [/wp-admin/js/user-profile.min.js]

[INF] Running nuclei with tags: [wordpress] against 22 targets
[INF] Temporary file created: /tmp/swagger.yaml869050629
[wordpress-config] [http] [info] Found on 1 URLs [/wp-config.bak]
[wordpress-debug-log] [http] [low] Found on 1 URLs ["3987 bytes"] [/wp-content/debug.log]
[wordpress-login:login] [http] [info] Found on 1 URLs [/mysecretlogin/]
[wordpress-login:register] [http] [info] Found on 1 URLs [/mysecretlogin/]
[wordpress-repair] [http] [low] Found on 1 URLs [/wp-admin/maint/repair.php]
[wordpress-themes-detect] [http] [info] Found on 1 URLs ["twentytwentyfive"] [/wp-content/themes/twentytwentyfive/]
[wordpress-users:author] [http] [info] Found on 1 URLs ["testwp"] [/]
[wordpress-users:rdf] [http] [info] Found on 2 URLs ["testwp"] [/2025/02/14/bonjour-tout-le-monde/feed/rdf/, /feed/rdf/]
[wordpress-users:usernames] [http] [low] Found on 1 URLs ["testwp"] [/]
[wordpress-users:yoast] [http] [info] Found on 2 URLs ["testwp"] [/author-sitemap.xml]
[wordpress-version:by_css] [http] [info] Found on 4 URLs ["6.7.2"] [/, /2025/02/14/bonjour-tout-le-monde/, /404/, /wp-admin/maint/repair.php]
[wordpress-version:by_generator] [http] [info] Found on 3 URLs ["6.7.2"] [/, /2025/02/14/bonjour-tout-le-monde/, /404/]
[wordpress-version:by_js] [http] [info] Found on 3 URLs ["6.7.2"] [/, /2025/02/14/bonjour-tout-le-monde/, /404/]
[wordpress-wpjson:routes] [http] [info] Found on 1 URLs ["50 routes"] [/]
[wordpress-xmlrpc:listmethods] [http] [info] Found on 2 URLs ["80 methods"] [/xmlrpc.php, /xmlrpc.php/xmlrpc.php]
[wordpress-xmlrpc] [http] [info] Found on 2 URLs [/xmlrpc.php]
[wordpress-xmlrpc:pingback] [http] [medium] Found on 1 URLs [/xmlrpc.php]
```

This will display the following results on the test GLPI instance.

```
                     __     _
   ____  __  _______/ /__  (_)
  / __ \/ / / / ___/ / _ \/ /
 / / / / /_/ / /__/ /  __/ /
/_/ /_/\__,_/\___/_/\___/_/ ng  v1.1.2

		github.com/oneaudit

[INF] Running nuclei with tags: [html] against 83 targets
[INF] Temporary file created: /tmp/swagger.yaml1887317488
[directory-listing] [http] [info] Found on 35 URLs [/bin/, /css/, /css/includes/, ...]
[form-detect] [http] [info] Found on 3 URLs ["/front/login.php"] [/, /Index, /index]
[form-detect] [http] [info] Found on 1 URLs ["/public/front/login.php"] [/public/]

[INF] Running nuclei with tags: [generic] against 240 targets
[INF] Temporary file created: /tmp/swagger.yaml3295598526
[cookie-detect] [http] [info] Found on 22 URLs ["glpi_8c1bab8cdfb8e857ff0e6290f00f8ed5"] [/, ///front/locale.php, //front/locale.php, ...]
[cookies-without-httponly] [javascript] [info] Found on 22 URLs ["glpi_8c1bab8cdfb8e857ff0e6290f00f8ed5"] [/status/application/json.bak]
[cookies-without-secure] [javascript] [info] Found on 22 URLs ["glpi_8c1bab8cdfb8e857ff0e6290f00f8ed5"] [//index.php]
[email-detect] [http] [info] Found on 2 URLs ["wowohoo@qq.com"] [//public/lib/fuzzy.min.js, /public/lib/fuzzy.min.js]
[favicon-detect:glpi] [http] [info] Found on 1 URLs [/pics/favicon.ico]
[http-missing-security-headers:content-security-policy] [http] [info] Found on 241 URLs [/, ///front/locale.php, //front/locale.php, ...]
[http-missing-security-headers:permissions-policy] [http] [info] Found on 241 URLs [/, ///front/locale.php, //front/locale.php, ...]
[http-missing-security-headers:referrer-policy] [http] [info] Found on 241 URLs [/, ///front/locale.php, //front/locale.php, ...]
[http-missing-security-headers:strict-transport-security] [http] [info] Found on 241 URLs [/, ///front/locale.php, //front/locale.php, ...]
[http-missing-security-headers:x-content-type-options] [http] [info] Found on 241 URLs [/, ///front/locale.php, //front/locale.php, ...]
[known-403-page:apache] [http] [info] Found on 6 URLs [/:, /js/"+CFG_GLPI.root_doc+e.original.url+", /js/"+CFG_GLPI.root_doc+el.original.url+", ...]
[known-404-page:apache] [http] [info] Found on 3 URLs [/PAGENOTFOUND, /PAGENOTFOUND/, /public/front/login.php]
[license-file] [http] [info] Found on 2 URLs [/LICENSE, /license]
[open-redirect-detect:detect] [javascript] [low] Found on 2 URLs ["domain"] [/front/locale.php]
[options-method-generic] [http] [info] Found on 208 URLs [//js/common.min.js, //js/fileupload.min.js, //js/flatpickr_buttons_plugin.min.js, ...]
[options-method-non-generic] [javascript] [info] Found on 208 URLs ["POST"] [/css/includes/components/_fileupload.scss]
[options-method-non-generic] [javascript] [info] Found on 208 URLs ["TRACE"] [/css/includes/components/_fileupload.scss]
[php-errors] [http] [info] Found on 1 URLs ["Fatal error"] [//front/locale.php]
[php-errors] [http] [info] Found on 2 URLs ["uncaught exception"] [//public/lib/base.min.js, /public/lib/base.min.js]
[readme-file] [http] [info] Found on 1 URLs [/README.md]
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

## Current issues

When access to a route is blocked, but an alternative route works, common checks are not executed on the alternative route unless the route is added to the list.

- [ ] Exploits are not executed
- [ ] Jetty (Java), XXX (Python), WordPress (MySQL), etc.
- [ ] MIA

```
[open-redirect-detect:parameters] [javascript] [info] Found on 6 URLs ["redirect"] [/////index.php]
[source-map-js] [http] [low] Found on 1 URLs [/assets/js/bundle.js.map]
"Yoast SEO:24.5"
"WordPress Site Editor"
"WordPress Block Editor"
SASS/SCSS
.idea (path exposure, etc.)
.twig
```

* [ ] Incorrect

```
[source-map-js] [http] [low] Found on 8 URLs [//public/lib/leaflet.min.js, //public/lib/photoswipe.min.js, /js/planning.js, ...]
```