# BTCLotto

Inspired by [saracens](https://github.com/saracen) [directory.io](http://directory.io), this script generates (random, not sequenced) private keys using BTCSuite's utilities and go-electrum to check the
given keys against nodes.

Each node is given it's own goroutine to check the keys with, the more nodes you bother the quicker it should work, in theory. Addresses are given in the electrum.txt file, every row contains one node-address in the format address:port.
