# Testnet
Demonstrating a failed swap transaction in testnet

### When created answering to queries (see below)
![Image of created block](./stage1.png)
### eventually reverting
![Image of turning to reverted trx](./stage2.png)

### Build

```bash
make
```
### Config
Provide a runtime key arg for the const ethSenderAddr which you may change in main.go.
Contract addresses are set for testnet.

### Run Swap
```bash
./swap -ethkey 2434f
Sending HTLT with timestamp:  2021-02-17 01:23:26 +0200 EET expiration @ 2021-02-20 01:23:26 +0200 EET
Pending nonce: 66
init tx sent: 0x2394e7d1c4774a9e9dcf8ce876baa9cae6a7c5238e384de43a9e6c823bec2296
init eth tx sent: 0x2394e7d1c4774a9e9dcf8ce876baa9cae6a7c5238e384de43a9e6c823bec2296created trx hash: 0x2394e7d1c4774a9e9dcf8ce876baa9cae6a7c5238e384de43a9e6c823bec2296
sleeping 2 seconds...
query HasSwap() returns true, for swapID:114b68255e648281e979243579ff00f10f13b24e1d20e4af34a9ed9969ea3778
sender addr: 0x001d05573f994671E11d416b55cf3Cf1e5aF8dCC
receiver addr: 0x90F8bf6A479f320ead074411a4B0e7944Ea8c9C1
swap id: 114b68255e648281e979243579ff00f10f13b24e1d20e4af34a9ed9969ea3778
random number hash: db0fa6b8f020a972a90453f6c87296993b32a05e3f92b1f712a363b9cf64c886
addr: 9d108807294d526c1fcd66d0fcefab1c402753f5
timestamp: 1613406818
expire time: 1613666018
erc20 amount: 1
amount: 0
getSwap() retuns -> swap obj:
swap.Id 0x114b68255e648281e979243579ff00f10f13b24e1d20e4af34a9ed9969ea3778
swap.RandomNumberHash 0xdb0fa6b8f020a972a90453f6c87296993b32a05e3f92b1f712a363b9cf64c886
swap.ExpireTimestamp 1613666018
swap.OutAmount 1
swap.RecipientAddress 0x90F8bf6A479f320ead074411a4B0e7944Ea8c9C1
swap.RecipientOtherChain 0x9d108807294D526C1fCd66d0FCEFAb1C402753F5
swap.SenderAddress 0x001d05573f994671E11d416b55cf3Cf1e5aF8dCC
```