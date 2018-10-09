'use strict';

const functions = require('firebase-functions');
const admin = require('firebase-admin');
admin.initializeApp();
const path = require('path');
const os = require('os');
const fs = require('fs');

exports.processDivelog = functions.storage.object().onFinalize(async object => {
    const fileBucket = object.bucket;
    const filePath = object.name;
    const contentType = object.contentType;
    const metageneration = object.metageneration;

    if (!contentType.startsWith('xml')) {
        return console.log(
            filePath + ' is not an xml file. We will not process this.'
        );
    }
    const fileName = path.basename(filePath);

    const bucket = admin.storage().bucket(fileBucket);
    const tempFilePath = path.join(os.tmpdir(), fileName);
    const metadata = {
        contentType: contentType
    };
    await bucket.file(filePath).download({destination: tempFilePath});
    console.log('Divelog downloaded locally to', tempFilePath);
    return fs.unlinkSync(tempFilePath);
});
