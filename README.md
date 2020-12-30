# `RedisWrapper` Intro

![RedisWrapper Package Logo https://gitlab.com/evatix-go/rediswrapper.git](/uploads/23b06de55397346b05314842415e12f8/250.png)

Rediswrapper is wrapper and helper for using redis database in memory.

## Git Clone

`git clone https://gitlab.com/evatix-go/rediswrapper.git`

### 2FA enabled, for linux

`git clone https://[YourGitLabUserName]:[YourGitlabAcessTokenGenerateFromGitlabsTokens]@gitlab.com/evatix-go/rediswrapper.git`

### Prerequisites

- Update git to latest 2.29
- Update or install the latest of Go 1.15.2
- Either add your ssh key to your gitlab account
- Or, use your access token to clone it.

## Installation

`go get gitlab.com/evatix-go/rediswrapper`

Install [redis in your machine following issues.](https://gitlab.com/evatix-go/rediswrapper/-/issues/13)

### Go get issue for private package

- Update git to 2.29
- Enable go modules. (Windows : `go env -w GO111MODULE=on`, Unix : `export GO111MODULE=on`)
- Add `gitlab.com/evatix-go` to go env private
  
To set for Windows:

`go env -w GOPRIVATE=[AddExistingOnes;]gitlab.com/evatix-go`

To set for Unix:

`expoort GOPRIVATE=[AddExistingOnes;]gitlab.com/evatix-go`

## Why `YourModuleName?`

## Examples

`Code Smaples`

## Acknowledgement

Any other packages used

## Links

## Issues

- [Create your issues](https://gitlab.com/evatix-go/rediswrapper/-/issues)

## Notes

## Contributors

## License

[Evatix MIT License](/LICENSE)
