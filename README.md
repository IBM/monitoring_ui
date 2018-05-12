<!--Put badges at the very top -->
<!--change the repos -->
<!--change the tracking number -->
[![Build Status](https://travis-ci.org/IBM/watson-banking-chatbot.svg?branch=master)](https://travis-ci.org/IBM/watson-banking-chatbot)
![IBM Cloud Deployments](https://metrics-tracker.mybluemix.net/stats/527357940ca5e1027fbf945add3b15c4/badge.svg)
<!--Add a new Title and fill in the blanks -->
# Blockchain Monitoring UI
In this Code Pattern, we'll use React.js, Watson IoT Platform, and the Hyperledger Fabric Node SDK to interact with an IBM Blockchain service. The resulting application provides a dynamically generated user interface to monitor assets as they traverse through a supply chain. This solution can be applicable for both physical assets (shipping containers, packages) and financial assets.Operators can use this Monitoring UI to perform actions on the blockchain, see the results of those actions, and query the state of each asset in the blockchain ledger.

When the reader has completed this Code Pattern, they will understand how to:

* Deploy a Hyperledger Blockchain network on IBM Cloud
* Create and enroll a administrative client using the Hyperledger Node SDK
* Deploy and Instantiate a smart contract to handle asset updates/queries
* Create a schema describing the properties of an asset
* Monitor and propose blockchain transactions via a UI
* Integrate Watson IoT platform to directly receive asset updates from registered IoT devices via MQTT or HTTP

<!--Remember to dump an image in this path-->
<p align="center">
<!-- <img src="https://i.imgur.com/lNZxVxo.png"  data-canonical-src="https://i.imgur.com/lNZxVxo.png" width="650" height="450" style="margin-left: auto; margin-right: auto;"> -->
<img src="/images/blockchain_arch.png"  />
</p>

## Flow
<!--Add new flow steps based on the architecture diagram-->
<!-- 1. Upload and Instantiate smart contracts via the Bluemix Network Monitor
2. Deploy the node application locally or on bluemix
3. Input connection information such as service credentials, endpoint, etc into configuration form
4. Submitting form sends a request to pull a json file containing the connection profile. The information from this profile is used to create a "monitoring" client with administrative privileges
5. If form data is valid, user should be able to execute Chaincode operations, view individual blocks and their data, and request state of registered Assets -->

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
<!-- The response is sent back to the Monitoring UI and printed in the "Response Payloads" view.  to show latest blockchain transactions -->

<!-- TODO expand on this -->

<!--Update this section-->
## Included components
* [Blockchain](https://console.bluemix.net/catalog/services/blockchain)
* [Watson IoT Platform](https://console.bluemix.net/catalog/services/internet-of-things-platform)

<!--Update this section-->
## Featured technologies
<!-- Select components from [here](https://github.ibm.com/developer-journeys/journey-docs/tree/master/_content/dev#technologies), copy and paste the raw text for ease -->
* [Hyperledger Fabric](https://hyperledger-fabric.readthedocs.io/en/release-1.1/)
* [MQTT](http://mqtt.org/faq)
* [npm](https://www.npmjs.com/)
* [node.js](https://nodejs.org/en/)

<!--Update this section when the video is created-->
# Watch the Video
<!-- [![](http://img.youtube.com/vi/Jxi7U7VOMYg/0.jpg)](https://www.youtube.com/watch?v=Jxi7U7VOMYg) -->
In progress

https://www.youtube.com/watch?v=DYvdN_p_Ldk

https://www.youtube.com/watch?v=Mw6924hCAIc


# Steps
Use the ``Deploy to IBM Cloud`` button **OR** create the services and run locally.

## Deploy to IBM Cloud
<!--Update the repo and tracking id-->
TODO, In progress
[![Deploy to IBM Cloud](https://metrics-tracker.mybluemix.net/stats/527357940ca5e1027fbf945add3b15c4/button.svg)](https://bluemix.net/deploy?repository=https://github.com/IBM/monitoring_ui.git)

1. Press the above ``Deploy to IBM Cloud`` button and then click on ``Deploy``.

<!--optional step-->
2. In Toolchains, click on ``Delivery Pipeline`` to watch while the app is deployed. Once deployed, the app can be viewed by clicking ``View app``.
![](doc/source/images/toolchain-pipeline.png)

<!--update with service names from manifest.yml-->
3. To see the app and services created and configured for this Code Pattern, use the IBM Cloud dashboard. The app is named `monitoring-ui` with a unique suffix. The following services will be created:
    * Blockchain
    * Internet of Things Platform

## Run locally
> NOTE: These steps are only needed when running locally instead of using the ``Deploy to IBM Cloud`` button.

<!-- there are MANY updates necessary here, just screenshots where appropriate -->
1. [Clone the repo](#1-clone-the-repo)
2. [Create Watson services with IBM Cloud](#2-create-watson-services-with-ibm-cloud)
3. [Upload and Instantiate Chaincode](#3-import-the-conversation-workspace)
4. [Install dependencies](#4-load-the-discovery-documents)
5. [Configure credentials](#5-configure-credentials)
5. [Run the application](#6-run-the-application)

## 1. Clone the repository

Clone the `monitoring_ui` project locally. In a terminal, run:

```
git clone github.com/IBM/monitoring_ui
```

## 2. Create Services

Next, we'll need to deploy our service instances using the IBM Cloud dashboard.

### Watson IoT Platform
We can start by deploying an instance of the Watson IoT Service. In this pattern, the primary role of this service is to offer an secure MQTT broker that'll receive asset updates from IoT devices. Our Monitoring UI subscribes to the MQTT broker and processes incoming messages with the fabric-client node SDK.

First, log into the IBM Cloud dashboard at [https://console.bluemix.net/](https://console.bluemix.net/) and select the ``Catalog`` button in the upper right.

<p align="center">
<img src="https://i.imgur.com/0CctlyI.png"  data-canonical-src="https://i.imgur.com/0CctlyI.png">
</p>

In the search bar type "IoT" and click the icon titled ``Internet of Things Platform Starter``.

<p align="center">
<img src="https://i.imgur.com/GtCcdEJ.png"  data-canonical-src="https://i.imgur.com/GtCcdEJ.png">
</p>

Once this service is provisioned, we'll need to generate a set of credentials for connecting to the broker. We can do so by entering the IoT Platform dashboard, selecting ``Devices`` from the left hand menu, and then clicking the ``Add Device`` button
<p align="center">
<img src="https://i.imgur.com/fec24FG.png"  data-canonical-src="https://i.imgur.com/fec24FG.png">
</p>

Next, provide a device type and ID.
<p align="center">
<img src="https://i.imgur.com/REQfYIK.png"  data-canonical-src="https://i.imgur.com/REQfYIK.png">
</p>

The next two steps (Device Information, Groups) can be skipped.

In the ``Security`` tab, an Authentication token can be entered as long as it meets certain criteria (between 8 and 36 characters, contains mix of lowercase/uppercase letters, numbers, and symbols). Leave this field blank if you'd like for one to be generated instead.

<p align="center">
<img src="https://i.imgur.com/rycnjlF.png"  data-canonical-src="https://i.imgur.com/rycnjlF.png">
</p>

Clicking the ``Finish`` button will generate a set of credentials that can be used to publish messages to the IoT Platform

<p align="center">
<img src="https://i.imgur.com/A2A6yXW.png" width="650" height="450">
</p>

Now, MQTT publish commands can be made from a device in the following format </br>
Client ID: `d:${organization_id}:${device_type}:${device_id}` </br>
Username: `use-token-auth` </br>
Password: `${authentication_token}` </br>
Endpoint: `${organization_id}.messaging.internetofthings.ibmcloud.com` </br>

To publish messages, a MQTT client will need to be installed on the IoT devices responsible for reading and sending asset updates. These clients are very lightweight, and are able to run on resource constrained devices such as Arduino, Raspberry Pi, CHIP, etc.

Now that we have a valid set of credentials, we can use an MQTT client to send a sample command. There are a few clients available online, but for simplicity we'll use a node cli client. This particular client can be installed by running `npm install -g mqtt`, and is also used by the Monitoring UI backend. After exporting the MQTT credentials we can publish a json payload with the following commands

```
organization_id=agf5n9
device_type=assetTracker
device_id=702f6460
username=use-token-auth
password=YTiRp4jRdt4oyKTS3a

mqtt_pub -i "d:${organization_id}:${device_type}:${device_id}" -u "${username}" -P "${password}" -h "${organization_id}.messaging.internetofthings.ibmcloud.com" -p 1883 -t 'iot-2/evt/deviceupdate/fmt/json' -m '{
    "d" : {
          "fcn" : "updateAsset",
          "args" : "[{"carrier": "LBC Freight 647", "longitude":"34.754185", "latitude": "-119.214458", "temperature": "45 F"}]"
          }
}'
```

And then we can see that message has been received by the IoT Platform dashboard by going back to the ``Devices`` menu, selecting our corresponding device, and then selecting ``Recent Events``

<p align="center">
<img src="https://i.imgur.com/lNJ668W.png"  data-canonical-src="https://i.imgur.com/d4QbQFP.png">
</p>

### Blockchain

We can continue on by deploying the IBM Blockchain service. This can be found by logging in to the IBM Cloud [dashboard](https://console.bluemix.net/), selecting the ``Catalog`` button, searching for ``Blockchain``, and clicking on the resulting icon. Or click this [*link*](https://console.bluemix.net/catalog/services/blockchain).

<p align="center">
<img src="https://i.imgur.com/qWQOXq5.png"  data-canonical-src="https://i.imgur.com/qWQOXq5.png">
</p>

After selecting the blockchain icon, a form will be presented for configuring the service name, region, and pricing plan. The default values for these fields can be left as is. Also, be sure that the free pricing tier is selected, which is titled "Starter Membership Plan". If you are using an IBM Cloud Lite account, this plan can be used for free for up to 30 days. After validating that the information in the form is correct, scroll down and click the ``Create`` button in the lower right corner
<p align="center">
<img src="https://i.imgur.com/ROAjOzr.png"  data-canonical-src="https://i.imgur.com/ROAjOzr.png">
</p>

<!-- Provision the following services:
* [**IBM Blockchain**](https://console.bluemix.net/catalog/services/blockchain)
* [**Watson IoT Platform**](https://console.bluemix.net/catalog/services/internet-of-things-platform) -->

<!-- If you're deploying the application via the "Delivery Pipeline" on IBM Cloud, these services should be created automatically.

If you're manually deploying the application and services, -->

## 3. Upload / Instantiate Chaincode
"Smart contracts", commonly referred to as "Chaincode", can be used to execute business logic and validate incoming requests. In this context, the contracts are used to implement CRUD operations for tracking assets on the IBM Blockchain ledger.

To begin the process of uploading the smart contracts to the blockchain, we can start by opening the IBM Cloud dashboard, selecting your provisioned Blockchain service, and accessing the blockchain network monitor by clicking ``Enter Monitor``
<p align="center">
<img src="https://i.imgur.com/J2pbo7H.png"  data-canonical-src="https://i.imgur.com/J2pbo7H.png" width="650" height="450" style="margin-left: auto; margin-right: auto;">
</p>

Next, click the ``Install code`` option on the left hand menu, and then the ``Install Chaincode`` button on the right of the page
<p align="center">
<img src="https://i.imgur.com/HmdDsgm.png"  data-canonical-src="https://i.imgur.com/HmdDsgm.png" width="650" height="450" style="margin-left: auto; margin-right: auto;">
</p>

Enter an id and a version (here we'll use "simple_contract" and "v1"). Then, select the ``Add Files`` button to upload the [samples.go](contracts/basic/simple_contract/samples.go), [schemas.go](contracts/basic/simple_contract/schemas.go), and [simple_contract_hyperledger.go](contracts/basic/simple_contract/simple_contract_hyperledger.go) files

<p align="center">
<img src="https://i.imgur.com/nYwMM47.png"  data-canonical-src="https://i.imgur.com/nYwMM47.png" width="650" height="450" style="margin-left: auto; margin-right: auto;">
</p>

Finally, we'll need to Instantiate the chaincode. This can be done by opening the chaincode options menu and selecting "Instantiate"

This will present a form where arguments can be provided to the chaincodes `init` function. In this case, we'll just need to provide a json string `{"version":"1.0"}` in the Arguments section, and then click ``Submit``
<p align="center">
<img src="https://i.imgur.com/blo1Qx3.png"  data-canonical-src="https://i.imgur.com/blo1Qx3.png" width="450" height="450" style="margin-left: auto; margin-right: auto;">
</p>

For additional documentation on the chaincode implementation, please see the README in the [simple_contract](contracts/basic/simple_contract) directory

## 4. Install dependencies

To start the Monitoring UI, we'll need to install a few node libraries which are listed in our `package.json` file.
- React.js: Used to simplify the generation of front-end components
- MQTT: Client package to subscribe to Watson IoT Platform and handle incoming messages
- Hyperledger Fabric SDK: Enables backend to connect to IBM Blockchain service

### Docker setup (optional)
If you have Docker installed, you can install these dependencies in a virtual container instead. Run the application with the following commands, and then skip to [Step 5](#5-configure-credentials)
```
docker build -t monitoring_ui .
docker run -d -p 8081:8080 monitoring_ui
```
### Manual installation
Otherwise, continue by installing [Node.js](https://nodejs.org/en/) runtime and NPM. Currently the Hyperledger Fabric SDK only appears to work with node v8.9.0+, but [is not yet supported](https://github.com/hyperledger/fabric-sdk-node#build-and-test) on node v9.0+. If your system requires newer versions of node for other projects, we'd suggest using [nvm](https://github.com/creationix/nvm) to easily switch between node versions. We did so with the following commands
```
curl -o- https://raw.githubusercontent.com/creationix/nvm/v0.33.11/install.sh | bash
# Place next three lines in ~/.bash_profile
export NVM_DIR="$HOME/.nvm"
[ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh"  # This loads nvm
[ -s "$NVM_DIR/bash_completion" ] && \. "$NVM_DIR/bash_completion"  # This loads nvm bash_completion
nvm install v8.9.0
nvm use 8.9.0
```

Install the Monitoring UI node packages by running `npm install` in the project root directory and in the [react-backend](react-backend) directory. Both `python` and `build-essential` are required for these dependencies to install properly:
```
npm install
cd react-backend && npm install
```

Finally, compile the `bundle.js` file
```
cd public
npm run build
```
<!-- Method	| Command	|Comment
--- | --- | ---
Filesystem | `npm run build` | The build command generates the bundle.js file in the public directory. </br>To access the Monitoring UI, go to the `monitoring_ui/public` directory and open the *index.html* file in a browser. -->

<!-- Launch the **Watson Conversation** tool. Use the **import** icon button on the right

Find the local version of [`data/conversation/workspaces/banking.json`](data/conversation/workspaces/banking.json) and select
**Import**. Find the **Workspace ID** by clicking on the context menu of the new
workspace and select **View details**. Save this ID for later.

*Optionally*, to view the conversation dialog select the workspace and choose the
**Dialog** tab, here's a snippet of the dialog: -->

![](doc/source/images/dialog.PNG)

## 4. Run the application

1. Start the app locally with `npm run dev-server`.

<!-- This method is ideal for a development environment but not suitable for a production environment. TODO, this comment is from the original author, would like to understand why-->

2. To access the Monitoring UI, open the following URL in a browser: `http://localhost:8081/` </br>**Note:** If you run into an issue with the port already being used, set the `PORT` environment variable to the port you'd like to use. Note that hot reload is enabled for the webpack-dev-server. Changes that you save to the source are immediately reflected in the Monitoring UI. There is no need to manually reload.
> Note: server host can be changed as required in server.js and `PORT` can be set in `.env`.

<!--Add a section that explains to the reader what typical output looks like, include screenshots -->

<p align="center">
<img src="https://i.imgur.com/BMbb8Oq.png"  data-canonical-src="https://i.imgur.com/BMbb8Oq.png" width="750" height="450" style="margin-left: auto; margin-right: auto;">
</p>

<!--Include any troubleshooting tips (driver issues, etc)-->

## 5. Obtain credentials

The credentials for IBM Cloud services (Blockchain, Watson IoT Platform), can be found in the ``Services`` menu in IBM Cloud by selecting the ``Service Credentials`` option for each service.

The Blockchain credentials consist of the `key`, `secret`, and `network_id` parameters
<!-- ![]("https://i.imgur.com/Qof7sve.png" width="250" height="400") -->
<p align="center">
<img src="https://i.imgur.com/Qof7sve.png"  data-canonical-src="https://i.imgur.com/Qof7sve.png" width="450" height="450" style="margin-left: auto; margin-right: auto;">
</p>

These credentials will need to be provided to the UI in the next step
<!-- We can obtain the Blockchain credentials by downloading

Copy the [`env.sample`](env.sample) to `.env`.

```
$ cp env.sample .env
```
Edit the `.env` file with the necessary settings.

#### `env.sample:`

```
# Replace the credentials here with your own.
# Rename this file to .env before starting the app.

# Watson conversation
CONVERSATION_USERNAME=<add_conversation_username>
CONVERSATION_PASSWORD=<add_conversation_password>
WORKSPACE_ID=<add_conversation_workspace>

# Watson Discovery
DISCOVERY_USERNAME=<add_discovery_username>
DISCOVERY_PASSWORD=<add_discovery_password>
DISCOVERY_ENVIRONMENT_ID=<add_discovery_environment>
DISCOVERY_COLLECTION_ID=<add_discovery_collection>

# Watson Natural Language Understanding
NATURAL_LANGUAGE_UNDERSTANDING_USERNAME=<add_nlu_username>
NATURAL_LANGUAGE_UNDERSTANDING_PASSWORD=<add_nlu_password>

# Watson Tone Analyzer
TONE_ANALYZER_USERNAME=<add_tone_analyzer_username>
TONE_ANALYZER_PASSWORD=<add_tone_analyzer_password>

# Run locally on a non-default port (default is 3000)
# PORT=3000

``` -->

## 6. UI Configuration

Before we're able to access blockchain information via the Monitoring UI, we'll need to provide a bit of information about our Blockchain service, such as the API credentials/endpoint, and chaincode information.

First, access the configuration form by clicking **CONFIGURATION**.  

<p align="center">
<img src="https://i.imgur.com/pS3s5vg.png" width="650" height="450" style="margin-left: auto; margin-right: auto;">
</p>

This form accepts the following parameters:

Parameter	|Value	|Comment
--- | --- | ---
API Host and Port	| http://peer_URL:port	| The host and port for the IBM Blockchain REST API prepended with `http://`.
Chaincode ID	| The contract ID that was returned when you registered the contract.	| The contract ID should be a string that was provided in the previous "Install / Instantiate Chaincode" step. If the ID is incorrectly entered, the UI will display the blockchain ledger entries, but the asset search function will not work.
Secure Context|Your fabric user	| This is required for connecting to Blockchain instances on IBM Cloud. </br>
**Important:** For secureContext use the user name that was used to configure the fabric.
Number of blocks to display	| A positive integer. Default: 10	| The number of blockchain blocks to display.
Key | API Key | Provided in IBM Cloud Credentials
Secret | API Secret | Provided in IBM Cloud Credentials
Network Id | Blockchain network id | Provided in IBM Cloud Credentials

After filling out and submitting the form, a request will be sent to the `/init_client` endpoint with the provided parameters.

Once this endpoint is invoked, it will call a series of methods to carry out the following
- Fetch the network [configuration file](https://hyperledger.github.io/composer/latest/reference/connectionprofile). This file contains information about all of the components in the blockchain network, such as the name/endpoints of associated peers, channels, organizations, orderers, and certificate authorities
- Initialize our hyperledger "client" by loading the configuration file via the Node Hyperledger SDK
- Create/enroll a fabric user named "monitoring_user" with administrative privileges. These elevated privileges are required to invoke and query chaincode methods. To obtain these privileges, the SDK must send a request to the blockchain "certificate authority" listed in the connection profile for a certificate/key pair. Once this is complete, a PEM encoded Certificate will be output to the server logs like so.
</br>
<!-- point it to a blockchain peer server and provide a contract ID to monitor. A "peer" is a member of a blockchain network that hosts copies of a ledger and smart contracts -->
<p align="center">
<img src="https://i.imgur.com/5ZRGcux.png" width="450" height="350" style="margin-left: auto; margin-right: auto;">
</p>
</br>
The certificate will need to be manually uploaded to one of the blockchain peers via the blockchain network monitor. A "[peer](http://hyperledger-fabric.readthedocs.io/en/release-1.1/peers/peers.html)" is a member of a blockchain network that's responsible for hosting copies of the ledger and chaincode. These are the primary interaction endpoint for both administrators and members with limited access. For a client to carry out administrative requests, they must have a PEM certificate that has been uploaded to the peer. The chaincode operations will not work until this step has been completed.

This can be done by going back to the IBM Cloud dashboard, selecting your provisioned Blockchain service, and accessing the blockchain network monitor by clicking ``Enter Monitor``
<p align="center">
<img src="https://i.imgur.com/J2pbo7H.png" width="650" height="450" style="margin-left: auto; margin-right: auto;">
</p>

After entering the Network Monitor, select the ``Members`` section from the left hand menu. Then, select the ``Certificate`` option, and click ``Add certificate``
<p align="center">
<img src="https://i.imgur.com/cKZthHB.png" width="650" height="450" style="margin-left: auto; margin-right: auto;">
</p>

Clicking ``Add certificate`` will present the following form
<p align="center">
<img src="https://i.imgur.com/cD125af.png" width="650" height="450" style="margin-left: auto; margin-right: auto;">
</p>

# Troubleshooting

* `sendPeersProposal - Promise is rejected: Error: 2 UNKNOWN: chaincode error (status: 500, message: Authorization for GETINSTALLEDCHAINCODES on channel getinstalledchaincodes has been denied with error Failed verifying that proposal's creator satisfies local MSP principal during channelless check policy with policy [Admins]: [This identity is not an admin]`
> This error occurs if the certificate generated by the SDK user has not been uploaded to the peer

* `Error: The gRPC binary module was not installed. This may be fixed by running "npm rebuild"`
> `grpc` is a requirement for the fabric-client SDK. Confirm that is has been installed in the `react_backend` directory with `npm install grpc@1.11.0`

<!-- * Error: Environment {GUID} is still not active, retry once status is active

  > This is common during the first run. The app tries to start before the Discovery
environment is fully created. Allow a minute or two to pass. The environment should
be usable on restart. If you used `Deploy to IBM Cloud` the restart should be automatic.

* Error: Only one free environment is allowed per organization

  > To work with a free trial, a small free Discovery environment is created. If you already have
a Discovery environment, this will fail. If you are not using Discovery, check for an old
service thay you may want to delete. Otherwise use the .env DISCOVERY_ENVIRONMENT_ID to tell
the app which environment you want it to use. A collection will be created in this environment
using the default configuration. -->

<!--This can stay as-is if using Deploy to IBM Cloud-->

<!--Include any relevant links-->

# Links
Blockchain Supply Chain articles
- https://aqurus.ca/blockchain-crucial-link-supply-chain/
- https://medium.com/scandinavian-cryptocurrency-exchange/how-blockchain-technology-can-optimize-product-supply-chains-53164a11a1ba

Submit scanned barcode result to Watson IoT Platform
- https://www.kevinhoyt.com/2016/11/09/tessel-barcode-scanner-software/
<!-- * [Demo on Youtube](https://www.youtube.com/watch?v=Jxi7U7VOMYg) -->
* [Hyperledger Node.js SDK](https://github.com/hyperledger/fabric-sdk-node)

<!-- pick the relevant ones from below -->
# Learn more

* **IoT Code Patterns**: Enjoyed this Code Pattern? Check out our other [IoT Code Patterns](https://developer.ibm.com/code/technologies/iot/).
* **Blockchain Patterns**: Enjoyed this Code Pattern? Check out our other [Blockchain Patterns](https://developer.ibm.com/code/technologies/blockchain/)
* **Emerging Tech Code Pattern Playlist**: Bookmark our [playlist](https://www.youtube.com/playlist?list=PLzUbsvIyrNfkmf4_91eLqELe6e0tFR_9W) with all of our Code Pattern videos
* **Kubernetes on IBM Cloud**: Deliver your apps with the combined the power of [Kubernetes and Docker on IBM Cloud](https://www.ibm.com/cloud-computing/bluemix/containers)

<!--keep this-->

# License
[Apache 2.0](LICENSE)
