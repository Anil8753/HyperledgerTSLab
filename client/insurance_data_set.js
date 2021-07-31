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
    set_insurance_data,
}

async function set_insurance_data(record) {
    try {
        const regNumber = record['regNumber'];
        const uINNumber  = record['UINNumber'];
        const policyNumber  = record['PolicyNumber'];
        const insuredNameAndAddress = record['InsuredNameAndAddress'];
        const contactNumber = record['ContactNumber'];
        const emailId = record['EmailId'];
        const periodOfCover = record['PeriodOfCover'];
        const premiumDetails = record['PremiumDetails'];

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
            UIN Number: UIN240720211076
            Policy Number: BAJAL11012023421
            Insured Name &amp; Address: Somajiguda, Hyderabad
            Contact Number: 9979788665
            Email Id: vehiclelifecycle2@gmail.com
            Period of Cover: 15/05/2021 to 14/03/2022
            Premium Details: Rs 17,000
        */
        await contract.submitTransaction(
            'SetInsuranceData',     // function
            regNumber,              //'TS09AE0200',                   //Vehicle No - identifier
            uINNumber,              //'UIN240720211076',              // UIN Number
            policyNumber,           //'BAJAL11012023421',             // Policy Number
            insuredNameAndAddress,  //'Somajiguda, Hyderabad',        // Insured Name &amp; Address
            contactNumber,          //'9979788665',                   // Contact Number
            emailId,                //'vehiclelifecycle2@gmail.com',  // Email Id
            periodOfCover,          //'15/05/2021 to 14/03/2022',     // Period of Cover
            premiumDetails,         //'Rs 17,000'                     // Premium Details
        );

        console.log('Transaction has been submitted');

        // Disconnect from the gateway.
        await gateway.disconnect();

    } catch (error) {
        console.error(`Failed to submit transaction: ${error}`);
        return error.toString();
    }
}
