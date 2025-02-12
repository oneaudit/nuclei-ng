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
/_/ /_/\__,_/\___/_/\___/_/ ng  v1.0.7

		github.com/oneaudit

[INF] Temporary file created: /tmp/swagger.yaml286455905
[cookie-detect] [http] [info] Found on 1 URLs ["PHPSESSID"] [/comment]
[cookies-without-httponly] [javascript] [info] Found on 1 URLs ["PHPSESSID"]
[cookies-without-secure] [javascript] [info] Found on 1 URLs ["PHPSESSID"]
[html-comments-detect] [javascript] [info] Found on 1 URLs ["<!-- \\n\\n\\n\\nmy secret password is:\\n\\n\\n\\n toto123\\n\\n-->"]
[html-comments-detect] [javascript] [info] Found on 1 URLs ["<!-- my secret password is: toto123 -->"]
[http-missing-security-headers:content-security-policy] [http] [info] Found on 28 URLs [/, /apache-v, /comment, ...]
[http-missing-security-headers:permissions-policy] [http] [info] Found on 28 URLs [/, /apache-v, /comment, ...]
[http-missing-security-headers:referrer-policy] [http] [info] Found on 28 URLs [/, /apache-v, /comment, ...]
[http-missing-security-headers:strict-transport-security] [http] [info] Found on 28 URLs [/, /apache-v, /comment, ...]
[http-missing-security-headers:x-content-type-options] [http] [info] Found on 28 URLs [/, /apache-v, /comment, ...]
[options-method-generic] [http] [info] Found on 3 URLs [/ng_hidden_login, /ng_hidden_spy]
[options-method-non-generic] [javascript] [info] Found on 3 URLs ["POST"]
[tech-detect:apache] [http] [info] Found on 2 URLs [/apache-v, /php-v]
[tech-detect:jetty] [http] [info] Found on 1 URLs [/icons/]
[tech-detect:nginx] [http] [info] Found on 1 URLs [/nginx-v]
[tech-detect:php] [http] [info] Found on 1 URLs [/comment]
[tech-detect:python] [http] [info] Found on 2 URLs [/]
[tech-version-new] [javascript] [info] Found on 1 URLs ["Unknown/21.4.21"]
[tech-version:nginx] [http] [info] Found on 1 URLs ["1.33.7"] [/nginx-v]
[tech-version:werkzeug] [http] [info] Found on 2 URLs ["1.5.7"] [/]
[tech-version:jetty] [http] [info] Found on 1 URLs ["12.0.17.v20201231"] [/icons/]
[tech-version:apache] [http] [info] Found on 2 URLs ["2.4.41"] [/apache-v, /php-v]
[tech-version:python] [http] [info] Found on 2 URLs ["3.10.2"] [/]
[tech-version:php] [http] [info] Found on 2 URLs ["7.4.0"] [/comment, /php-v]
```