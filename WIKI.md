# Short Name

Track IoT assets via a Blockchain monitoring UI

# Short Description

Deploy a React UI that has the ability to visualize and submit blockchain transactions, allowing all interested parties to monitor and update assets.

# Offering Type

Emerging Tech (Blockchain + IoT)

# Introduction

This application provides a user interface that allows operators to monitor the progress of IoT assets as they move through a supply chain. This is done by utilizing the Hyperledger Fabric SDK to query a blockchain service. In addition to being able to view individual blocks and their transactions, operators can use the UI to leverage chaincode operations, which have the ability to register new assets and view/update their properties.

# Author

Kalonji Bankole and Alex Nyugen

# Code

https://github.com/IBM/monitoring_ui

# Demo

(In progress)

# Video

https://www.youtube.com/watch?v=DYvdN_p_Ldk

https://www.youtube.com/watch?v=Mw6924hCAIc

# Overview

In this code pattern, we will demonstrate a blockchain monitoring application using React + Node.js with the Hyperledger Fabric SDK. This integration allows users to easily execute actions against the blockchain and monitor the state of assets.

When the reader has completed this code pattern, they will understand how to:

* Deploy a Hyperledger Blockchain network on IBM Cloud
* Create and enroll a administrative client using the Hyperledger Node SDK
* Deploy and Instantiate a smart contract to handle asset updates/queries
* Create a schema describing the properties of an asset
* Monitor and propose blockchain transactions via a UI
* Integrate Watson IoT platform to directly receive asset updates from registered IoT devices via MQTT or HTTP


<!-- This application provides a visual
for individual blocks in a hyperledger blockchain.

Can query any individual assets to view their last known state

 applicable for monitoring assets as they
 toward interested parties of a supply chain

A few examples presented here are

This project is built for developers who are interested in
 are interested in using blockchain
The user interface allows

Interact with a Hyperledger blockchain to track the status of IoT assets in real time.

 developers with a central

Location

This application allows users to monitor their
 real time

Deploy a application that leverages Blockchain integrated with Watson IoT Platform to track  -->



# Flow

1. A request is submitted to Create, Read, Update, or Delete an asset from a blockchain ledger. This request may either be submitted manually by a user via the monitoring UI browser, or from a IoT device (NFC/barcode scanner, etc) publishing a MQTT message to the Watson IoT Platform

2. Node Express backend formats CRUD request into a [jsonrpc](http://www.jsonrpc.org/specification#examples) object like below, and submits it to a Hyperledger peer as a transaction proposal
```
{
    jsonrpc: '2.0',
    method: 'invoke',
    params: {
        type: 1,
        chaincodeID: {
            name: 'simple_contract'
        },
        ctorMsg: {
            function: 'createAsset',
            args: '["assetID", {"carrier": "Port of Long Beach", "longitude":"33.754185", "latitude": "-118.216458", "temperature": "44 F"}]'
        },
        secureContext: 'kkbankol@us.ibm.com'
    },
    id: 5
}
```
<!-- 3. Fabric Node SDK submits CRUD request to Hyperledger peer as a transaction proposal -->

3. Peer uses an "endorsement" service to simulate the proposed transaction against the relevant smart contracts. This endorsement service is used to confirm that the transaction is possible given the current state of the ledger. Examples of invalid proposals might be creating an asset that already exists, querying the state of an asset that does not exist, etc.
<!-- to simulate the transaction request against smart contracts and the current ledger state -->

4. If the simulation is successful, the proposal is then "signed" by the peer's endorser.

5. The signed transaction is forwarded to an ordering service, which executes the transactions and places the result into the ledger

6. The Monitoring UI auto-refreshes to show the transaction result and updated ledger in the "Response Payloads" and "Blockchain" columns, respectively

# Included components

* [Blockchain](https://console.bluemix.net/catalog/services/blockchain)
* [Watson IoT Platform](https://console.bluemix.net/catalog/services/internet-of-things-platform)

# Featured technologies

* [Hyperledger Fabric](https://hyperledger-fabric.readthedocs.io/en/release-1.1/)
* [MQTT](http://mqtt.org/faq)

# Blog

## Blog Title

## Blog Author

## Blog Content

The user interface is divided into three columns.  
1. Chaincode Operations
2. Response Payloads
3. Blockchain

<img src="https://i.imgur.com/BMbb8Oq.png" width="650" height="450" align="center">

### The Chaincode Operations column
The first section of the Monitoring UI is dynamically generated through a combination of JSON Schema and convention. The tabs each represent a subset of the available contract functions and are hardcoded in the `ChaincodeReducer.js` file.     

The contract functions can be selected from the menu in each tab. The functions and their related input fields are defined in the JSON schema.

For example, if we are connected to the IBM sample conctract on the blockchain fabric the **Create** tab includes just one function: `createAsset`. This function maps to the `createAsset` function that is defined in the JSON schema. The UI knows to put `createAsset` under the `create` tab because it matches the tab's name as a substring of the function. The **Read** tab, in contrast, contains three functions each of which are defined in the JSON schema. Each tab also has a corresponding `type`, which controls the use of Hyperledger invoke or query endpoint.

The *arguments* form is generated when you select a particular function. The form creates input fields for the arguments that are defined in the JSON schema. The basic contract includes the following fields: `assetID`, `carrier`, `location`, `temperature`, and `timestamp`.  

**Note:**
- Fields such as `assetID` that are denoted with asterisk are required. Required fields are defined in the JSON schema.
- Location is a nested object with its own properties that in turn are exposed as data fields.
- Validation is defined in the JSON schema and is reflected in the form.  
For example, if you submit the form without an entry for the required `assetID` field, you are prompted to enter a value. If you enter a non-numeric value in the `latitude` or `longitude` fields, you will also be prompted.

When you submit a form the Monitoring UI creates a valid blockchain REST payload with the field input as arguments. The payload and a request is sent to the configured blockchain peer. The Monitoring UI then waits for a response from the peer. The response is displayed in the Request Payload column.

### The Response Payload column
The second column displays the response from the blockchain peer by recursively traversing the payload and writing the responses to the card. If you submit multiple requests from a combination of tabs the Monitoring UI generates cards as needed to display the payload. **Note:** Duplicate REST request with the exact same function and arguments will not create extra cards.

Request payload cards are displayed in the collapsed state. You must expand the cards to view the contents of the card.

Close individual cards by clicking **x** next to the card header.

Click **Clear** on the Request Payload header to remove all payloads from the display.

**Tip:** Enable the **Poll for changes** toggle to have the Monitoring UI actively check for changes to a particular query every time the blockchain height changes. For the basic contract, use this feature to monitoring a particular asset for changes.


### The Blockchain column
The third column shows the current state of the blockchain.

To expand a block, click the expander. The contents of the block show the transactions in the block and the details for each transaction.

**Important:** Any transactions that occur against a specific blockchain will appear within blocks on the blockchain. These include invalid transactions as well as transactions against other contracts. To see a change on a specific contract, the Monitoring UI must be configured to connect to that contract.

# Links

* [title](link): description
