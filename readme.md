# Readme
- This repo demo for `how to build & use pkg` in project & cmd.

## About `./pkg/`
- Build pkg rule

## About `./cmd/`
- Implement pkg

## About `./go.mod` & `./go.sum`

1. We could see the `pkg path` in the 'go.sum' & 'go.mod'.
    - In general, go.mod format is `<RepoPath> v1.0.0`
    - The version has a difference with the regular version. Because regular version repository, itself is a package.
2. see `https://github.com/jhaoheng/gobuildpkgdemo2` and check its go.mod