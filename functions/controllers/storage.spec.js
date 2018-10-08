const chai = require('chai');
const chaiAsPromised = require('chai-as-promised');
const os = require('os');
const path = require('path');
const sinon = require('sinon');
const StorageController = require('./storage');
const { Storage } = require('@google-cloud/storage');
const expect = chai.expect;
chai.use(chaiAsPromised);

describe('Storage controller...', () => {
    let sandbox;
    var controller, storageObj, bucketStub, fileStub, downloadStub, gen;
    beforeEach(() => {
        sandbox = sinon.createSandbox();
    });

    after(() => {
        sandbox.restore();
    });

    beforeEach(() => {
        storage = new Storage();

        downloadStub = sinon.stub();
        fileStub = sinon.stub().returns({
            download: downloadStub
        });
        bucketStub = sinon.stub(storage, 'bucket').returns({
            file: fileStub
        });

        controller = new StorageController(storage);
    });

    describe('read file...', () => {
        beforeEach(() => {
            storageObj = {
                name: '/sub-dir/test-file.json',
                bucket: 'test-bucket',
            };
            gen = controller.readFile(storageObj);
            gen.next();
        });

        afterEach(() => {
            bucketStub.reset();
            fileStub.reset();
            downloadStub.reset();
        });

        describe('storage bucket function...', () => {
            it('is called one time', () => {
                expect(bucketStub.callCount).to.equal(1);
            });

            it('is called with one argument', () => {
                expect(bucketStub.firstCall.args.length).to.equal(1);
            });

            describe('first argument...', () => {
                var arg;
                beforeEach(() => {
                    arg = bucketStub.firstCall.args[0];
                });

                it('storage object bucket name', () => {
                    expect(arg).to.equal('test-bucket');
                });
            });
        });

        describe(`bucket's file function...`, () => {
            it('is called one time', () => {
                expect(fileStub.callCount).to.equal(1);
            });

            it('is called with one argument', () => {
                expect(fileStub.firstCall.args.length).to.equal(1);
            });

            describe('first argument...', () => {
                var arg;
                beforeEach(() => {
                    arg = fileStub.firstCall.args[0];
                });

                it('storage object bucket name', () => {
                    expect(arg).to.equal('/sub-dir/test-file.json');
                });
            });
        });

        describe(`file's download function...`, () => {
            beforeEach(() => {
            });

            it('is called one time', () => {
                expect(downloadStub.callCount).to.equal(1);
            });

            it('is called with one argument', () => {
                expect(downloadStub.firstCall.args.length).to.equal(1);
            });

            describe('first argument...', () => {
                var arg;
                beforeEach(() => {
                    arg = downloadStub.firstCall.args[0];
                });

                it('storage object bucket name', () => {
                    const tempLocalDir = path.join(os.tmpdir(), 'sub-dir');
                    const tempLocalFile = path.join(tempLocalDir, 'test-file.json');
                    expect(arg['destination']).to.equal(tempLocalFile);
                });
            });
        });

    });


});