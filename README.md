# VITALIK

<br/>
<p align="center">
<img src="img/vitalik.png">
</a>
</p>
<br/>

`Geth` based project that listens to `real time events` from below smart contract and executes:

1. `BQ` insterts upon received events
2. `reads-writes` from contract
3. publish messages to GCPs `PUB/SUB`

- `smart-contract` Deployed and Verified on [alfajores](https://explorer.celo.org/alfajores/address/0xa883d9C6F7FC4baB52AcD2E42E51c4c528d7F7D3/contracts)

- with [owner](https://explorer.celo.org/alfajores/address/0x741e0608906B74B8754a99413A7374FdE7B9779a/transactions)

## logic

- smart contract allows users to save for a specific goal, which funds are stored in the contaratcs escrow
- users can scheule multiple payments subcriptions against these goals
- built in order to execute off-chain inetgrators as per scheduled plan and provide business analytics

## dependencies:

- `PRIVATE_KEY="xxx"` in .env
- `GCP.json` in global with rights to read/write to PUB/SUB and BQ
- `BQ-TABLES` as per schema
- `intervals` as pub/sub topic

## run:

- `go run cmd/main.go`

## TBD

1. UNIT TEST
2. NETWORK TEST
3. PB in PUB/SUB
4. Docker/K8s deployment
5. Detach smart contract implementation from lib
