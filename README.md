[![Let's FIWARE Banner](https://github.com/lets-fiware/ngsi-go/blob/gh-pages/img/lets-fiware-logo-non-free.png)](https://www.letsfiware.jp/)

[![NGSI v2](https://img.shields.io/badge/NGSI-v2-5dc0cf.svg)](https://fiware-ges.github.io/orion/api/v2/stable/)
[![NGSI LD](https://img.shields.io/badge/NGSI-LD-d6604d.svg)](https://www.etsi.org/deliver/etsi_gs/CIM/001_099/009/01.03.01_60/gs_cim009v010301p.pdf)
[![License: MIT](https://img.shields.io/github/license/lets-fiware/ngsi-go.svg)](https://opensource.org/licenses/MIT)
![GitHub top language](https://img.shields.io/github/languages/top/lets-fiware/ngsi-go)
![Lines of code](https://img.shields.io/tokei/lines/github/lets-fiware/ngsi-go)
[![Build Status](https://travis-ci.com/lets-fiware/ngsi-go.svg?branch=main)](https://travis-ci.com/lets-fiware/ngsi-go)
[![Coverage Status](https://coveralls.io/repos/github/lets-fiware/ngsi-go/badge.svg?branch=main)](https://coveralls.io/github/lets-fiware/ngsi-go?branch=main)

The NGSI Go is a Unix-like command-line tool for FIWARE NGSIv2 and NGSI-LD.

## Contents
 
<details>
<summary><strong>Details</strong></summary>

-   [Getting Started with NGSI Go](#getting-started-with-ngsi-go)
-   [Usage](#usage)
-   [Install](#install)
-   [Document](#document)
-   [Third party packages](#third-party-packages)
-   [Copyright and License](#copyright-and-license)


</details>

# What's NGSI Go

> "Brave (hero), bearer of the blood of Erdrick, hero of legend! Know that your weapon will not
> serve to vanquish the Dragonload."
>
> — DRAGON WARRIOR (DRAGON QUEST)

The NGSI Go is a Unix-like command-line tool for FIWARE NGSIv2 and NGSI-LD.

## Getting Started with NGSI Go

You can get the version of your context broker instance as shown:

```json
$ ngsi version -h localhost:1026
{
"orion" : {
  "version" : "2.5.0",
  "uptime" : "0 d, 5 h, 7 m, 50 s",
  "git_hash" : "63cc107657ae10aa03f1c83bdea0be869d8e26a1",
  "compile_time" : "Fri Oct 30 09:02:37 UTC 2020",
  "compiled_by" : "root",
  "compiled_in" : "320890801dd4",
  "release_date" : "Fri Oct 30 09:02:37 UTC 2020",
  "doc" : "https://fiware-orion.rtfd.io/en/2.5.0/",
  "libversions": {
     "boost": "1_53",
     "libcurl": "libcurl/7.29.0 NSS/3.44 zlib/1.2.7 libidn/1.28 libssh2/1.8.0",
     "libmicrohttpd": "0.9.70",
     "openssl": "1.0.2k",
     "rapidjson": "1.1.0",
     "mongodriver": "legacy-1.1.2"
  }
}
}
```

You can register an alias to access the broker.

```
ngsi broker add --host letsfiware --brokerHost http://localhost:1026 --ngsiType v2
```

You can get the version by using the alias `letsfiware`.

```
$ ngsi version -h letsfiware
{
"orion" : {
  "version" : "2.5.0",
  "uptime" : "0 d, 5 h, 7 m, 50 s",
  "git_hash" : "63cc107657ae10aa03f1c83bdea0be869d8e26a1",
  "compile_time" : "Fri Oct 30 09:02:37 UTC 2020",
  "compiled_by" : "root",
  "compiled_in" : "320890801dd4",
  "release_date" : "Fri Oct 30 09:02:37 UTC 2020",
  "doc" : "https://fiware-orion.rtfd.io/en/2.5.0/",
  "libversions": {
     "boost": "1_53",
     "libcurl": "libcurl/7.29.0 NSS/3.44 zlib/1.2.7 libidn/1.28 libssh2/1.8.0",
     "libmicrohttpd": "0.9.70",
     "openssl": "1.0.2k",
     "rapidjson": "1.1.0",
     "mongodriver": "legacy-1.1.2"
  }
}
}
```

Once you access the broker, you can omit to specify the broker.

```
ngsi version
```

If you want to check the current settings, you can run the following command.

```
ngsi settings list
```

## Usage

```
NAME:
   ngsi - unix-like command-line tool for FIWARE NGSI and NGSI-LD

USAGE:
   ngsi [global options] command [command options] [arguments...]

VERSION:
   0.1.0 (git_hash:d64ff8e24b70c000561528baad2517a6ae8575d6)

COMMANDS:
   help, h  Shows a list of commands or help for one command
   CONVENIENCE:
     cp        copy entities
     wc        print number of entities, subscriptions, registrations, or types
     man       print urls of document
     ls        list entities
     rm        remove entities
     template  create template of subscription or registration
     version   print the version of Context Broker
   MANAGEMENT:
     broker    manage config for broker
     context   manage @context
     settings  manage settings
     token     manage token
   NGSI:
     append   append attributes
     create   create entity(ies), subscription or registration
     delete   delete entity(ies), attribute, subscription, or registration
     get      get entity(ies), attribute(s), subscription, registration, or type
     list     list types, entities, subscriptions, or registrations
     replace  replace entities or attributes
     update   update entities, attribute(s), or subscription
     upsert   upsert entities

GLOBAL OPTIONS:
   --syslog LEVEL  specify logging LEVEL (off, err, info, debug)
   --stderr LEVEL  specify logging LEVEL (off, err, info, debug)
   --config FILE   specify configuration FILE
   --cache FILE    specify cache FILE
   --batch, -B     don't use previous args (batch) (default: false)
   --help          show help (default: false)
   --version, -v   print the version (default: false)

COPYRIGHT:
   (c) 2020 Kazuhito Suda
```

## Install

### Install NGSI Go binary

Install NGSI Go binary in `/usr/local/bin`.

```
curl -OL https://github.com/lets-fiware/ngsi-go/releases/download/v0.1.0/ngsi-v0.1.0-linux-amd64.tar.gz
sudo tar zxvf ngsi-v0.1.0-linux-amd64.tar.gz -C /usr/local/bin
```

### Install bash autocomplete file for NGSI Go

Install ngsi_bash_autocomplete file in `/etc/bash_completion.d`.

```
curl -OL https://raw.githubusercontent.com/lets-fiware/ngsi-go/main/autocomplete/ngsi_bash_autocomplete
sudo mv ngsi_bash_autocomplete /etc/bash_completion.d/
source /etc/bash_completion.d/ngsi_bash_autocomplete
echo "source /etc/bash_completion.d/ngsi_bash_autocomplete" >> ~/.bashrc
```

### Other binaries

-    ngsi-v0.1.0-linux-arm.tar.gz
-    ngsi-v0.1.0-linux-arm64.tar.gz
-    ngsi-v0.1.0-darwin-amd64.tar.gz

## Document

-    [NGSI Go document](https://ngsi-go.letsfiware.jp/)

## Third party packages

The NGSI Go makes use of the following package:

| Package                                         | OSS License        |
| ----------------------------------------------- | ------------------ |
| [urfave/cli](https://github.com/urfave/cli)     | MIT License        |

The dependencies of dependencies have been omitted from the list.

## Copyright and License

Copyright (c) 2020 Kazuhito Suda<br>
Licensed under the MIT license.
