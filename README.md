# VITALIK

<br/>
<p align="center">
<img src="img/vitalik.png">
</a>
</p>
<br/>

`Geth` example of `EMV` contract listener that executues upon received `real time events`:

1. `BQ` data insterts from contarcts logs and read funcs
2. `reads` from contract to check logs vs contracts state vars
3. `publish` messages to GCPs `PUB/SUB`
4. `receives` messages from GCPs `PUB/SUB`
5. `writes` to the contracts given received params

- `smart-contract` example is deployed and verified on [mumbai](https://mumbai.polygonscan.com/address/0x54713127daf2bFD5129C980Ea800E3fCD616B547#code)

[polygon-mainnet](https://polygonscan.com/address/0x6Bd249181BAdf2a389296D68f80A8B1c74fDDAC1#code)
[bsc-testnet](https://testnet.bscscan.com/address/0x54713127daf2bFD5129C980Ea800E3fCD616B547#code)
[bsc-mainnet](https://bscscan.com/address/0x6Bd249181BAdf2a389296D68f80A8B1c74fDDAC1#code)

## logic

- smart contract allows users to save for a specific goal, which funds are stored in the contaratcs escrow
- users can scheule multiple payments subcriptions against these goals
- built in order to execute off-chain inetgrators as per scheduled plan and provide business analytics

## dependencies:

- `PRIVATE_KEY="xxx"` in .env
- `GCP_worker.json` in global with rights to read/write to PUB/SUB and BQ
- `BQ-TABLES` as per schema
- `topic-sub` as pub/sub topic
- `gcp-project-id-sub` as pub/sub topic

## run:

- logs listener, contract reader ,bq inserter and pubsub publisher`go run cmd/sub/main.go -config=POLYGON-USDC-PROD | -config=POLYGON-USDC-TEST`
- pubsub listerner and contaract writer`go run cmd/pub/main.go`
