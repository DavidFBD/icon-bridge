# BTP (Block Transmission Protocol) Relay System

## Introduction

We need to build a usable [BTP](doc/btp.md) Relay System which can deliver digital tokens between multiple chains.

Target chains
* ICON (goloop)
* Polkadot parachain

Terminologies

| Word            | Description                                                                                            |
|:----------------|:-------------------------------------------------------------------------------------------------------|
| BTP             | Blockchain Transmission Protocol, [ICON BTP Standard](https://github.com/icon-project/IIPs/blob/master/IIPS/iip-25.md) defined by ICON. |
| BTP Message     | A verified message which is delivered by the relay                                                     |
| Service Message | A payload in a BTP message                                                                             |
| Relay Message   | A message including BTPMessages with proofs for that, and other block update messages.                 |
| NetworkAddress  | Network Type and Network ID <br/> *0x1.icon* <br/> *0x1.iconee*                                        |
| ContractAddress | Addressing contract in the network <br/> *btp://0x1.iconee/cx87ed9048b594b95199f326fc76e76a9d33dd665b* |

> BTP Standard

### Components

* [BTP Message Verifier(BMV)](doc/bmc.md) - smart contract
  - Update blockchain verification information
  - Verify delivered BTP message and decode it

* [BTP Message Center(BMC)](doc/bmv.md) - smart contract
  - Receive BTP messages through transactions.
  - Send BTP messages through events.

* [BTP Service Handler(BSH)](doc/bsh.md) - smart contract
  - Handle service messages related to the service.
  - Send service messages through the BMC

* [BTP Message Relay(BMR)](doc/bmr.md) - external software
  - Monitor BTP events
  - Gather proofs for the events
  - Send BTP Relay Message

### Blockchain specifics
* [ICON](doc/icon.md)

## BTP Project

### Documents

* [Build Guide](doc/build.md)
* [Tutorial](doc/tutorial.md)
* [btpsimple command line](doc/btpsimple_cli.md)

### Layout

| Directory                | Description  |
|:--------------------|:-------|
| /cmd           |   Root of implement of BMR |
| /cmd/btpsimple           |   Reference implement of BMR. only provide unidirectional relay. (golang) |
| /cmd/btpsimple/chain    |   Implement of common logic of BMR, use module |
| /cmd/btpsimple/module    |   BMR module interface and common codes |
| /cmd/btpsimple/module/`<src>`    | Implement of BMR module (`Sender`,`Receiver`), `<src>` is name of source blockchain |
| /common | Common codes (golang) |
| /doc | Documents |
| /docker | Docker related resources |
| /`<env>` | Root of implement of BTP smart contracts, `<env>` is name of smart contract environment |
| /`<env>`/bmc | Implement of BMC |
| /`<env>`/bmv | Root of implement of BMV |
| /`<env>`/bmv/`<src>` | Implement of BMV, `<src>` is name of source blockchain |
| /`<env>`/`<svc>` | Root of implement of BSH, `<svc>` is name of BTP service |
| /`<env>`/token_bsh | Reference implement of BSH for Interchain-Token transfer service |
| /`<env>`/token_bsh/sample/irc2_token | Implement of IRC-2.0, example for support legacy smart contract |

### BTP implement for ICON blockchain

| Directory                | Description  |
|:--------------------|:-------|
| /cmd/btpsimple/module/iconee    | Module for ICON blockchain |
| /pyscore | Implement of BTP smart contracts for Python SCORE of ICON blockchain |
| /pyscore/bmv/icon | Implement of BMV for ICON blockchain |
| /pyscore/lib/icon | ICON related common codes |
