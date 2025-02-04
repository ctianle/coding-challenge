# learning notes

## Installing Ignite
`brew install ignite` (macos)
requires golang

## Creation of new Cosmos SDK blockchain
`ignite scaffold chain crudechain`

## Creation of Resource module
`ignite scaffold module resourcemodule`

## Function for creation of resource
`ignite scaffold message create-resource <field1> <field2> --module resourcemodule`

## Function for listing of resource
`ignite scaffold query list-resource --module resourcemodule`

## Function for getting details of a resource
`ignite scaffold query get-resource id --module resourcemodule`

## Function for updating a resource`
`ignite scaffold message update-resource id title description --module resourcemodule`

## Function for deletion of a resource`
`ignite scaffold message delete-resource id --module resourcemodule`

## Build and deploy
```
ignite chain build
ignite chain serve
```


# Different ports and their purpose of usage

## Tendermint Node 
- Core Consensus Layer for the blockchain
- Ensures validity of transactions and state of the blockchain
- Allows interaction at a low level
- API for interacting with blockchain's internal state and sending/querying transaction data

## Blockchain API 
- REST API for external applications or users to interact
- Allows performing high level interactions such as CRUD without neeidng to interact with the Tendermint or low-level data structures.

## Token Faucet
- Service to allow requesting of test token for usage on this blockchain as its a testnet environment. 
- Commonly used for development/testing environments.