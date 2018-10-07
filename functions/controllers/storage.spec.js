const chai = require('chai');
const chaiAsPromised = require('chai-as-promised');
const sinon = require('sinon');
const StorageController = require('./storage');
const expect = chai.expect;
chai.use(chaiAsPromised);

describe('Storage controller...', () => {
    var controller, requestStub, info;
    before(() => {

    });

    after(() => {

    });

    beforeEach(() => {
        requestStub = sinon.stub();
        controller = new StorageController();
    });

    describe('read file...', () => {
        beforeEach(() => {
            info = {
                from: 'from@test.com',
                to: 'test@test.com',
                subject: 'This is a Test',
                html: '<body>Test</body>',
                text: 'Test'
            };
            controller.readFile();
        });

        afterEach(() => {
            requestStub.reset();
        });

        describe('request...', () => {
            it('is called one time', () => {
                expect(requestStub.callCount).to.equal(1);
            });

            // it('is called with one argument', () => {
            //     expect(requestStub.firstCall.args.length).to.equal(1);
            // });

            describe('first argument...', () => {
                var arg;
                beforeEach(() => {
                    arg = requestStub.firstCall.args[0];
                });

                // it('one element in header', () => {
                //     let headers = arg.headers;
                //     expect(Object.keys(headers).length).to.equal(1);
                // });

                // it('contains authorization header', () => {
                //     expect(arg.headers).to.haveOwnProperty('Authorization');
                // });

                // it('authorization value is a base64 encoded string', () => {
                //     let auth = new Buffer(`api:key12345`).toString('base64');
                //     let basic = `Basic ${auth}`;
                //     expect(arg.headers['Authorization']).to.equal(basic);
                // });

                // it('method is POST', () => {
                //     expect(arg.method).to.equal('POST');
                // });

                // it('uri is mailgun v3 with domain and messages appended', () => {
                //     let expectedUri = 'https://api.mailgun.net/v3/test.flyworkouts.com/messages';
                //     expect(arg.uri).to.equal(expectedUri);
                // });

            });
        });

    });


});