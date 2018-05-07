var express = require('express');
var router = express.Router();
var _ = require('underscore')
const hfc = require('fabric-client')
var CAClient = require('fabric-ca-client')
//var config = hfc.loadFromConfig('./connection_profile.json')
var crypto = require('crypto')
var util = require('util')
var fs = require('fs');
const request = require('request')
// var config
/* GET home page. */
router.get('/', function(req, res, next) {
  res.render('index', { title: 'Express' });
});

router.post('/init_client', function (req, res) {
  console.log("Initializing Client")
  if (! fs.existsSync('./connection_profile.json')) {
    var options = {
        url: req.body.urlRestRoot + '/networks/' + req.body.networkId + '/connection_profile',
        method: 'GET',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json',
            'Accept-Charset': 'utf-8',
            "Authorization": "Basic " + new Buffer(req.body.key + ":" + req.body.secret, "utf8").toString("base64")
        }
    }
    request(options, function(err, res, body) {
        let json = JSON.parse(body);
        fs.writeFile('./connection_profile.json', body, 'utf8', function(err){
          if(err) {console.log(err)}
        })
        config = hfc.loadFromConfig(json)
    });
  } else {
    config = hfc.loadFromConfig('./connection_profile.json')
  }

  var org = Object.keys(config._network_config._network_config.organizations)[0]
  var certificateAuthorities = config._network_config._network_config.certificateAuthorities
  var certificateAuthorityName = Object.keys(certificateAuthorities)[0]
  var certificateAuthObj = certificateAuthorities[certificateAuthorityName]
  var registrar = config._network_config._network_config.certificateAuthorities[certificateAuthorityName].registrar[0]
  var mspId = config._network_config._network_config.organizations[org]['mspid']
  var storePath = './'
  var client_crypto_suite = hfc.newCryptoSuite()
  var crypto_store = hfc.newCryptoKeyStore({path: storePath})
  var crypto_suite = hfc.newCryptoSuite()
  crypto_suite.setCryptoKeyStore(crypto_store)
  // var crypto_store = hfc.newCryptoKeyStore({path: storePath})
  // crypto_suite.setCryptoKeyStore(crypto_store)
  config.setCryptoSuite(crypto_suite)
  // config.setCryptoSuite(client_crypto_suite);
  hfc.newDefaultKeyValueStore({path: storePath}).then( (store) => {
    config.setStateStore(store)
  }).then( (result) => {
    config.getUserContext('monitoring_admin', true).then ( (user) => {
    // res.send("Client Initialized")
    // console.log("Client Initialized")
    if (user && user.isEnrolled()) {
      console.log("Client Loaded From Persistence")
      res.send("Client Loaded From Persistence")
      console.log("Be sure to upload following cert via blockchain UI: \n" + req.body.urlRestRoot + "/network/" + req.body.networkId + "/members/certificates")
      console.log(user._signingIdentity._certificate + '\n')
    } else {
      console.log("Monitoring Client User doesn't exist. Loading CA client to enroll")
      var ca = new CAClient(certificateAuthObj.url, {
        trustedRoots: [],
        verify: false
      }, certificateAuthObj.caName, crypto_suite)
      enrollment = ca.enroll({
        enrollmentID: registrar.enrollId,
        enrollmentSecret: registrar.enrollSecret
      }).then( (result) => {
        console.log("Enrolling client")
        // user = new User('admin', config );
        return config.createUser({
          username: 'monitoring_admin',
          mspid: mspId,
          cryptoContent: { privateKeyPEM: result.key.toBytes(), signedCertPEM: result.certificate }
        })
      }).then((user) => {
        //user.setEnrollment(res.key, res.certificate, 'org2')
        //user.setRoles('admin')
        config.setUserContext(user)
        console.log("\"monitoring_admin\" user created. Please upload following certificate via blockchain UI: \n " + req.body.urlRestRoot + "/network/" + req.body.networkId + "/members/certificates")
        console.log(user._signingIdentity._certificate + '\n')
        res.send("User created. Please upload following certificate via blockchain UI: " + user._signingIdentity._certificate)

        //return user
      }).catch((err) => {
        console.error('Failed to enroll and persist admin. Error: ' + err.stack ? err.stack : err);
        throw new Error('Failed to enroll admin');
        });
      }
    })
  })
  // channel = config.getChannel()
});

var chaincodes, peer, channel, marbles_chaincode
router.post('/getchaincodes', function (req, res) {
  config.getUserContext('admin', true)
  // console.log(req)
  // res.send('received chaincode call')
  peer = config.getPeersForOrgOnChannel()[0]._name
  channel = config.getChannel()
  config.queryInstalledChaincodes(peer, true).then( (response) => {
    chaincodes = response
  }).then ( (result) =>  {
    chaincode = _.where( chaincodes.chaincodes, {name: 'simple_contract', version: 'v1'} )[0]
    // console.log(chaincodes)
    res.sendStatus(200)
  });
});

router.post('/chaincode', function (req, res) {
  // console.log("obc request received")
  // console.log(req.body)
  // console.log("args")
  // console.log(req.body.params.ctorMsg.args)
  console.log(req.body)
  if (req.body.method && req.body.method === 'invoke') {
    var transaction_id = config.newTransactionID(true)
    var createMarbleRequest = {
      chaincodeId: chaincode.name,
      chaincodeVersion: chaincode.version,
      txId: transaction_id,
      fcn: req.body.params.ctorMsg.function,
      args: req.body.params.ctorMsg.args
    }
    channel.sendTransactionProposal(createMarbleRequest).then ( (proposalRes) => {
      console.log("sending transaction proposal")
      var proposalResponses = proposalRes[0];
      var proposal = proposalRes[1];
      let isProposalGood = false;
      console.log("proposalResponses")
      console.log(proposalResponses)
      if (proposalResponses && proposalResponses[0].response && proposalResponses[0].response.status === 200) {
          isProposalGood = true;
          console.log('Transaction proposal was good');
        } else {
          console.error('Transaction proposal was bad');
        }
      if (isProposalGood) {
        console.log(util.format('Successfully sent Proposal and received ProposalResponse: Status - %s, message - "%s"', proposalResponses[0].response.status, proposalResponses[0].response.message));
        var promises = []
        var sendPromise = channel.sendTransaction({
          proposalResponses: proposalResponses,
          proposal: proposal
        })
        sendPromise.then( (result) => {
          console.log("transaction result")
          console.log(result)
          res.send(result)
        })
      }
    }).catch ( (err) => {
      console.log(err)
      res.send(err)
    });
  } else { // query
    var assetId = JSON.parse(req.body.params.ctorMsg.args).assetID
    var request = {
        chaincodeId: chaincode.name,
        chaincodeVersion: chaincode.version,
        txId: transaction_id,
        fcn: req.body.params.ctorMsg.function,
        args: req.body.params.ctorMsg.args
      }
      console.log(request)
    channel.queryByChaincode(request).then( (cc_response) => {
      // console.log(cc_response[0].toString())
      console.log(request)
      res.send( cc_response[0].toString() )
    }).catch ( (err) => {
      res.send(err)
    });
  }
   //else if ( req.body.method && req.body.method) === 'query' {
  //   console.log("Chaincode query requested")
  // }
});

module.exports = router;
