# AWS Refresh Access Key

A small go util for rotating your AWS Secret Access Key from the comfort of your CLI.

It uses your curent aws access key to issue a new key and rotate your key in `~/.aws/credentials`

Requires the following aws IAM permissions on your own user.

 * `iam:CreateAccessKey`
 * `iam:DeleteAccessKey`


## Usage

**Basic Usage**
```bash
$ aws-refresh-access-keys
New Access Key Created: AKIASMJZFUZETLGX3VGP
Deleting old access key: AKIASMJZFUZE5RAUPZM4
Saving Access Key (AKIASMJZFUZETLGX3VGP) to /Users/simon/.aws/credentials. (profile = [default])
```

**Supports [named profiles](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-profiles.html) for AWS CLI**
```bash
$ AWS_PROFILE="prod" aws-refresh-access-keys
# [...]
Saving Access Key (AKIASMJZFUZETLGX3VGP) to /Users/simon/.aws/credentials. (profile = [prod])
```

**Intended to work with [aws-mfa](https://github.com/broamski/aws-mfa)**
```bash
$ AWS_PROFILE="prod" aws-mfa
# [...]
$ AWS_PROFILE="prod" aws-refresh-access-keys
New Access Key Created: AKIASMJZFUZETLGX3VGP
Deleting old access key: AKIASMJZFUZE5RAUPZM4
Saving Access Key (AKIASMJZFUZETLGX3VGP) to /Users/simon/.aws/credentials. (profile = [prod-long-term])
```

## Building the Binary

```bash
cd src
go build
```
