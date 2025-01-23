[![Moov Banner Logo](https://user-images.githubusercontent.com/20115216/104214617-885b3c80-53ec-11eb-8ce0-9fc745fb5bfc.png)](https://github.com/moov-io)

<p align="center">
  <a href="https://slack.moov.io/">Community</a>
  ·
  <a href="https://moov.io/blog/">Blog</a>
  <br>
  <br>
</p>

[![GoDoc](https://godoc.org/github.com/moov-io/iso3166?status.svg)](https://godoc.org/github.com/moov-io/iso3166)
[![Build Status](https://github.com/moov-io/iso3166/workflows/Go/badge.svg)](https://github.com/moov-io/iso3166/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/moov-io/iso3166)](https://goreportcard.com/report/github.com/moov-io/iso3166)
[![Repo Size](https://img.shields.io/github/languages/code-size/moov-io/iso3166?label=project%20size)](https://github.com/moov-io/iso3166)
[![Apache 2 License](https://img.shields.io/badge/license-Apache2-blue.svg)](https://raw.githubusercontent.com/moov-io/iso3166/master/LICENSE)
[![Slack Channel](https://slack.moov.io/badge.svg?bg=e01563&fgColor=fffff)](https://slack.moov.io/)
[![GitHub Stars](https://img.shields.io/github/stars/moov-io/iso3166)](https://github.com/moov-io/iso3166)
[![Twitter](https://img.shields.io/twitter/follow/moov?style=social)](https://twitter.com/moov?lang=en)

# moov-io/iso3166

The `iso3166` package implements ISO 3166-1-alpha2 code and name lookup.

### Usage

**Validate Country Codes**

```
if iso3166.Valid("US") {
    // do something
}
```

**Lookup Country Code from Name**

```
countryCode := iso3166.LookupCode("United States") // "US"
```

## Getting help

 channel | info
 ------- | -------
Twitter [@moov](https://twitter.com/moov)	| You can follow Moov.io's Twitter feed to get updates on our project(s). You can also tweet us questions or just share blogs or stories.
[GitHub Issue](https://github.com/moov-io/iso3166/issues/new) | If you are able to reproduce a problem please open a GitHub Issue under the specific project that caused the error.
[moov-io slack](https://slack.moov.io/) | Join our slack channel to have an interactive discussion about the development of the project.

## Supported and tested platforms

- 64-bit Linux (Ubuntu, Debian), macOS, and Windows

## License

Apache License 2.0 - See [LICENSE](LICENSE) for details.
