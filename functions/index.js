'use strict';

const functions = require('firebase-functions');
const admin = require('firebase-admin');
admin.initializeApp();
const path = require('path');
const os = require('os');
const convert = require('xml-js');

exports.processDivelog = functions.storage.object().onFinalize(async object => {
    const fileBucket = object.bucket;
    const filePath = object.name;
    const contentType = object.contentType;
    const metageneration = object.metageneration;

    if (!contentType.endsWith('xml')) {
        return console.log(
            filePath +
                ' with ' +
                contentType +
                'is not an xml file. We will not process this.'
        );
    }
    const fileName = path.basename(filePath);

    const bucket = admin.storage().bucket(fileBucket);
    const tempFilePath = path.join(os.tmpdir(), fileName);
    const metadata = {
        contentType: contentType
    };
    var data = await bucket.file(filePath).download();
    console.log('Divelog loaded');
    var diveData = convert.xml2js(data.toString(), {
        compact: true,
        spaces: 4
    });
    console.log('Converted dive log');
    for (var d in diveData.dives.dive) {
        console.log(
            diveData.dives.dive[d].date._text,
            diveData.dives.dive[d].duration._text,
            diveData.dives.dive[d].maxDepth._text,
            diveData.dives.dive[d].averageDepth._text,
            diveData.dives.dive[d].tempHigh._text,
            diveData.dives.dive[d].tempLow._text,
            diveData.dives.dive[d].site.lat,
            diveData.dives.dive[d].site.lon,
            diveData.dives.dive[d].gear.item.name._text,
            diveData.dives.dive[d].gear.item.serial._text
        );
    }
});
