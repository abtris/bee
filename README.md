# Bee

Bee is a command line tool for publishing and fetching existing documents from Apiary.io.

## Install

```
go get github.com/abtris/bee
```

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


# TODO

- [] tests
- [] docs