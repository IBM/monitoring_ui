<!--Put badges at the very top -->
<!--change the repos -->
<!--change the tracking number -->
[![Build Status](https://travis-ci.org/IBM/watson-banking-chatbot.svg?branch=master)](https://travis-ci.org/IBM/watson-banking-chatbot)
![IBM Cloud Deployments](https://metrics-tracker.mybluemix.net/stats/527357940ca5e1027fbf945add3b15c4/badge.svg)
<!--Add a new Title and fill in the blanks -->
# Blockchain Monitoring UI
In this Code Pattern, we'll use React.js, Watson IoT Platform, and the Hyperledger Fabric SDK to interact with an IBM Blockchain service. The resulting application provides a dynamically generated user interface to monitor assets as they traverse through a supply chain. Operators can use this Monitoring UI to perform actions on the blockchain, see the results of those actions, and query the state of each asset in the blockchain ledger.

When the reader has completed this Code Pattern, they will understand how to:

* Deploy a smart contract to handle asset updates/queries
* Create a schema describing the properties of an asset
* Monitor and propose blockchain transactions via a UI
* Integrate Watson IoT platform to directly receive asset updates from registered IoT devices via MQTT or HTTP

<!--Remember to dump an image in this path-->
<p align="center">
<!-- <img src="https://i.imgur.com/lNZxVxo.png"  data-canonical-src="https://i.imgur.com/lNZxVxo.png" width="650" height="450" style="margin-left: auto; margin-right: auto;"> -->
<img src="/images/architecture.png"  />
</p>

## Flow
<!--Add new flow steps based on the architecture diagram-->
<!-- 1. Upload and Instantiate smart contracts via the Bluemix Network Monitor
2. Deploy the node application locally or on bluemix
3. Input connection information such as service credentials, endpoint, etc into configuration form
4. Submitting form sends a request to pull a json file containing the connection profile. The information from this profile is used to create a "monitoring" client with administrative privileges
5. If form data is valid, user should be able to execute Chaincode operations, view individual blocks and their data, and request state of registered Assets -->
1. User submits CRUD request through monitoring_ui **OR** IoT Device scans Asset (barcode, NFC) and publishes "update" message to Watson IoT Platform

2. Node Express backend receives request from user or from Watson IoT platform via MQTT subscriber

2. Request is formatted into a jsonrpc object like so.
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
            args: ["assetID", '{"carrier": "Port of Long Beach", "longitude":"33.754185", "latitude": "-118.216458", "temperature": "44 F"}']
        },
        secureContext: 'kkbankol@us.ibm.com'
    },
    id: 5
}
```
4. Fabric SDK is used to forward formatted request as a transaction proposal to hyperledger service

5. If proposal is accepted, transaction is then submitted to hyperledger peer

5. Result is printed in "Response Payloads" section in monitoring UI

6. Monitoring UI auto-refreshes to show latest blockchain transactions

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
2. In Toolchains, click on Delivery Pipeline to watch while the app is deployed. Once deployed, the app can be viewed by clicking 'View app'.
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

### 1. Clone the repo

Clone the `monitoring_ui` project locally. In a terminal, run:

```
$ git clone github.com/IBM/monitoring_ui
```

### 2. Create Watson services with IBM Cloud

Provision the following services:
* [**IBM Blockchain**](https://console.bluemix.net/catalog/services/blockchain)
* [**Watson IoT Platform**](https://console.bluemix.net/catalog/services/internet-of-things-platform)

### 3. Upload / Instantiate Chaincode
The smart contracts, commonly referred to as "Chaincode", can be used to execute business logic and validate incoming requests

In this context, the contracts are used to implement CRUD operations for tracking assets on the IBM Blockchain ledger.

To begin the process of uploading the smart contracts, we can start by opening the Bluemix UI, selecting your provisioned Blockchain service, and accessing the blockchain network monitor by clicking "Enter Monitor"
<p align="center">
<img src="https://i.imgur.com/J2pbo7H.png"  data-canonical-src="https://i.imgur.com/J2pbo7H.png" width="650" height="450" style="margin-left: auto; margin-right: auto;">
</p>

Next, click the "Install code" option on the left hand menu, and then the "Install Chaincode" button on the right of the page
<p align="center">
<img src="https://i.imgur.com/HmdDsgm.png"  data-canonical-src="https://i.imgur.com/HmdDsgm.png" width="650" height="450" style="margin-left: auto; margin-right: auto;">
</p>

Enter an id ("simple_contract") and a version ("v1"). Then, select the "Add Files" button to upload the [samples.go](contracts/basic/simple_contract/samples.go), [schemas.go](contracts/basic/simple_contract/schemas.go), and [simple_contract_hyperledger.go](contracts/basic/simple_contract/simple_contract_hyperledger.go) files

<p align="center">
<img src="https://i.imgur.com/nYwMM47.png"  data-canonical-src="https://i.imgur.com/nYwMM47.png" width="650" height="450" style="margin-left: auto; margin-right: auto;">
</p>

Finally, we'll need to Instantiate the chaincode. This can be done by opening the chaincode options menu and selecting "Instantiate"

This will present a form where arguments can be provided to the chaincodes `init` function. In this case, we'll just need to provide a json string `{"version":"1.0"}` in the Arguments section, and then click "Submit"
<p align="center">
<img src="https://i.imgur.com/blo1Qx3.png"  data-canonical-src="https://i.imgur.com/blo1Qx3.png" width="450" height="450" style="margin-left: auto; margin-right: auto;">
</p>

For additional documentation on the chaincode implementation, please see the README in the [simple_contract](contracts/basic/simple_contract) directory

### 4. Install dependencies

If you have docker installed, you can run the application with the following commands, and then skip to [Step 5](#5-configure-credentials)
```
docker build -t monitoring_ui .
docker run -d -p 8081:8080 monitoring_ui
```


Otherwise, continue by installing [Node.js](https://nodejs.org/en/) runtime and NPM. Currently the Hyperledger fabric-sdk only appears to work with node v8.9.0+, but [is not yet supported](https://github.com/hyperledger/fabric-sdk-node#build-and-test) on node v9.0+. If your system requires newer versions of node for other projects, we'd suggest using [nvm](https://github.com/creationix/nvm) to easily switch between node versions. We did so with the following commands
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

### 4. Run the application

1. Start the app locally with `npm run dev-server`.

This method is ideal for a development environment but not suitable for a production environment.

2. To access the Monitoring UI, open the following URL in a browser: `http://localhost:8081/` </br>**Note:** If you run into an issue with the port already being used, set the `PORT` environment variable to the port you'd like to use. Note that hot reload is enabled for the webpack-dev-server. Changes that you save to the source are immediately reflected in the Monitoring UI. There is no need to manually reload.
> Note: server host can be changed as required in server.js and `PORT` can be set in `.env`.

<!--Add a section that explains to the reader what typical output looks like, include screenshots -->

<p align="center">
<img src="https://i.imgur.com/BMbb8Oq.png"  data-canonical-src="https://i.imgur.com/BMbb8Oq.png" width="750" height="450" style="margin-left: auto; margin-right: auto;">
</p>

<!--Include any troubleshooting tips (driver issues, etc)-->

### 5. Configure credentials

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

### 6. UI Configuration

Before you can access blockchain information via the Monitoring UI, you must point it to a blockchain peer server and provide a contract ID to monitor. Access the configuration by clicking **CONFIGURATION**.  

<p align="center">
<img src="https://i.imgur.com/pS3s5vg.png" width="650" height="450" style="margin-left: auto; margin-right: auto;">
</p>

This form accepts the following parameters:

Parameter	|Value	|Comment
--- | --- | ---
API Host and Port	| http://peer_URL:port	| The host and port for the IBM Blockchain REST API prepended with `http://`.
Chaincode ID	| The contract ID that was returned when you registered the contract.	| The contract ID is a 128-character alphanumeric hash that corresponds to the Contract ID entry. </br> **Important:** As you cut-and-paste the contract ID, make sure that no spaces are included in the ID. If the ID is incorrectly entered, the UI will display the blockchain ledger entries, but the asset search function will not work.
Secure Context|Your fabric user	| This is required for connecting to IBM Blockchain instances on Bluemix. </br>
**Important:** For secureContext use the user name that was used to configure the fabric.
Number of blocks to display	| A positive integer. Default: 10	| The number of blockchain blocks to display.
<!-- Key | API Key | This is the Hyperledger API key that can be found in the Bluemix console credentials
Secret | API Secret | This is the Hyperledger API secret that can be found in the Bluemix console credentials
Network Id |  | -->

**Important** After submitting the form, a request will be sent to the `/init_client` endpoint with the provided parameters. This will fetch the network configuration file and create/enroll a fabric user named "monitoring_user". Once this is complete, a PEM encoded Certificate will be output to the server logs like so.

<p align="center">
<img src="https://i.imgur.com/5ZRGcux.png" width="450" height="350" style="margin-left: auto; margin-right: auto;">
</p>

This certificate will need to be manually uploaded via the blockchain service UI. The chaincode operations will not work until this step has been completed.

This can be done by going back to the Bluemix UI, selecting your provisioned Blockchain service, and accessing the blockchain network monitor by clicking "Enter Monitor"
<p align="center">
<img src="https://i.imgur.com/J2pbo7H.png" width="650" height="450" style="margin-left: auto; margin-right: auto;">
</p>

After entering the Network Monitor, select the "Members" section from the left hand menu. Then, select the "Certificate" option, and click "Add certificate"
<p align="center">
<img src="https://i.imgur.com/cKZthHB.png" width="650" height="450" style="margin-left: auto; margin-right: auto;">
</p>

Clicking "Add certificate" will present the following form
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

# Privacy Notice
If using the `Deploy to IBM Cloud` button some metrics are tracked, the following
information is sent to a [Deployment Tracker](https://github.com/IBM/cf-deployment-tracker-service) service
on each deployment:

* Node.js package version
* Node.js repository URL
* Application Name (`application_name`)
* Application GUID (`application_id`)
* Application instance index number (`instance_index`)
* Space ID (`space_id`)
* Application Version (`application_version`)
* Application URIs (`application_uris`)
* Labels of bound services
* Number of instances for each bound service and associated plan information

This data is collected from the `package.json` file in the sample application and the `VCAP_APPLICATION` and `VCAP_SERVICES` environment variables in IBM Cloud and other Cloud Foundry platforms. This data is used by IBM to track metrics around deployments of sample applications to IBM Cloud to measure the usefulness of our examples, so that we can continuously improve the content we offer to you. Only deployments of sample applications that include code to ping the Deployment Tracker service will be tracked.

## Disabling Deployment Tracking

To disable tracking, simply remove ``require("cf-deployment-tracker-client").track();`` from the ``app.js`` file in the top level directory.

<!--Include any relevant links-->

# Links
Blockchain Supply Chain articles
- https://aqurus.ca/blockchain-crucial-link-supply-chain/
- https://medium.com/scandinavian-cryptocurrency-exchange/how-blockchain-technology-can-optimize-product-supply-chains-53164a11a1ba

<!-- * [Demo on Youtube](https://www.youtube.com/watch?v=Jxi7U7VOMYg) -->
* [Hyperledger Node.js SDK](https://github.com/hyperledger/fabric-sdk-node)

<!-- pick the relevant ones from below -->
# Learn more

* **Artificial Intelligence Code Patterns**: Enjoyed this Code Pattern? Check out our other [AI Code Patterns](https://developer.ibm.com/code/technologies/artificial-intelligence/).
* **Data Analytics Code Patterns**: Enjoyed this Code Pattern? Check out our other [Data Analytics Code Patterns](https://developer.ibm.com/code/technologies/data-science/)
* **AI and Data Code Pattern Playlist**: Bookmark our [playlist](https://www.youtube.com/playlist?list=PLzUbsvIyrNfknNewObx5N7uGZ5FKH0Fde) with all of our Code Pattern videos
* **With Watson**: Want to take your Watson app to the next level? Looking to utilize Watson Brand assets? [Join the With Watson program](https://www.ibm.com/watson/with-watson/) to leverage exclusive brand, marketing, and tech resources to amplify and accelerate your Watson embedded commercial solution.
* **Data Science Experience**: Master the art of data science with IBM's [Data Science Experience](https://datascience.ibm.com/)
* **PowerAI**: Get started or get scaling, faster, with a software distribution for machine learning running on the Enterprise Platform for AI: [IBM Power Systems](https://www.ibm.com/ms-en/marketplace/deep-learning-platform)
* **Spark on IBM Cloud**: Need a Spark cluster? Create up to 30 Spark executors on IBM Cloud with our [Spark service](https://console.bluemix.net/catalog/services/apache-spark)
* **Kubernetes on IBM Cloud**: Deliver your apps with the combined the power of [Kubernetes and Docker on IBM Cloud](https://www.ibm.com/cloud-computing/bluemix/containers)

<!--keep this-->

# License
[Apache 2.0](LICENSE)
