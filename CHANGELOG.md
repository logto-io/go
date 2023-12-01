# Changelog

## [v1.0.4](https://github.com/logto-io/go/compare/v1.0.3...v1.0.4) (2023-12-01)

### Feat

* support organizations ([#89](https://github.com/logto-io/go/issues/89))
* **core:** add organization data to userinfo response ([#91](https://github.com/logto-io/go/issues/91))

### Fix

* **deps:** update module github.com/jarcoal/httpmock to v1.3.1 ([#80](https://github.com/logto-io/go/issues/80))
* **deps:** update golang.org/x/exp digest to 6522937 ([#79](https://github.com/logto-io/go/issues/79))
* **deps:** update module github.com/stretchr/testify to v1.8.4 ([#73](https://github.com/logto-io/go/issues/73))
* **deps:** update module github.com/agiledragon/gomonkey/v2 to v2.10.1 ([#69](https://github.com/logto-io/go/issues/69))
* **deps:** update golang.org/x/exp digest to 2e198f4 ([#68](https://github.com/logto-io/go/issues/68))


## [v1.0.3](https://github.com/logto-io/go/compare/1.0.3...v1.0.3) (2023-05-29)


## [1.0.3](https://github.com/logto-io/go/compare/v1.0.2...1.0.3) (2023-05-29)

### Fix

* ensure the basic auth header is included when fetching token ([#75](https://github.com/logto-io/go/issues/75))


## [v1.0.2](https://github.com/logto-io/go/compare/v1.0.1...v1.0.2) (2023-05-18)

### Fix

* fetch user info ([#71](https://github.com/logto-io/go/issues/71))
* **client:** remove id token for sign out ([#67](https://github.com/logto-io/go/issues/67))
* **deps:** update module github.com/jarcoal/httpmock to v1.3.0 ([#60](https://github.com/logto-io/go/issues/60))
* **deps:** update module github.com/gin-gonic/gin to v1.9.0 ([#59](https://github.com/logto-io/go/issues/59))
* **deps:** update module github.com/stretchr/testify to v1.8.2 ([#52](https://github.com/logto-io/go/issues/52))
* **deps:** update module github.com/agiledragon/gomonkey/v2 to v2.9.0 ([#50](https://github.com/logto-io/go/issues/50))
* **deps:** update golang.org/x/exp digest to 10a5072 ([#48](https://github.com/logto-io/go/issues/48))


## [v1.0.1](https://github.com/logto-io/go/compare/v1.0.0...v1.0.1) (2023-04-11)

### Fix

* **client:** add resource when calling FetchTokenByRefreshToken ([#64](https://github.com/logto-io/go/issues/64))


## [v1.0.0](https://github.com/logto-io/go/compare/v1.0.0-rc.0...v1.0.0) (2023-03-22)

### Fix

* **docs:** document link issue ([#61](https://github.com/logto-io/go/issues/61))

### Refactor

* **core:** replace `idTokenHint` with `clientId` in sign-out url ([#57](https://github.com/logto-io/go/issues/57))
* **core:** remove deprecated role names ([#54](https://github.com/logto-io/go/issues/54))


## [v1.0.0-rc.0](https://github.com/logto-io/go/compare/v0.1.5...v1.0.0-rc.0) (2023-02-07)

### Feat

* **client:** remove scope for access token grant ([#53](https://github.com/logto-io/go/issues/53))


## [v0.1.5](https://github.com/logto-io/go/compare/v0.1.4...v0.1.5) (2022-09-27)

### Feat

* fetch user info ([#51](https://github.com/logto-io/go/issues/51))


## [v0.1.4](https://github.com/logto-io/go/compare/v0.1.3...v0.1.4) (2022-09-09)


## [v0.1.3](https://github.com/logto-io/go/compare/v0.1.2...v0.1.3) (2022-09-09)

### Feat

* **client:** persist access tokens all the time ([#45](https://github.com/logto-io/go/issues/45))

### Fix

* **deps:** update golang.org/x/exp digest to 145caa8 ([#41](https://github.com/logto-io/go/issues/41))

### Refactor

* **gin-sample:** rename storage -> `session_storage` ([#24](https://github.com/logto-io/go/issues/24))


## [v0.1.2](https://github.com/logto-io/go/compare/v0.1.1...v0.1.2) (2022-09-07)

### Fix

* **client:** skip resource check in `GetAccessToken` when resource is empty string ([#38](https://github.com/logto-io/go/issues/38))
* **client:** `GetIdTokenClaims` should return `ErrNotAuthenticated` if the user is not authenticated ([#37](https://github.com/logto-io/go/issues/37))


## v0.1.1 (2022-09-06)

### Feat

* web sample ([#13](https://github.com/logto-io/go/issues/13))
* add authenticated user related method ([#12](https://github.com/logto-io/go/issues/12))
* logto client ([#10](https://github.com/logto-io/go/issues/10))
* **client:** handle signing out ([#11](https://github.com/logto-io/go/issues/11))
* **core:** `verifyAndParseCodeFromCallbackUri` ([#5](https://github.com/logto-io/go/issues/5))
* **core:** core functions to interact with OIDC APIs ([#7](https://github.com/logto-io/go/issues/7))
* **core:** generate sign in & sign out uri ([#6](https://github.com/logto-io/go/issues/6))
* **core:** token utils ([#4](https://github.com/logto-io/go/issues/4))
* **core:** generator functions ([#3](https://github.com/logto-io/go/issues/3))
* **core:** types ([#2](https://github.com/logto-io/go/issues/2))

### Fix

* **deps:** update golang.org/x/exp digest to 334a238 ([#23](https://github.com/logto-io/go/issues/23))

### Refactor

* use `assert` in test cases ([#35](https://github.com/logto-io/go/issues/35))
* use packages in project ([#19](https://github.com/logto-io/go/issues/19))
* rename go modules ([#14](https://github.com/logto-io/go/issues/14))
* **client:** rename `SignInContext` -> `SignInSession` ([#36](https://github.com/logto-io/go/issues/36))
* **client:** replace sign-in context key with constant value ([#32](https://github.com/logto-io/go/issues/32))
* **client:** extract errors ([#33](https://github.com/logto-io/go/issues/33))
* **client:** extract storage key ([#27](https://github.com/logto-io/go/issues/27))

