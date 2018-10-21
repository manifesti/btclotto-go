# BTCLotto

Inspired by [saracens](https://github.com/saracen) [directory.io](http://directory.io), this script generates (random, not sequenced) private keys using BTCSuite's utilities and go-electrum to check the
given keys against nodes.

Each node is given it's own goroutine to check the keys with, the more nodes you bother the quicker it should work, in theory. Addresses are given in the electrum.txt file, every row contains one node-address in the format address:port.

### Example Output
```
address 16xxhmdtX9JNhHSKThTmqSHFcpR6MsM6s5 has no history..
address 1CJ3ohcpNf71VFRKUEURstXvv3yJfCUvmy has no history..
address 1HcztYNxu78KdMU5dafU9fbmPDij43zGio has no history..
address 1LSpNmNDzEjkSVMkQMrmWAy3UQfrRy8ESA has no history..
address 1KGYd3WnfhHPheffJN5cw78WB2Zvj2eB4Q has no history..
address 18puh9U1bjPBRkV6pLo8iwxpBsmhnuU2yn has no history..
address 12NJQtaH2NYfjpf81mrNovvNkKnZHfuoQm has no history..
address 19xDh8H5knH9qC7QHkmV8ZzwfWgrHN1fsc has no history..
address 14EvnwYuiAyCYhDQSKCWgKrvfgcVpj7QHn has no history..
address 1AaiaWR13Qiy96EyfoXdVyVE5GGnpDwBMi has no history..
address 1MTzmttet4yxJJmEAgHSkbXJgYsx4gcx4x has no history..
address 1KsamVRSR7NzBMLEc5MauTBdAEGCKMoPf2 has no history..
address 1BAdC1oU48K3F1ADVhudiZ4cte6L6WHDLy has no history..
address 17aJX1GPhXt2sFRoF1qqu2csozvHwqqYQH has no history..
address 1FFCVSS8D3h2QND1RPWm5BHPUDZrAMvgvJ has no history..
```
