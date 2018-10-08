const os = require('os');
const path = require('path');

const StorageController = function (storageInt, fs) {

    this.storageInt = storageInt;

    const readFile = function* (storageObject) {
        // const filePath = bucketObject.name;
        const parsedPath = path.parse(storageObject.name);
        const fileName = parsedPath.base;
    
        const bucket = storageInt.bucket(storageObject.bucket);
        const file = bucket.file(storageObject.name);

        const tempLocalDir = path.join(os.tmpdir(), parsedPath.dir);
        const tempLocalFile = path.join(tempLocalDir, fileName);

        const fileDownload = yield file.download({ destination: tempLocalFile});

        return fileDownload;
    };

    return {
        readFile
    };
};

module.exports = StorageController;