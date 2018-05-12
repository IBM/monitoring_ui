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
Leverage Blockchain and IoT to secure your supply chain

## Blog Author
Kalonji Bankole

## Blog Content

The Internet of Things (IoT) is turning out to one of the most disruptive and useful forms of technology. However, there are a quite a few challenges that come with implementing an IoT system, such as:
- Security: How can we prevent IoT devices from being compromised? If a device does get compromised, how can we minimize the potential damage from said device?
- Redundancy: How can we prevent data from being lost due to corruption, accidental deletion, etc?
- Transparency: How can we allow interested parties to see the state of IoT devices without compromising security?

In this code pattern, we'll demonstrate an application that can address a few of these concerns by integrating a Hyperledger blockchain service with the Watson IoT platform. This blockchain integration allows registered IoT devices to efficiently create and track an Asset as it travels through a supply chain. Using the blockchain in this case will allow interested parties to see the owner/state of an asset in real time as it is updated, and for transactions to be validated before they are carried out.

This application also offers a front end UI that allows users to view each individual block and associated transaction. This pattern uses two IBM Cloud hosted Services: IBM Blockchain and the Watson IoT Platform.

# Links
Blockchain Supply Chain articles
- https://aqurus.ca/blockchain-crucial-link-supply-chain/
- https://medium.com/scandinavian-cryptocurrency-exchange/how-blockchain-technology-can-optimize-product-supply-chains-53164a11a1ba

Tutorial to scanned barcode result to Watson IoT Platform
- https://www.kevinhoyt.com/2016/11/09/tessel-barcode-scanner-software/
