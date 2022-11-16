# VITALIK

<br/>
<p align="center">
<img src="img/vitalik.png">
</a>
</p>
<br/>

`Geth` example of `EMV` contract listener that executues upon received `real time events`:

1. `BQ` insterts ie `data engineering`
2. `reads` from contract to check logs vs contracts state vars
3. `publish` messages to GCPs `PUB/SUB`
4. `receives` messages from GCPs `PUB/SUB`
5. `writes` to the contracts given received params

- `smart-contract` example is deployed and verified on [alfajores](https://explorer.celo.org/alfajores/address/0xa883d9C6F7FC4baB52AcD2E42E51c4c528d7F7D3/contracts)

- with [owner](https://explorer.celo.org/alfajores/address/0x741e0608906B74B8754a99413A7374FdE7B9779a/transactions)

## logic

- smart contract allows users to save for a specific goal, which funds are stored in the contaratcs escrow
- users can scheule multiple payments subcriptions against these goals
- built in order to execute off-chain inetgrators as per scheduled plan and provide business analytics

## dependencies:

- `PRIVATE_KEY="xxx"` in .env
- `GCP_worker.json` in global with rights to read/write to PUB/SUB and BQ
- `BQ-TABLES` as per schema
- `intervals-sub` as pub/sub topic

## run:

- logs listener, contract reader ,bq inserter and pubsub publisher`go run cmd/sub/main.go`
- pubsub listerner and contaract writer`go run cmd/pub/main.go`

## TBD

1. UNIT TEST
2. NETWORK TEST
3. PB in PUB/SUB
4. Docker/K8s deployment
5. Detach smart contract implementation from lib
