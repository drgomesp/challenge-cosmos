#!/usr/bin/env bash

rm -r ~/.rpsd || true
BIN=$(which rpsd)
CHAIN_ID="rps-1"

# configure rpsd
$BIN config set client chain-id $CHAIN_ID
$BIN config set client keyring-backend test #Define the key handle (test, os, file)

# Create my test accounts (alice and bob)
$BIN keys add alice
$BIN keys add bob

# Init node (generate the genesis file)
$BIN init test --chain-id $CHAIN_ID --default-denom rps

# Copy my genesis into the blockchain gensis file
cp ./genesis.json ~/.rpsd/config/genesis.json

# Assisign funds to the alice and bob accounts
$BIN genesis add-genesis-account alice 10000000rps --keyring-backend test
$BIN genesis add-genesis-account bob 1000rps --keyring-backend test

# Create the validator (create gentx tx)
$BIN genesis gentx alice 1000000rps --chain-id $CHAIN_ID

# Save the transactions gentx (transactions to register validator)
$BIN genesis collect-gentxs
