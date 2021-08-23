# sso2aws [![Go](https://github.com/PennState/sso2aws/actions/workflows/go.yml/badge.svg)](https://github.com/PennState/sso2aws/actions/workflows/go.yml)

A simple CLI wrapper, like a stripped down [saml2aws](https://github.com/Versent/saml2aws), to manage multiple AWS SSO profiles through config files and ENV vars.

## Usage

```
A tool to make operating under multiple AWS SSO accounts and roles easier

Usage:
  ./sso2aws [command]

Available Commands:
  completion        generate the autocompletion script for the specified shell
  exec              execute commands with the configured profile
  get-config        show runtime config
  help              Help about any command
  login             login
  update-aws-config write aws profile to shared config
  version

Flags:
  -d, --debug     debug output
  -h, --help      help for ./sso2aws
  -v, --version   version for ./sso2aws
```

## Configuration

Config can be set via YAML (~/.sso2aws.yaml) and overriden via environment variable (eg with direnv).

|cfg name|envvar|default|usage|
|---|---|---|---|
|sso_config.region|SSO2AWS_REGION|us-east-1|AWS Region|
|sso_config.sso_account_id|SSO2AWS_SSO_ACCOUNT_ID|-|AWS Account ID|
|sso_config.sso_start_url|SSO2AWS_SSO_START_URL|-|AWS SSO Start URL|
|sso_config.sso_role_name|SSO2AWS_SSO_ROLE_NAME|AdministratorAccess|AWS role name|
|sso_config.sso_region|SSO2AWS_SSO_REGION|us-east-1|AWS SSO Region|
|profile|SSO2AWS_PROFILE|sso2aws|AWS config profile name to use|
