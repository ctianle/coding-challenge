# learning notes

## Installing Ignite
`brew install ignite` (macos)
requires golang

## Creation of new Cosmos SDK blockchain with a starting 'crude' prefix
`ignite scaffold chain crudechain --address-prefix crude`

## 1 Command Trigger for creation of CRUD functionality 
`ignite scaffold list crudeblock name description`

Help
`crudechaind tx crudechain --help`


## Build and deploy
```
ignite chain build
ignite chain serve
```

# Different ports and their purpose of usage

## Tendermint Node (CometBFT RPC endpoint. Default Port 26657)
- Core Consensus Layer for the blockchain
- Ensures validity of transactions and state of the blockchain
- Allows interaction at a low level
- API for interacting with blockchain's internal state and sending/querying transaction data

## Blockchain API (Default Port 1317)
- REST API for external applications or users to interact
- Allows performing high level interactions such as CRUD without neeidng to interact with the Tendermint or low-level data structures.

## Token Faucet (Default Port 4050)
- Service to allow requesting of test token for usage on this blockchain as its a testnet environment. 
- Commonly used for development/testing environments.

## Debug commands
`ignite chain serve --reset-once`

## Debug paths
`~/go/bin/crudechaind`