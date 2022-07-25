# helmet
 HTTP security middleware for Go(lang) inspired by HelmetJS.

## Features

| Module | Default | Option Function | Docs |
| --- | --- | --- | --- |
| ``Content-Security-Policy[-Report-Only]`` | ``default-src 'self'; object-src 'none'; child-src https:;`` | ``WithContentSecurityPolicyReportOnly(bool)`` <br> ``WithContentSecurityPolicyDirectives(string)`` | [MDN](https://developer.mozilla.org/en-US/docs/Web/HTTP/CSP) |
| ``Cross-Origin-Embedder-Policy`` | - | ``WithCrossOriginEmbedderPolicy(enum)`` | - |
| ``Cross-Origin-Opener-Policy`` | - | ``WithCrossOriginOpenerPolicy(enum)`` | - |
| ``Cross-Origin-Resource-Policy`` | - | ``WithCrossOriginResourcePolicy(enum)`` | - |
| ``Expect-CT`` | ``max-age:0`` | ``WithExpectCTMaxAge(time.Duration)`` <br> ``WithExpectCTEnforce(bool)`` <br> ``WithExpectCTReportUri(string)`` | [MDN](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Expect-CT) |
| ``Origin-Agent-Cluster`` | ``?1`` | - | - |
| ``Referrer-Policy`` | ``no-referrer`` | ``WithReferrerPolicyDirectives([]enum)`` | [MDN](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Referrer-Policy) |
| ``Strict-Transport-Security`` | ``max-age:<180 days>`` | ``WithStrictTransportSecurityMaxAge(time.Duration)`` <br> ``WithStrictTransportSecurityIncludeSubDomains(bool)`` <br> ``WithStrictTransportSecurityPreload(bool)`` | - |
| ``X-Content-Type-Options`` | ``nosniff`` | - | [MDN](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Content-Type-Options) |
| ``X-DNS-Prefetch-Control`` | ``off`` | ``WithXDNSPrefetchControl(enum)`` | [MDN](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-DNS-Prefetch-Control) |
| ``X-Download-Options`` | ``noopen`` | - | - |
| ``X-Frame-Options`` | ``SAMEORIGIN`` | ``WithXFrameOptions(enum)`` | [MDN](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Frame-Options) |
| ``X-Permitted-Cross-Domain-Policies`` | ``none`` | ``WithXPermittedCrossDomainPolicies(enum)`` | - |
| ``X-Powered-By`` | - | ``WithHideXPoweredBy(bool)`` <br> ``WithReplacementXPoweredBy(string)`` | - |
| ``X-XSS-Protection`` | ``0`` | - | [MDN](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-XSS-Protection) |

## Contrib
| Engine | Version |
| --- | --- |
| [gin](https://github.com/gin-gonic/gin) | ``v1.8.1`` |
| [iris](https://github.com/kataras/iris) | ``v12.2.0-beta4`` |
