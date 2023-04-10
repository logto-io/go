# Changelog

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

### Chore

* update changelog ([#46](https://github.com/logto-io/go/issues/46))


## [v0.1.3](https://github.com/logto-io/go/compare/v0.1.2...v0.1.3) (2022-09-09)

### Feat

* **client:** persist access tokens all the time ([#45](https://github.com/logto-io/go/issues/45))

### Fix

* **deps:** update golang.org/x/exp digest to 145caa8 ([#41](https://github.com/logto-io/go/issues/41))

### Refactor

* **gin-sample:** rename storage -> `session_storage` ([#24](https://github.com/logto-io/go/issues/24))


## [v0.1.2](https://github.com/logto-io/go/compare/v0.1.1...v0.1.2) (2022-09-07)

### Chore

* add readme ([#39](https://github.com/logto-io/go/issues/39))

### Fix

* **client:** skip resource check in `GetAccessToken` when resource is empty string ([#38](https://github.com/logto-io/go/issues/38))
* **client:** `GetIdTokenClaims` should return `ErrNotAuthenticated` if the user is not authenticated ([#37](https://github.com/logto-io/go/issues/37))


## v0.1.1 (2022-09-06)

### Chore

* bump dependencies version ([#16](https://github.com/logto-io/go/issues/16))
* bump core version to v0.1.0 (client, gin-sample) ([#15](https://github.com/logto-io/go/issues/15))
* add license ([#8](https://github.com/logto-io/go/issues/8))
* project setup ([#1](https://github.com/logto-io/go/issues/1))
* add readme
* **deps:** add renovate.json ([#20](https://github.com/logto-io/go/issues/20))

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

