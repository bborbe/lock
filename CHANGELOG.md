# Changelog

All notable changes to this project will be documented in this file.

Please choose versions by [Semantic Versioning](http://semver.org/).

* MAJOR version when you make incompatible API changes,
* MINOR version when you add functionality in a backwards-compatible manner, and
* PATCH version when you make backwards-compatible bug fixes.

## Unreleased

- fix: bump `osv-scanner` to v2.6.0 and `golang.org/x/net` to v0.60.0 so the Linux vulnerability gates stop failing. v2.3.1 pins `golang.org/x/tools` v0.38.0, whose SSA builder aborts with `unexpected expr: *ast.KeyValueExpr` on the promoted-field composite-literal key Go 1.27 permits in the Linux stdlib, so a repo on the old pin passes locally on darwin and fails only in Linux CI. `x/net` v0.58.0 carries `GO-2026-6603/6610/6611/6612/6617`, which fail both `vulncheck` and `trivy`.

## v1.0.6

- chore: update Go to 1.27.1

## v1.0.5

- chore: update Go to 1.27.0

## v1.0.4

- chore: Bump errcheck to v1.20.0 and golangci-lint to v2.13.1 for Go 1.27 support
## v1.0.3

- update Go to 1.26.6 and update dependencies

## v1.0.2

- align repo to current lib standard: drop vendor + tools.go, add CI, maintainer pipeline, lint + vuln-scan configs
- update Go to 1.26.5 and update dependencies

## v1.0.1

- go mod update

## v1.0.0

- Initial Version
