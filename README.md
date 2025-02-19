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
nuclei-ng -config config.yaml -i oppa_localhost_5000.yaml
```

This will display the following results.

```
                     __     _
   ____  __  _______/ /__  (_)
  / __ \/ / / / ___/ / _ \/ /
 / / / / /_/ / /__/ /  __/ /
/_/ /_/\__,_/\___/_/\___/_/ ng  v1.0.10

		github.com/oneaudit

[INF] Running nuclei with tags: html against 27 targets
[INF] Temporary file created: /tmp/swagger.yaml1269739532
[directory-listing] [http] [info] Found on 3 URLs [/, /secret/]
[form-detect] [http] [info] Found on 1 URLs ["#"] [/simple-form]
[form-detect] [http] [info] Found on 1 URLs ["/ng_hidden_login"] [/cookie-form]
[html-comments-detect] [javascript] [info] Found on 1 URLs ["<!-- \\n\\n\\n\\nmy secret password is:\\n\\n\\n\\n toto123\\n\\n-->"]
[html-comments-detect] [javascript] [info] Found on 1 URLs ["<!-- my secret password is: toto123 -->"]

[INF] Running nuclei with tags: [generic] against 31 targets
[INF] Temporary file created: /tmp/swagger.yaml747212024
[cookie-detect] [http] [info] Found on 2 URLs ["PHPSESSID"] [/comment]
[cookie-detect] [http] [info] Found on 2 URLs ["user"] [/ng_hidden_login]
[cookies-without-httponly] [javascript] [info] Found on 1 URLs ["PHPSESSID"]
[cookies-without-httponly] [javascript] [info] Found on 1 URLs ["user"]
[cookies-without-secure] [javascript] [info] Found on 1 URLs ["PHPSESSID"]
[cookies-without-secure] [javascript] [info] Found on 1 URLs ["user"]
[cors-detect:headers] [http] [info] Found on 2 URLs [/cors]
[cors-exploit:arbitrary-origin] [http] [info] Found on 1 URLs [header:Origin] [GET] [/cors]
[favicon-detect:Vue] [http] [info] Found on 1 URLs [/favicon.ico]
[favicon-new:mmh3] [http] [info] Found on 2 URLs ["1823185746"] [/secret.ico]
[favicon-new:md5] [http] [info] Found on 2 URLs ["b7f5b488d0b802ed63ea4ffefbbb1d6d"] [/secret.ico]
[http-missing-security-headers:content-security-policy] [http] [info] Found on 63 URLs [/, /apache-v, ...]
[http-missing-security-headers:permissions-policy] [http] [info] Found on 65 URLs [/, /apache-v, ...]
[http-missing-security-headers:referrer-policy] [http] [info] Found on 65 URLs [/, /apache-v, ...]
[http-missing-security-headers:strict-transport-security] [http] [info] Found on 65 URLs [/, /apache-v, ...]
[http-missing-security-headers:x-content-type-options] [http] [info] Found on 65 URLs [/, /apache-v, ...]
[http-suspicious-request-headers] [javascript] [info] Found on 1 URLs ["X-Api-Key: MYS3cr374P1K3y"]
[http-suspicious-response-headers] [javascript] [info] Found on 1 URLs ["X-Entrypoint: /empty_page/1234/"]
[open-redirect-detect:parameters] [javascript] [info] Found on 1 URLs ["redirect"]
[open-redirect-exploit] [http] [medium] Found on 1 URLs [query:redirect] [GET] [/]
[options-method-generic] [http] [info] Found on 3 URLs [/ng_hidden_login, /ng_hidden_spy]
[options-method-non-generic] [javascript] [info] Found on 3 URLs ["POST"]
[tech-detect:apache] [http] [info] Found on 4 URLs [/apache-v, /php-v]
[tech-detect:bootstrap] [http] [info] Found on 2 URLs [/libs]
[tech-detect:bulma] [http] [info] Found on 2 URLs [/libs]
[tech-detect:font-awesome] [http] [info] Found on 2 URLs [/libs]
[tech-detect:jetty] [http] [info] Found on 2 URLs [/icons/]
[tech-detect:jsdelivr] [http] [info] Found on 2 URLs [/libs]
[tech-detect:materialize-css] [http] [info] Found on 2 URLs [/libs]
[tech-detect:milligram] [http] [info] Found on 2 URLs [/libs]
[tech-detect:nginx] [http] [info] Found on 2 URLs [/nginx-v]
[tech-detect:php] [http] [info] Found on 2 URLs [/comment]
[tech-detect:python] [http] [info] Found on 4 URLs [/]
[tech-detect:semantic-ui] [http] [info] Found on 2 URLs [/libs]
[tech-detect:zurb-foundation] [http] [info] Found on 2 URLs [/libs]
[tech-version-new] [javascript] [info] Found on 1 URLs ["Unknown/21.4.21"]
[tech-version-new] [javascript] [info] Found on 1 URLs ["empty_page/1234"]
[tech-version:nginx] [http] [info] Found on 2 URLs ["1.33.7"] [/nginx-v]
[tech-version:werkzeug] [http] [info] Found on 4 URLs ["1.5.7"] [/]
[tech-version:jetty] [http] [info] Found on 2 URLs ["12.0.17.v20201231"] [/icons/]
[tech-version:apache] [http] [info] Found on 4 URLs ["2.4.41"] [/apache-v, /php-v]
[tech-version:python] [http] [info] Found on 4 URLs ["3.10.2"] [/]
[tech-version:php] [http] [info] Found on 4 URLs ["7.4.0"] [/comment, /php-v]
```