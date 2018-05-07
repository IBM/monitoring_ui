/*****************************************************************************
Copyright (c) 2016 IBM Corporation and other Contributors.


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and limitations under the License.


Contributors:

Alex Nguyen - Initial Contribution
*****************************************************************************/
import { combineReducers } from 'redux'
import {blockchain} from './BlockchainReducer.js'
import {configuration} from './ConfigurationReducer.js'
import {chaincode} from './ChaincodeReducer.js'
import {app} from './AppReducer'

import {
  modelReducer,
  formReducer
} from 'react-redux-form';

const initialConfigurationState = {
  urlRestRoot: "https://ibmblockchain-starter.ng.bluemix.net/api/v1",
  chaincodeId: "simple_contract",
  secureContext: "<username>@<email_domain>",
  blocksPerPage: "10"
};

const initialChaincodeOpsFormState = {
  tabOne:{
    selectedFn: "firstOne",
    selectFns: ['firstOne','secondOne','thirdOne']
  }
}

/**
Combines all other reducers into one reducer called the root reducer. We will be using the root
reducer when creating the redux store.
**/
const rootReducer = combineReducers({
  blockchain,
  configuration,
  //obcConfiguration is the model that deals with any configuration related to obc
  obcConfiguration: modelReducer('obcConfiguration', initialConfigurationState),
  chaincodeOpsForm: modelReducer('chaincodeOpsForm', initialChaincodeOpsFormState),
  obcConfigurationForm: formReducer('obcConfiguration'),
  //
  chaincode,
  app

})

export default rootReducer
