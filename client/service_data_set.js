/*
 * Copyright IBM Corp. All Rights Reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

'use strict';

const { Gateway, Wallets } = require('fabric-network');
const fs = require('fs');
const path = require('path');

module.exports =
{
    set_service_data,
}

async function set_service_data(record) {
    try {
        const regNumber = record['regNumber'];
        const chassisNumber = record['chassisNumber'];
        const engineNumber=record['engineNumber'];
        const monthYearOfMfg= record['monthYearOfMfg'];
        const serviceDetails= record['serviceDetails'];

        // load the network configuration
        const ccpPath = path.resolve(__dirname, '..', '..', 'test-network', 'organizations', 'peerOrganizations', 'org1.example.com', 'connection-org1.json');
        let ccp = JSON.parse(fs.readFileSync(ccpPath, 'utf8'));

        // Create a new file system based wallet for managing identities.
        const walletPath = path.join(process.cwd(), 'wallet');
        const wallet = await Wallets.newFileSystemWallet(walletPath);
        console.log(`Wallet path: ${walletPath}`);

        // Check to see if we've already enrolled the user.
        const identity = await wallet.get('appUser');
        if (!identity) {
            console.log('An identity for the user "appUser" does not exist in the wallet');
            console.log('Run the registerUser.js application before retrying');
            return;
        }

        // Create a new gateway for connecting to our peer node.
        const gateway = new Gateway();
        await gateway.connect(ccp, { wallet, identity: 'appUser', discovery: { enabled: true, asLocalhost: true } });

        // Get the network (channel) our contract is deployed to.
        const network = await gateway.getNetwork('mychannel');

        // Get the contract from the network.
        const contract = network.getContract('fabcar');

        // Submit the specified transaction.
        /*
            Vehicle No: TS09AE0200
            Chassis Number: MAKGM553EH624052017
            Engine Number: L15Z22022017
            Mth.Yr. of Mfg:2020
            Service Details: Free Service, Oil Change,
        */
        await contract.submitTransaction(
            'SetServiceData',           // function
            regNumber,                  //'TS09AE0200',           // Vehicle No
            chassisNumber,              //'MAKGM553EH624052017',  // Chassis Number
            engineNumber,               //'L15Z22022017',         // Engine Number
            monthYearOfMfg,             //'2020'                  // Mth.Yr. of Mfg
            serviceDetails              // Oil Change             // Service Details
        );

        console.log('Transaction has been submitted');

        // Disconnect from the gateway.
        await gateway.disconnect();

    } catch (error) {
        console.error(`Failed to submit transaction: ${error}`);
        return error.toString();
    }
}

