[![Build Status](https://travis-ci.org/abtris/bee.svg?branch=master)](https://travis-ci.org/abtris/bee)

# Bee

Bee is a command line tool for publishing and fetching existing documents from Apiary.io.

## Install

Need golang 1.11+ - deps using go modules

```
go get github.com/abtris/bee
```

or binary downloads:

### Mac

```
wget https://github.com/abtris/bee/releases/download/1.0.0/bee-darwin-amd64 -O /usr/local/bin/bee
chmod +x /usr/local/bin/bee
```

### Linux

```
wget https://github.com/abtris/bee/releases/download/1.0.0/bee-linux-amd64 -O /usr/local/bin/bee
chmod +x /usr/local/bin/bee
```

### Windows

Just [download file](https://github.com/abtris/bee/releases/download/1.0.0/bee-windows-amd64) and make executable (add `.exe`).

## Config File

In order to fetch and publish documentation, you will need to generate an authentication token by going to [https://login.apiary.io/tokens](https://login.apiary.io/tokens). Keep this token safe, because a token is like a password and will give anyone with this token access to fetch and publish to your documentation.

`~/.bee.yaml`

```yaml
APIARY_API_KEY: XXXXXXXXXXXXXX
```

## Usage

```
bee fetch <subdomain>
```

or

```
bee publish <subdomain> filename
```

## Debug - using proxy as [Charles](https://www.charlesproxy.com/)

HTTP_PROXY=http://127.0.0.1:8888 bee fetch

# Release


You need use [goreleaser](https://goreleaser.com/). For work need GITHUB_TOKEN with repo scope and installed goreleaser.

```
export GITHUB_TOKEN=XXXX
git tag -a v1.1.1 -m "Release 1.1.1"
git push origin v1.1.1
goreleaser --rm-dist
```

## License

MIT
