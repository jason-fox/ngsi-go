[![Let's FIWARE Banner](https://raw.githubusercontent.com/lets-fiware/ngsi-go/gh-pages/img/lets-fiware-logo-non-free.png)](https://www.letsfiware.jp/)

[![NGSI v2](https://img.shields.io/badge/NGSI-v2-5dc0cf.svg)](https://fiware-ges.github.io/orion/api/v2/stable/)
[![NGSI LD](https://img.shields.io/badge/NGSI-LD-d6604d.svg)](https://www.etsi.org/deliver/etsi_gs/CIM/001_099/009/01.03.01_60/gs_cim009v010301p.pdf)
[![License: MIT](https://img.shields.io/github/license/lets-fiware/ngsi-go.svg)](https://opensource.org/licenses/MIT)
![GitHub top language](https://img.shields.io/github/languages/top/lets-fiware/ngsi-go)
![Lines of code](https://img.shields.io/tokei/lines/github/lets-fiware/ngsi-go)
[![Build Status](https://travis-ci.com/lets-fiware/ngsi-go.svg?branch=main)](https://travis-ci.com/lets-fiware/ngsi-go)
[![Coverage Status](https://coveralls.io/repos/github/lets-fiware/ngsi-go/badge.svg?branch=main)](https://coveralls.io/github/lets-fiware/ngsi-go?branch=main)

# NGSI Go

The NGSI Go is a Unix-like command-line tool for FIWARE NGSI and NGSI-LD.

## Contents

-   [Usage](usage.md)
-   [Quick Start Guide](quick_start_guide.md)
-   [Install](install.md)
-   [Build from source](build_source.md)

## Walkthrough

-   [NGSI-LD CRUD](walkthrough/ngsi-ld-crud.md)
-   [NGSIv2 CRUD](walkthrough/ngsi-v2-crud.md)

## Command reference

### NGSI
-   [append](ngsi/append.md): append attributes
-   [create](ngsi/create.md): create entity(ies), subscription or registration
-   [delete](ngsi/delete.md): delete entity(ies), attribute, subscription or registration
-   [get](ngsi/get.md): get entity(ies) or attribute(s)
-   [list](ngsi/list.md): list types, entities, subscriptions or registrations
-   [replace](ngsi/replace.md): replace entities or attributes
-   [update](ngsi/update.md): update entities, attribute(s) or subscription
-   [upsert](ngsi/upsert.md): upsert entities

### Convenience
-   [cp](convenience/cp.md): copy entities
-   [wc](convenience/wc.md): print number of entities, subscriptions or registrations
-   [man](convenience/man.md): print  URLs of the documents related to the NGSI Go
-   [ls](convenience/ls.md): list entities
-   [rm](convenience/rm.md): remove entities
-   [template](convenience/template.md): create template of subscription or registration
-   [version](convenience/version.md): print the version of Context Broker

### Management
-    [broker](management/broker.md): manage config for broker
-    [context](management/context.md): manage @context
-    [settings](management/settings.md):  manage settings
-    [token](management/token.md): manage token

### Global Options

-   [Global Options](global.md)

### files

-   [Files](files.md)

## GitHub repository

-    [lets-fiware/ngsi-go](https://github.com/lets-fiware/ngsi-go/)

## Copyright and License

Copyright (c) 2020 Kazuhito Suda<br>
Licensed under the MIT license.
