# SimpleKPI Client SDK for Go

[![Build Status][build-status-svg]][build-status-link]
[![Go Report Card][goreport-svg]][goreport-link]
[![Docs][docs-godoc-svg]][docs-godoc-link]
[![License][license-svg]][license-link]

 [build-status-svg]: https://api.travis-ci.org/grokify/go-simplekpi.svg?branch=master
 [build-status-link]: https://travis-ci.org/grokify/go-simplekpi
 [goreport-svg]: https://goreportcard.com/badge/github.com/grokify/go-simplekpi
 [goreport-link]: https://goreportcard.com/report/github.com/grokify/go-simplekpi
 [docs-godoc-svg]: https://img.shields.io/badge/docs-godoc-blue.svg
 [docs-godoc-link]: https://godoc.org/github.com/grokify/go-simplekpi
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-link]: https://github.com/grokify/go-simplekpi/blob/master/LICENSE

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
