# Pingdom CLI

Written in Go and built with [mitchellh/cli](https://github.com/mitchellh/cli)
this library provides a CLI to some of [Pingdom's API](https://www.pingdom.com/resources/api).

At this time it has a single subcommand, `uptime`, which can be used to fetch
uptime for all or some checks for a given period.

Uptime is reported as a percentage.

## Requirements
Authenticating with Pingdom's API requires your Pingdom username, password,
and an [App-Key](https://www.pingdom.com/resources/api#authentication).
This are provided using environment variables:

 - `PINGDOM_USER`
 - `PINGDOM_PASS`
 - `PINGDOM_APPKEY`

## Installation
If you are a Go developer you can install and build locally by running:

`go get github.com/fillup/pingdom`

Or, you can download the latest binary for your operating system from the
[dist/](https://github.com/fillup/pingdom-cli/tree/master/dist/) directory of this
repo.

## Usage

### uptime
Get a report of uptime by Check

Optional Flags:

 - `-period`: Specify a relative time period, example: `-period LastMonth`.
Default is `Today`.
 - `-tags`: Comma-separated list of tags. Only checks that have at least one
of these tags will be returned. Example: `-tags tag1,tag2`.
If not specified, all Checks will be returned

Example:

```
$ pingdom uptime -period LastMonth -tags appsdev

Pingdom Uptime Report
Reporting Period (LastMonth):
From: 2017-07-01 00:00:00 +0000 UTC
To:   2017-07-31 23:59:59 +0000 UTC

Check                    Uptime
Developer Portal         99.980%
Addressbook              100.000%
Search Assistant         100.000%
View IPC                 99.982%
Doorman UI               100.000%
Doorman API              99.998%
...
```
