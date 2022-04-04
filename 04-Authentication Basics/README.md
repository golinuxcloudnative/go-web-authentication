# Authentication vs Authorization

| Authentication                            | Authorization             |
| ----------------------------------------- | ------------------------- |
| Who you are                               | Which permission you have |
| Passwords, one-time pins, biometric, card | Rule, policie, scope      |
| First step                                | After authentication      |
| ID tokens                                 | access token              |
| OpenID Connection (OIDC)                  | OAuth2                    |

# HTTP Basic Authentication

| :exclamation: Only use this method over HTTPS/TLS |
| ------------------------------------------------- |

[RFC7617 - The 'Basic' HTTP Authentication Scheme](https://datatracker.ietf.org/doc/html/rfc7617)

**Header -** `Authorization`  
**Format -** `Authorization: Basic userid:password`  
**usedid:password -** must convert to base64 (easily reversible)

Curl example  `curl -u user:passwd -v google.com`  


| ![This is an image](/99-utils/imgs/http-auth-sequence-diagram.png) |
| :----------------------------------------------------------------: |
|                     Authentication flow[^1].                      |

[^1]: https://developer.mozilla.org/en-US/docs/Web/HTTP/Authentication



# Simple App

```bash
# 1) Unauthorized -> no Authentication header
curl -v localhost:8000

# 2) Unauthorized -> Authentication header OK, but wrong user or password
curl -v -u admin:passwd  localhost:8000

# 3) Authenticated -> Authentication header OK and user and password OK
curl -v -u admin:admin  localhost:8000

# 4) Authenticated -> Authentication header OK and user and password OK
# You need to generate a base64 from "admin:admin"
curl -v localhost:8000 -H "Authorization: Basic YWRtaW46YWRtaW4="
```