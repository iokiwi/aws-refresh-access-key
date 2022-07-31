# AWS Refresh Access Key

A small go util for rotating your AWS Secret Access Key from the comfort of your CLI

## Usage

```bash
$ aws-refresh-access-keys
New Access Key Created: AKIASMJZFUZETLGX3VGP
Deleting old access key: AKIASMJZFUZE5RAUPZM4
Saving Access Key (AKIASMJZFUZETLGX3VGP) to /Users/simon/.aws/credentials. (profile = [default])
```

## Building

```bash
cd src
go build
```
