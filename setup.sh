#!/bin/bash
# apt-get update -y
# apt-get install -y git curl python build-essential
# install nvm
curl -o- https://raw.githubusercontent.com/creationix/nvm/v0.33.11/install.sh | bash
# add nvm values to .bash_profile
cat .nvm_profile >> ~/.bash_profile
. ~/.bash_profile
nvm install 8.9.0
#nvm use 8.9.0
npm install
npm install --prefix react-backend/
