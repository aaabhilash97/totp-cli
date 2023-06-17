# Simple TOTP CLI

## Add new TOTP

totp -add -account "zero-tele-otp" -value "otpauth://totp/<host>?algorithm=SHA1&digits=6&issuer=<issuer>&period=30&secret=<secret>"

## Copy new TOTP PassCode to clipboard

totp -account "zero-tele-otp"

## Delete TOTP

totp -delete -account "zero-tele-otp"