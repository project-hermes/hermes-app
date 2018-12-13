const admin = require('firebase-admin');
const moment = require('moment');

const DiveProcessor = function (storageController) {

    this.storageController = storageController;

    const save = function* (dive, firestoreDoc) {
        const data = {
            startPoint: new admin.firestore.GeoPoint(dive.startLat, dive.startLong),
            endPoint: new admin.firestore.GeoPoint(dive.endLat, dive.endLong),
            sensorId: dive.sensorId,
            startTime: moment(dive.startTime).toDate(),
            endTime: moment(dive.endTime).toDate(),
            stringExample: 'path to GCS data'
        };

        console.log(moment().toDate());

    };

    return {
        save
    };
};

module.exports = DiveProcessor;