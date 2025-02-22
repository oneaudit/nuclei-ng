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

[INF] Running nuclei with tags: [html] against 32 targets
[INF] Temporary file created: /tmp/swagger.yaml1269739532
[directory-listing] [http] [info] Found on 3 URLs [/, /secret/]
[form-detect] [http] [info] Found on 1 URLs ["#"] [/simple-form]
[form-detect] [http] [info] Found on 1 URLs ["/ng_hidden_login"] [/cookie-form]
[html-comments-detect] [javascript] [info] Found on 1 URLs ["<!--<!-- @version (+redirect) -->-->"] [/libs]
[html-comments-detect] [javascript] [info] Found on 1 URLs ["<!--<!-- \\n\\n\\n\\nmy secret password is:\\n\\n\\n\\n toto123\\n\\n-->-->"] [/comment-long]
[html-comments-detect] [javascript] [info] Found on 1 URLs ["<!--<!-- local -->-->"] [/libs]
[html-comments-detect] [javascript] [info] Found on 1 URLs ["<!--<!-- my secret password is: toto123 -->-->"] [/comment]
[html-comments-detect] [javascript] [info] Found on 1 URLs ["<!--<!-- no version -->-->"] [/libs]
[html-comments-detect] [javascript] [info] Found on 1 URLs ["<!--<!-- version in copyright -->-->"] [/libs]
[html-comments-detect] [javascript] [info] Found on 1 URLs ["<!--<!-- version in name -->-->"] [/libs]
[html-comments-detect] [javascript] [info] Found on 1 URLs ["<!--<!-- version in path -->-->"] [/libs]

[INF] Running nuclei with tags: [generic] against 43 targets
[INF] Temporary file created: /tmp/swagger.yaml747212024
[cookie-detect] [http] [info] Found on 1 URLs ["PHPSESSID"] [/comment]
[cookie-detect] [http] [info] Found on 1 URLs ["user"] [/ng_hidden_login]
[cookies-without-httponly] [javascript] [info] Found on 1 URLs ["PHPSESSID"] [/comment]
[cookies-without-httponly] [javascript] [info] Found on 1 URLs ["user"] [/ng_hidden_login]
[cookies-without-secure] [javascript] [info] Found on 1 URLs ["PHPSESSID"] [/comment]
[cookies-without-secure] [javascript] [info] Found on 1 URLs ["user"] [/ng_hidden_login]
[cors-detect:headers] [http] [info] Found on 1 URLs [/cors]
[cors-exploit:arbitrary-origin] [http] [info] Found on 1 URLs [header:Origin] [GET] [/cors]
[email-detect] [http] [info] Found on 1 URLs ["stuart [at] stuartk.com"] [/assets/js/jszip.js]
[favicon-detect:Vue] [http] [info] Found on 1 URLs [/favicon.ico]
[favicon-new:mmh3] [http] [info] Found on 1 URLs ["1823185746"] [/secret.ico]
[favicon-new:md5] [http] [info] Found on 1 URLs ["b7f5b488d0b802ed63ea4ffefbbb1d6d"] [/secret.ico]
[http-missing-security-headers:content-security-policy] [http] [info] Found on 44 URLs [/, /apache-v, /assets/, ...]
[http-missing-security-headers:permissions-policy] [http] [info] Found on 45 URLs [/, /apache-v, /assets/, ...]
[http-missing-security-headers:referrer-policy] [http] [info] Found on 45 URLs [/, /apache-v, /assets/, ...]
[http-missing-security-headers:strict-transport-security] [http] [info] Found on 45 URLs [/, /apache-v, /assets/, ...]
[http-missing-security-headers:x-content-type-options] [http] [info] Found on 45 URLs [/, /apache-v, /assets/, ...]
[http-suspicious-request-headers] [javascript] [info] Found on 1 URLs ["X-Api-Key: MYS3cr374P1K3y"] [/ng_hidden_spy]
[open-redirect-detect:parameters] [javascript] [info] Found on 1 URLs ["redirect"]
[open-redirect-exploit] [http] [medium] Found on 1 URLs [query:redirect] [GET] [/]
[options-method-generic] [http] [info] Found on 3 URLs [/ng_hidden_login, /ng_hidden_spy]
[options-method-non-generic] [javascript] [info] Found on 3 URLs ["POST"] [undefined]
[tech-detect:apache] [http] [info] Found on 2 URLs [/apache-v, /php-v]
[tech-detect:jetty] [http] [info] Found on 1 URLs [/icons/]
[tech-detect:jsdelivr] [http] [info] Found on 1 URLs [/libs]
[tech-detect:nginx] [http] [info] Found on 1 URLs [/nginx-v]
[tech-detect:php] [http] [info] Found on 1 URLs [/comment]
[tech-detect:python] [http] [info] Found on 2 URLs [/]
[tech-version-new] [javascript] [info] Found on 1 URLs ["Unknown/21.4.21"] [undefined]
[tech-version-new] [javascript] [info] Found on 1 URLs ["empty_page/1234"] [undefined]
[tech-version:nginx] [http] [info] Found on 1 URLs ["1.33.7"] [/nginx-v]
[tech-version:werkzeug] [http] [info] Found on 2 URLs ["1.5.7"] [/]
[tech-version:jetty] [http] [info] Found on 1 URLs ["12.0.17.v20201231"] [/icons/]
[tech-version:apache] [http] [info] Found on 2 URLs ["2.4.41"] [/apache-v, /php-v]
[tech-version:python] [http] [info] Found on 2 URLs ["3.10.2"] [/]
[tech-version:php] [http] [info] Found on 2 URLs ["7.4.0"] [/comment, /php-v]

[INF] Running nuclei with tags: [javascript] against 9 targets
[INF] Temporary file created: /tmp/swagger.yaml3600235712
[javascript-library] [javascript] [info] Found on 1 URLs ["angular.js==1.8.3"] [/assets/js/angular.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["bootstrap.js==5.3.3"] [/assets/js/bootstrap.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["jquery-1.14.0.js==1.14.0"] [/assets/js/jquery-1.14.0.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["jquery-migrate.js==3.5.2"] [/assets/js/jquery-migrate.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["jquery-ui.js==1.14.1"] [/assets/js/jquery-ui.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["jquery.js==3.7.1"] [/assets/js/jquery.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["jszip.js==3.10.1"] [/assets/js/jszip.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["leaflet.js==unknown"] [/assets/js/leaflet.js]
[javascript-library] [javascript] [info] Found on 1 URLs ["link_id.js==unknown"] [/link_id.js]

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

[INF] Running nuclei with tags: [html] against 64 targets
[INF] Temporary file created: /tmp/swagger.yaml2790815807
[form-detect] [http] [info] Found on 2 URLs ["http://wp.sec2/"] [//, /comments/]
[form-detect] [http] [info] Found on 3 URLs ["http://wp.sec2/wp-comments-post.php"] [/2025/02/14/bonjour-tout-le-monde/]
[form-detect] [http] [info] Found on 3 URLs ["http://wp.sec2/wp-login.php"] [//wp-login, /wp-login, /wp-login.php]
[html-comments-detect] [javascript] [info] Found on 11 URLs ["<!--<!-- / Yoast SEO plugin. -->-->"] [/category/non-classe/]
[html-comments-detect] [javascript] [info] Found on 11 URLs ["<!--<!-- This site is optimized with the Yoast SEO plugin v24.5 - https://yoast.com/wordpress/plugins/seo/ -->-->"] [/category/non-classe/]
[html-comments-detect] [javascript] [info] Found on 4 URLs ["<!--<!-- This site is optimized with the Yoast SEO plugin v24.5 - https:\\/\\/yoast.com\\/wordpress\\/plugins\\/seo\\/ -->-->"] [/wp-json/wp/v2/users]
[html-comments-detect] [javascript] [info] Found on 4 URLs ["<!--<!-- \\/ Yoast SEO plugin. -->-->"] [/wp-json/wp/v2/users]
[html-comments-detect] [javascript] [info] Found on 1 URLs ["<!--<!-- generator="WordPress/6.7.2" -->-->"] [/wp-links-opml]

[INF] Running nuclei with tags: [generic] against 89 targets
[INF] Temporary file created: /tmp/swagger.yaml4026525154
[cookie-detect] [http] [info] Found on 7 URLs ["wordpress_c9fd5a0c97d0e084fdf80980af45bd8b"] [/////////wp-login.php, ///////wp-login.php, //////wp-login.php, ...]
[cookie-detect] [http] [info] Found on 7 URLs ["wordpress_logged_in_c9fd5a0c97d0e084fdf80980af45bd8b"] [/////////wp-login.php, ///////wp-login.php, //////wp-login.php, ...]
[cookie-detect] [http] [info] Found on 7 URLs ["wordpress_sec_c9fd5a0c97d0e084fdf80980af45bd8b"] [/////////wp-login.php, ///////wp-login.php, //////wp-login.php, ...]
[cookie-detect] [http] [info] Found on 16 URLs ["wordpress_test_cookie"] [/////////wp-login.php, ////////wp-login.php, ///////wp-login.php, ...]
[cookie-detect] [http] [info] Found on 7 URLs ["wordpresspass_c9fd5a0c97d0e084fdf80980af45bd8b"] [/////////wp-login.php, ///////wp-login.php, //////wp-login.php, ...]
[cookie-detect] [http] [info] Found on 7 URLs ["wordpressuser_c9fd5a0c97d0e084fdf80980af45bd8b"] [/////////wp-login.php, ///////wp-login.php, //////wp-login.php, ...]
[cookie-detect] [http] [info] Found on 7 URLs ["wp-postpass_c9fd5a0c97d0e084fdf80980af45bd8b"] [/////////wp-login.php, ///////wp-login.php, //////wp-login.php, ...]
[cookie-detect] [http] [info] Found on 7 URLs ["wp-settings-0"] [/////////wp-login.php, ///////wp-login.php, //////wp-login.php, ...]
[cookie-detect] [http] [info] Found on 7 URLs ["wp-settings-time-0"] [/////////wp-login.php, ///////wp-login.php, //////wp-login.php, ...]
[cookie-detect] [http] [info] Found on 4 URLs ["wp_lang"] [////wp-login.php, ///wp-login.php, //wp-login, //wp-login.php]
[cookies-without-httponly] [javascript] [info] Found on 4 URLs ["wordpress_c9fd5a0c97d0e084fdf80980af45bd8b"] [/////////wp-login.php]
[cookies-without-httponly] [javascript] [info] Found on 4 URLs ["wordpress_logged_in_c9fd5a0c97d0e084fdf80980af45bd8b"] [/////////wp-login.php]
[cookies-without-httponly] [javascript] [info] Found on 4 URLs ["wordpress_sec_c9fd5a0c97d0e084fdf80980af45bd8b"] [/////////wp-login.php]
[cookies-without-httponly] [javascript] [info] Found on 13 URLs ["wordpress_test_cookie"] [////wp-login.php]
[cookies-without-httponly] [javascript] [info] Found on 4 URLs ["wordpresspass_c9fd5a0c97d0e084fdf80980af45bd8b"] [/////////wp-login.php]
[cookies-without-httponly] [javascript] [info] Found on 4 URLs ["wordpressuser_c9fd5a0c97d0e084fdf80980af45bd8b"] [/////////wp-login.php]
[cookies-without-httponly] [javascript] [info] Found on 4 URLs ["wp-postpass_c9fd5a0c97d0e084fdf80980af45bd8b"] [/////////wp-login.php]
[cookies-without-httponly] [javascript] [info] Found on 4 URLs ["wp-settings-0"] [/////////wp-login.php]
[cookies-without-httponly] [javascript] [info] Found on 4 URLs ["wp-settings-time-0"] [/////////wp-login.php]
[cookies-without-httponly] [javascript] [info] Found on 4 URLs ["wp_lang"] [////wp-login.php]
[cookies-without-secure] [javascript] [info] Found on 4 URLs ["wordpress_c9fd5a0c97d0e084fdf80980af45bd8b"] [//////wp-login.php]
[cookies-without-secure] [javascript] [info] Found on 4 URLs ["wordpress_logged_in_c9fd5a0c97d0e084fdf80980af45bd8b"] [//////wp-login.php]
[cookies-without-secure] [javascript] [info] Found on 4 URLs ["wordpress_sec_c9fd5a0c97d0e084fdf80980af45bd8b"] [//////wp-login.php]
[cookies-without-secure] [javascript] [info] Found on 13 URLs ["wordpress_test_cookie"] [/wp-login.php]
[cookies-without-secure] [javascript] [info] Found on 4 URLs ["wordpresspass_c9fd5a0c97d0e084fdf80980af45bd8b"] [//////wp-login.php]
[cookies-without-secure] [javascript] [info] Found on 4 URLs ["wordpressuser_c9fd5a0c97d0e084fdf80980af45bd8b"] [//////wp-login.php]
[cookies-without-secure] [javascript] [info] Found on 4 URLs ["wp-postpass_c9fd5a0c97d0e084fdf80980af45bd8b"] [//////wp-login.php]
[cookies-without-secure] [javascript] [info] Found on 4 URLs ["wp-settings-0"] [//////wp-login.php]
[cookies-without-secure] [javascript] [info] Found on 4 URLs ["wp-settings-time-0"] [//////wp-login.php]
[cookies-without-secure] [javascript] [info] Found on 4 URLs ["wp_lang"] [//wp-login]
[cors-detect:headers] [http] [info] Found on 17 URLs [////wp-json/oembed/1.0/embed, ///wp-json/oembed/1.0/embed, //wp-json/oembed/1.0/embed, ...]
[cors-exploit:arbitrary-origin] [http] [info] Found on 16 URLs [header:Origin] [GET] [////wp-json/oembed/1.0/embed, ///wp-json/oembed/1.0/embed, //wp-json/oembed/1.0/embed, ...]
[http-missing-security-headers:content-security-policy] [http] [info] Found on 115 URLs [/, /.git/, /.git/HEAD, ...]
[http-missing-security-headers:permissions-policy] [http] [info] Found on 115 URLs [/, /.git/, /.git/HEAD, ...]
[http-missing-security-headers:referrer-policy] [http] [info] Found on 115 URLs [/, /.git/, /.git/HEAD, ...]
[http-missing-security-headers:strict-transport-security] [http] [info] Found on 115 URLs [/, /.git/, /.git/HEAD, ...]
[http-missing-security-headers:x-content-type-options] [http] [info] Found on 99 URLs [/, /.git/, /.git/HEAD, ...]
[open-redirect-detect:parameters] [javascript] [info] Found on 1 URLs ["dir"]
[open-redirect-detect:parameters] [javascript] [info] Found on 3 URLs ["redirect_to"]
[open-redirect-detect:parameters] [javascript] [info] Found on 4 URLs ["url"]
[options-method-generic] [http] [info] Found on 54 URLs [/.git/, /.git/HEAD, /.git/config, ...]
[options-method-non-generic] [javascript] [info] Found on 40 URLs ["POST"] [undefined]
[options-method-non-generic] [javascript] [info] Found on 36 URLs ["TRACE"] [undefined]
[tech-detect:apache] [http] [info] Found on 122 URLs [/, /.git/, /.git/HEAD, ...]
[tech-detect:gravatar] [http] [info] Found on 4 URLs [/2025/02/14/bonjour-tout-le-monde/, /author/testwp/]
[tech-detect:php] [http] [info] Found on 86 URLs [/, //, ///, ...]
[tech-version:mod_fcgid] [http] [info] Found on 122 URLs ["2.3.10"] [/, /.git/, /.git/HEAD, ...]
[tech-version:apache] [http] [info] Found on 122 URLs ["2.4.59"] [/, /.git/, /.git/HEAD, ...]
[tech-version:php] [http] [info] Found on 122 URLs ["8.3.6"] [/, /.git/, /.git/HEAD, ...]
[tech-version:cms] [http] [info] Found on 15 URLs ["WordPress 6.7.2"] [/, //, /0/, ...]

[INF] Running nuclei with tags: [javascript] against 4 targets
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
[javascript-library] [code] [info] Found on 9 URLs ["view.min.js==unknown"] [/, /0/, /2025/, ...]
[javascript-library] [code] [info] Found on 4 URLs ["wp-util.min.js==unknown"] [/wp-admin/import.php, /wp-admin/update-core.php, /wp-login, /wp-login.php]
```