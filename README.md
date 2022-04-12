# SimpleKPI Client SDK for Go

[![Build Status][build-status-svg]][build-status-url]
[![Go Report Card][goreport-svg]][goreport-url]
[![Docs][docs-godoc-svg]][docs-godoc-url]
[![License][license-svg]][license-url]

 [build-status-svg]: https://github.com/grokify/go-simplekpi/workflows/test/badge.svg
 [build-status-url]: https://github.com/grokify/go-simplekpi/actions
 [goreport-svg]: https://goreportcard.com/badge/github.com/grokify/go-simplekpi
 [goreport-url]: https://goreportcard.com/report/github.com/grokify/go-simplekpi
 [docs-godoc-svg]: https://pkg.go.dev/badge/github.com/grokify/go-simplekpi
 [docs-godoc-url]: https://pkg.go.dev/github.com/grokify/go-simplekpi
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-url]: https://github.com/grokify/go-simplekpi/blob/master/LICENSE

This is a SimpleKPI.com SDK for Go built using OpenAPI Generator.

The API is documented here:

* https://support.simplekpi.com/Developers

The generated SDK is in the [`simplekpi`](simplekpi) folder.

## Usage

See the [`examples`](examples) folder for usage.

## API Coverage

- [ ] KPIs
  - [x] /api/kpis GET
- [x] KPI Entries
  - [x] /api/kpientries GET
  - [x] /api/kpientries POST
  - [x] /api/kpientries/{id} GET
  - [x] /api/kpientries/{id} PUT
  - [x] /api/kpientries/{id} DELETE
- [ ] Users
  - [x] /api/users GET
