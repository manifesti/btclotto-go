# BTCLotto

Inspired by [saracens](https://github.com/saracen) [directory.io](http://directory.io), this script generates (random, not sequenced) private keys using BTCSuite's utilities and goquery for checking the keys for prior use.

Once it finds a key with prior transactions or any balance (check done through [blockchain.info](https://blockchain.info)), it writes the private and public key into needles.txt in the application root.

## Installing
```
go get github.com/manifesti/btclotto-go
```
