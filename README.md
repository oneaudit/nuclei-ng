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
/_/ /_/\__,_/\___/_/\___/_/ ng  v1.0.11

		github.com/oneaudit

[INF] Running nuclei with tags: [html] against 39 targets
[INF] Temporary file created: /tmp/swagger.yaml1269739532
[directory-listing] [http] [info] Found on 10 URLs [/, /secret/]
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

[INF] Running nuclei with tags: [generic] against 51 targets
[INF] Temporary file created: /tmp/swagger.yaml747212024
[cookie-detect] [http] [info] Found on 1 URLs ["PHPSESSID"] [/comment]
[cookie-detect] [http] [info] Found on 1 URLs ["user"] [/ng_hidden_login]
[cookies-without-httponly] [javascript] [info] Found on 1 URLs ["PHPSESSID"] [/comment]
[cookies-without-httponly] [javascript] [info] Found on 1 URLs ["user"] [/ng_hidden_login]
[cookies-without-secure] [javascript] [info] Found on 1 URLs ["PHPSESSID"] [/comment]
[cookies-without-secure] [javascript] [info] Found on 1 URLs ["user"] [/ng_hidden_login]
[cors-detect:headers] [http] [info] Found on 1 URLs [/cors]
[cors-exploit:arbitrary-origin] [http] [info] Found on 1 URLs [header:Origin] [GET] [/cors]
[email-detect] [http] [info] Found on 1 URLs ["jdoe [at] example.com"] [/humans.txt]
[email-detect] [http] [info] Found on 1 URLs ["security@example.com"] [/.well-known/security.txt]
[email-detect] [http] [info] Found on 1 URLs ["security[at]example.com"] [/.well-known/security.txt]
[email-detect] [http] [info] Found on 1 URLs ["stuart [at] stuartk.com"] [/assets/js/jszip.js]
[favicon-detect:Vue] [http] [info] Found on 1 URLs [/favicon.ico]
[favicon-new:md5] [http] [info] Found on 1 URLs ["b7f5b488d0b802ed63ea4ffefbbb1d6d"] [/secret.ico]
[favicon-new:mmh3] [http] [info] Found on 1 URLs ["1823185746"] [/secret.ico]
[http-missing-security-headers:content-security-policy] [http] [info] Found on 59 URLs [/, /.well-known/, /.well-known/security.txt, ...]
[http-missing-security-headers:permissions-policy] [http] [info] Found on 60 URLs [/, /.well-known/, /.well-known/security.txt, ...]
[http-missing-security-headers:referrer-policy] [http] [info] Found on 60 URLs [/, /.well-known/, /.well-known/security.txt, ...]
[http-missing-security-headers:strict-transport-security] [http] [info] Found on 60 URLs [/, /.well-known/, /.well-known/security.txt, ...]
[http-missing-security-headers:x-content-type-options] [http] [info] Found on 51 URLs [/, /.well-known/, /.well-known/security.txt, ...]
[http-suspicious-request-headers] [javascript] [info] Found on 1 URLs ["X-Api-Key: MYS3cr374P1K3y"] [/ng_hidden_spy]
[http-suspicious-response-headers] [javascript] [info] Found on 1 URLs ["X-Entrypoint: /empty_page/1234/"] [/js-event-link]
[open-redirect-detect:parameters] [javascript] [info] Found on 1 URLs ["redirect"] [/]
[open-redirect-exploit] [http] [medium] Found on 1 URLs [query:redirect] [GET] [/]
[options-method-generic] [http] [info] Found on 3 URLs [/ng_hidden_login, /ng_hidden_spy]
[options-method-non-generic] [javascript] [info] Found on 3 URLs ["POST"] [/ng_hidden_spy]
[robots-txt-generic] [http] [info] Found on 1 URLs [/robots.txt]
[robots-txt-non-generic] [http] [info] Found on 1 URLs ["b77a70c632c0be0cea6bbce3dff8a2317392303b0730327589b3485c4b892dcd"] [/icons/robots.txt]
[sitemap-detect] [http] [info] Found on 1 URLs [/sitemap.xml]
[tech-detect:apache] [http] [info] Found on 2 URLs [/apache-v, /php-v]
[tech-detect:jetty] [http] [info] Found on 5 URLs [/icons/, /icons/.robots.txt.swp, /icons/robots.bak, ...]
[tech-detect:jsdelivr] [http] [info] Found on 1 URLs [/libs]
[tech-detect:nginx] [http] [info] Found on 1 URLs [/nginx-v]
[tech-detect:php] [http] [info] Found on 1 URLs [/comment]
[tech-detect:python] [http] [info] Found on 6 URLs [/]
[tech-version-new] [javascript] [info] Found on 1 URLs ["Unknown/21.4.21"] [/icons/robots.txt]
[tech-version-new] [javascript] [info] Found on 1 URLs ["empty_page/1234"] [/js-event-link]
[tech-version:apache] [http] [info] Found on 2 URLs ["2.4.41"] [/apache-v, /php-v]
[tech-version:jetty] [http] [info] Found on 5 URLs ["12.0.17.v20201231"] [/icons/, /icons/.robots.txt.swp, /icons/robots.bak, ...]
[tech-version:nginx] [http] [info] Found on 1 URLs ["1.33.7"] [/nginx-v]
[tech-version:php] [http] [info] Found on 2 URLs ["7.4.0"] [/comment, /php-v]
[tech-version:python] [http] [info] Found on 6 URLs ["3.10.2"] [/]
[tech-version:werkzeug] [http] [info] Found on 6 URLs ["1.5.7"] [/]
[well-known-detect] [http] [info] Found on 5 URLs [/.well-known/security.txt, /humans.txt, /icons/robots.txt, ...]

[INF] Running nuclei with tags: [javascript] against 9 targets
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
[source-map-js] [http] [low] Found on 1 URLs [/assets/js/bundle.js.map]
[source-map-js] [http] [low] Found on 1 URLs ["bundle.js.map"] [/assets/js/bundle.js]

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
/_/ /_/\__,_/\___/_/\___/_/ ng  v1.0.11

		github.com/oneaudit

[INF] Running nuclei with tags: [html] against 72 targets
[INF] Temporary file created: /tmp/swagger.yaml2790815807
[form-detect] [http] [info] Found on 14 URLs ["http://wp.sec2/"] [//, ///, ////, ...]
[form-detect] [http] [info] Found on 2 URLs ["http://wp.sec2/wp-comments-post.php"] [/2025/02/14/bonjour-tout-le-monde/]
[html-comments-detect] [javascript] [info] Found on 1 URLs ["<!-- #respond -->"] [/2025/02/14/bonjour-tout-le-monde/]
[html-comments-detect] [javascript] [info] Found on 16 URLs ["<!-- / Yoast SEO plugin. -->"] [/docs/]
[html-comments-detect] [javascript] [info] Found on 16 URLs ["<!-- This site is optimized with the Yoast SEO plugin v24.5 - https://yoast.com/wordpress/plugins/seo/ -->"] [/docs/]
[html-comments-detect] [javascript] [info] Found on 1 URLs ["<!-- generator="WordPress/6.7.2" -->"] [/wp-links-opml]

[INF] Running nuclei with tags: [generic] against 103 targets
[INF] Temporary file created: /tmp/swagger.yaml4026525154
[email-detect] [http] [info] Found on 2 URLs ["jdoe@example.com"] [/.git/config, /.git/logs/HEAD]
[git-config:config] [http] [medium] Found on 1 URLs [/.git/HEAD]
[git-config:folder] [http] [medium] Found on 1 URLs [/.git/]
[git-config:ignore] [http] [medium] Found on 1 URLs [/.gitignore]
[git-config:logs] [http] [medium] Found on 1 URLs [/.git/logs/HEAD]
[http-missing-security-headers:content-security-policy] [http] [info] Found on 68 URLs [/.git/HEAD, /.git/config, /.git/index, ...]
[http-missing-security-headers:permissions-policy] [http] [info] Found on 68 URLs [/.git/HEAD, /.git/config, /.git/index, ...]
[http-missing-security-headers:referrer-policy] [http] [info] Found on 68 URLs [/.git/HEAD, /.git/config, /.git/index, ...]
[http-missing-security-headers:strict-transport-security] [http] [info] Found on 68 URLs [/.git/HEAD, /.git/config, /.git/index, ...]
[http-missing-security-headers:x-content-type-options] [http] [info] Found on 68 URLs [/.git/HEAD, /.git/config, /.git/index, ...]
[open-redirect-detect:parameters] [javascript] [info] Found on 2 URLs ["url"] [/wp-json/oembed/1.0/embed]
[options-method-generic] [http] [info] Found on 31 URLs [/.git/HEAD, /.git/config, /.git/logs/, ...]
[options-method-non-generic] [javascript] [info] Found on 35 URLs ["POST"] [/readme.html]
[options-method-non-generic] [javascript] [info] Found on 33 URLs ["TRACE"] [/readme.html]
[robots-txt-non-generic] [http] [info] Found on 1 URLs ["673fe3c92956f2618d0d11e0a9f0be1023e8a63ca3fbe977816ced034e767534"] [/robots.txt]
[tech-detect:apache] [http] [info] Found on 64 URLs [/.git/, /.git/HEAD, /.git/config, ...]
[tech-detect:gravatar] [http] [info] Found on 1 URLs [/2025/02/14/bonjour-tout-le-monde/]
[tech-detect:php] [http] [info] Found on 31 URLs [////, /////, //xmlrpc.php, ...]
[tech-version:apache] [http] [info] Found on 68 URLs ["2.4.59"] [/, /.git/, /.git/HEAD, ...]
[tech-version:cms] [http] [info] Found on 13 URLs ["WordPress 6.7.2"] [/, ///, ////, ...]
[tech-version:mod_fcgid] [http] [info] Found on 68 URLs ["2.3.10"] [/, /.git/, /.git/HEAD, ...]
[tech-version:php] [http] [info] Found on 68 URLs ["8.3.6"] [/, /.git/, /.git/HEAD, ...]
[well-known-detect] [http] [info] Found on 5 URLs [/LICENSE, /LICENSE.txt, /license, ...]

[INF] Running nuclei with tags: [javascript] against 6 targets
[INF] Temporary file created: /tmp/swagger.yaml259532497
[javascript-library] [javascript] [info] Found on 1 URLs ["password-strength-meter.js==unknown"] [/wp-admin/js/password-strength-meter.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["password-strength-meter.min.js==6.7.2"] [/wp-admin/js/password-strength-meter.min.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["password-strength-meter.min.js==unknown"] [//wp-admin/js/password-strength-meter.min.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["user-profile.js==unknown"] [/wp-admin/js/user-profile.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["user-profile.min.js==6.7.2"] [/wp-admin/js/user-profile.min.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["user-profile.min.js==unknown"] [//wp-admin/js/user-profile.min.js]

[INF] Running nuclei with tags: [jsext] against 10 targets
[javascript-library] [code] [info] Found on 4 URLs ["a11y.min.js==unknown"] [/wp-admin/import.php, /wp-admin/update-core.php, /wp-login, /wp-login.php]
[javascript-library] [code] [info] Found on 1 URLs ["comment-reply.min.js==unknown"] [/2025/02/14/bonjour-tout-le-monde/]
[javascript-library] [code] [info] Found on 4 URLs ["dom-ready.min.js==unknown"] [/wp-admin/import.php, /wp-admin/update-core.php, /wp-login, /wp-login.php]
[javascript-library] [code] [info] Found on 4 URLs ["i18n.min.js==unknown"] [/wp-admin/import.php, /wp-admin/update-core.php, /wp-login, /wp-login.php]
[javascript-library] [code] [info] Found on 4 URLs ["load-scripts.php==unknown"] [/wp-admin/import.php, /wp-admin/update-core.php, /wp-login, /wp-login.php]
[javascript-library] [code] [info] Found on 4 URLs ["password-strength-meter.min.js==unknown"] [/wp-admin/import.php, /wp-admin/update-core.php, /wp-login, /wp-login.php]
[javascript-library] [code] [info] Found on 4 URLs ["underscore.min.js==unknown"] [/wp-admin/import.php, /wp-admin/update-core.php, /wp-login, /wp-login.php]
[javascript-library] [code] [info] Found on 4 URLs ["user-profile.min.js==unknown"] [/wp-admin/import.php, /wp-admin/update-core.php, /wp-login, /wp-login.php]
[javascript-library] [code] [info] Found on 8 URLs ["view.min.js==unknown"] [/0/, /2025/, /2025/02/, ...]
[javascript-library] [code] [info] Found on 4 URLs ["wp-util.min.js==unknown"] [/wp-admin/import.php, /wp-admin/update-core.php, /wp-login, /wp-login.php]

[INF] Running nuclei with tags: [wordpress] against 103 targets
[INF] Temporary file created: /tmp/swagger.yaml869050629
[wordpress-debug-log] [http] [low] Found on 1 URLs ["1494043 bytes"] [/wp-content/debug.log]
[wordpress-debug-log] [http] [low] Found on 1 URLs ["1496111 bytes"] [//wp-content/debug.log]
[wordpress-license] [http] [info] Found on 4 URLs [/LICENSE, /LICENSE.txt, /license, /license.txt]
[wordpress-login] [http] [info] Found on 1 URLs [/mysecretlogin/]
[wordpress-readme] [http] [info] Found on 4 URLs [/README, /Readme, /readme, /readme.html]
[wordpress-repair] [http] [low] Found on 1 URLs [//wp-admin/maint/repair.php]
[wordpress-themes-detect] [http] [info] Found on 10 URLs ["twentytwentyfive"] [/wp-content/themes/twentytwentyfive/, /wp-content/themes/twentytwentyfive/assets/, /wp-content/themes/twentytwentyfive/assets/fonts/, ...]
[wordpress-users:author] [http] [info] Found on 1 URLs ["testwp"] [/]
[wordpress-users:rdf] [http] [info] Found on 9 URLs ["testwp"] [//rdf/, /2025/02/14/bonjour-tout-le-monde/feed/rdf/, /comments/feed/rdf/, ...]
[wordpress-users:usernames] [http] [low] Found on 1 URLs ["testwp"] [/]
[wordpress-users:yoast] [http] [info] Found on 2 URLs ["testwp"] [//author-sitemap.xml, /author-sitemap.xml]
[wordpress-version:by_css] [http] [info] Found on 28 URLs ["6.7.2"] [/, //, ///, ...]
[wordpress-version:by_generator] [http] [info] Found on 32 URLs ["6.7.2"] [/, //, ///, ...]
[wordpress-version:by_js] [http] [info] Found on 26 URLs ["6.7.2"] [/, //, ///, ...]
[wordpress-wpjson:routes] [http] [info] Found on 1 URLs ["50 routes"] [/]
[wordpress-xmlrpc:listmethods] [http] [info] Found on 2 URLs ["80 methods"] [//xmlrpc.php/xmlrpc.php, /xmlrpc/xmlrpc.php]
[wordpress-xmlrpc] [http] [info] Found on 3 URLs [//xmlrpc.php, /xmlrpc, /xmlrpc.php]
[wordpress-xmlrpc:pingback] [http] [medium] Found on 2 URLs [//xmlrpc.php, /xmlrpc]
```

## Current issues

When access to a route is blocked, but an alternative route works, common checks are not executed on the alternative route unless the route is added to the list.

- [ ] `wp-json` and `?rest_route`
- [ ] `wp-login`/`wp-admin` and CVE bypasses