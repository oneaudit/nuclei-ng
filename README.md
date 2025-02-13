<h1 align="center">
  Nuclei-ng
</h1>
<h4 align="center">...</h4>

```console
CGO_ENABLED=1 go install github.com/oneaudit/nuclei-ng/cmd/nuclei-ng@latest
```

```

                     __     _
   ____  __  _______/ /__  (_)
  / __ \/ / / / ___/ / _ \/ /
 / / / / /_/ / /__/ /  __/ /
/_/ /_/\__,_/\___/_/\___/_/ ng  v1.0.8

		github.com/oneaudit

[INF] Running nuclei with tags: [generic] against 32 targets
[INF] Temporary file created: /tmp/swagger.yaml1498543949
[cookie-detect] [http] [info] Found on 1 URLs ["PHPSESSID"] [/comment]
[cookie-detect] [http] [info] Found on 1 URLs ["user"] [/ng_hidden_login]
[cookies-without-httponly] [javascript] [info] Found on 1 URLs ["PHPSESSID"]
[cookies-without-httponly] [javascript] [info] Found on 1 URLs ["user"]
[cookies-without-secure] [javascript] [info] Found on 1 URLs ["PHPSESSID"]
[cookies-without-secure] [javascript] [info] Found on 1 URLs ["user"]
[favicon-detect:vuejs] [http] [info] Found on 1 URLs [/favicon.ico]
[favicon-new] [http] [info] Found on 1 URLs ["b7f5b488d0b802ed63ea4ffefbbb1d6d"] [/secret.ico]
[http-missing-security-headers:content-security-policy] [http] [info] Found on 33 URLs [/, /apache-v, /comment, ...]
[http-missing-security-headers:permissions-policy] [http] [info] Found on 33 URLs [/, /apache-v, /comment, ...]
[http-missing-security-headers:referrer-policy] [http] [info] Found on 33 URLs [/, /apache-v, /comment, ...]
[http-missing-security-headers:strict-transport-security] [http] [info] Found on 33 URLs [/, /apache-v, /comment, ...]
[http-missing-security-headers:x-content-type-options] [http] [info] Found on 33 URLs [/, /apache-v, /comment, ...]
[options-method-generic] [http] [info] Found on 3 URLs [/ng_hidden_login, /ng_hidden_spy]
[options-method-non-generic] [javascript] [info] Found on 3 URLs ["POST"]
[tech-detect:apache] [http] [info] Found on 2 URLs [/apache-v, /php-v]
[tech-detect:jetty] [http] [info] Found on 1 URLs [/icons/]
[tech-detect:nginx] [http] [info] Found on 1 URLs [/nginx-v]
[tech-detect:php] [http] [info] Found on 1 URLs [/comment]
[tech-detect:python] [http] [info] Found on 2 URLs [/]
[tech-version-new] [javascript] [info] Found on 1 URLs ["Unknown/21.4.21"]
[INF] Running nuclei with tags: [html] against 27 targets
[INF] Temporary file created: /tmp/swagger.yaml3547701473
[directory-listing] [http] [info] Found on 3 URLs [/, /secret/]
[form-detect] [http] [info] Found on 1 URLs ["#"] [/simple-form]
[form-detect] [http] [info] Found on 1 URLs ["/ng_hidden_login"] [/cookie-form]
[html-comments-detect] [javascript] [info] Found on 1 URLs ["<!-- \\n\\n\\n\\nmy secret password is:\\n\\n\\n\\n toto123\\n\\n-->"]
[html-comments-detect] [javascript] [info] Found on 1 URLs ["<!-- my secret password is: toto123 -->"]
```