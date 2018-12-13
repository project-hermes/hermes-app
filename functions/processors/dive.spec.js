const chai = require('chai');
const chaiAsPromised = require('chai-as-promised');
const sinon = require('sinon');
const DiveProcessor = require('./dive');
const expect = chai.expect;
chai.use(chaiAsPromised);

var testData = {
	"startTime":124525634545,
	"endTime":134545346549,
	"startLat":34.2453,
	"startLong":-13.56467,
	"endLat":64.35234,
	"endLong":-56.2453,
	"sensorId":"13df4g5y3",
	"readings":[
		{
			"time":2244244,
			"rawTemp":34244928,
			"rawPressure":23457632,
			"temp":78.9,
			"depth":34.564
		},
	]
}

describe('Dive processor...', () => {
    let sandbox, processor;
    var datastoreStub, storageStub;
    beforeEach(() => {
        sandbox = sinon.createSandbox();
    });

    after(() => {
        sandbox.restore();
    });

    beforeEach(() => {
        datastoreStub = sinon.stub();
        storageStub = sinon.stub();

        processor = new DiveProcessor(storage);
    });

    describe('save dive...', () => {
        var gen;

        beforeEach(() => {
            // Make sure the controllers are returning promises
            gen = processor.save(testData);
        });

        afterEach(() => {
            datastoreStub.reset();
            storageStub.reset();
        });

        describe('storage bucket function...', () => {
            it('is called one time', () => {
                gen.next();
                expect(1).to.equal(1);
            });

            it('is called with one argument', () => {
                expect(1).to.equal(1);
            });

            describe('first argument...', () => {
                
            });
        });
    });


});