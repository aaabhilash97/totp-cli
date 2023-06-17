# Simple TOTP CLI

## Add new TOTP

```sh
totp -add -account "zero-tele-otp" -url "otpauth://totp/<host>?algorithm=SHA1&digits=6&issuer=<issuer>&period=30&secret=<secret>"
```

## Copy new TOTP PassCode to clipboard

```sh
totp -account "zero-tele-otp"
```

## Delete TOTP

```sh
totp -delete -account "zero-tele-otp"
```