const functions = require('firebase-functions');
const admin = require('firebase-admin');
const moment = require('moment');

admin.initializeApp(functions.config().firebase);
const firestore = admin.firestore();
const settings = { timestampsInSnapshots: true };
firestore.settings(settings);

exports.processDive = functions.https.onRequest((req, res) => {
    if (req.method !== 'POST') {
        res.statusCode = 404;
        res.send("Not Found");
        return;
    }

    if (req.get('content-type') !== 'application/json') {
        res.statusCode = 400;
        res.send("Invalid content type");
        return;
    }

    const body = req.body

    var data = {
        startPoint: new admin.firestore.GeoPoint(body.startLat, body.startLong),
        endPoint: new admin.firestore.GeoPoint(body.endLat, body.endLong),
        sensorId: body.sensorId,
        startTime: moment(body.startTime).toDate(),
        endTime: moment(body.endTime).toDate(),
        stringExample: 'path to GCS data'
    };

    const document = firestore.collection('dive').doc();

    document.set(data).then(() => {
        // write is complete here
        res.send("Hello Firebase");
    }).catch(err => {
        res.send(`Firebase is sad: ${err}`);
    });

});


