# loan
**loan** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Get started

```
ignite scaffold chain loan --no-module && cd loan
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com).

### Web Frontend

Additionally, Ignite CLI offers both Vue and React options for frontend scaffolding:

For a Vue frontend, use: `ignite scaffold vue`
For a React frontend, use: `ignite scaffold react`
These commands can be run within your scaffolded blockchain project. 


For more information see the [monorepo for Ignite front-end development](https://github.com/ignite/web).

## Release
To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

### Install
To install the latest version of your blockchain node's binary, execute the following command on your machine:

```
curl https://get.ignite.com/username/loan@latest! | sudo bash
```
`username/loan` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).

## Learn more

- [Ignite CLI](https://ignite.com/cli)
- [Tutorials](https://docs.ignite.com/guide)
- [Ignite CLI docs](https://docs.ignite.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/ignite)


## Create a module
Create a new "loan" module that is based on the standard Cosmos SDK bank module.
```
ignite scaffold module loan --dep bank
```

## Define the loan module
The "list" scaffolding command is used to generate files that implement the logic for storing and interacting with data stored as a list in the blockchain state.
```
ignite scaffold list loan amount fee collateral deadline state borrower lender --no-message
```

## Scaffold the Messages
Scaffold the code for handling the messages for requesting, approving, repaying, liquidating, and cancelling loans.

- Handling loan requests
```
ignite scaffold message request-loan amount fee collateral deadline
```

- Approving and Canceling Loans

```
ignite scaffold message approve-loan id:uint
```
```
ignite scaffold message cancel-loan id:uint
```

- Repaying and Liquidating Loans

```
ignite scaffold message repay-loan id:uint
```
```
ignite scaffold message liquidate-loan id:uint
```

## Testing the application
### Add test tokens
Configure config.yml to add tokens (e.g., 10000foocoin) to test accounts.
```
version: 1
accounts:
  - name: alice
    coins:
      - 20000token
      - 10000foocoin
      - 200000000stake
  - name: bob
    coins:
      - 10000token
      - 100000000stake
client:
  openapi:
    path: docs/static/openapi.yml
faucet:
  name: bob
  coins:
    - 5token
    - 100000stake
validators:
  - name: alice
    bonded: 100000000stake
```


### Start the Blockchain:
```
ignite chain serve
```
If everything works successful, you should see the Blockchain is running message in the Terminal.

### Perform loan operations
```
loand tx loan request-loan 1000token 100token 1000foocoin 500 --from alice --chain-id loan
loand tx loan approve-loan 0 --from bob --chain-id loan
loand tx loan repay-loan 0 --from alice --chain-id loan

loand tx loan request-loan 1000token 100token 1000foocoin 20 --from alice --chain-id loan -y
loand tx loan approve-loan 1 --from bob --chain-id loan -y
loand tx loan liquidate-loan 1 --from bob --chain-id loan -y

loand q loan list-loan

loand q bank balances $(loand keys show alice -a)
```