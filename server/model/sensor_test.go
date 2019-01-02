package model_test

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/project-hermes/hermes-app/server/model"
	"github.com/project-hermes/hermes-app/server/wrapper/mock"

	// . "github.com/project-hermes/hermes-app/server/model"
)

var _ = Describe("Sensor Tests", func() {
	var (
		mockCtrl   *gomock.Controller
		mockClient *mock_wrapper.MockDBClientInterface
		mockCollectionRef *mock_wrapper.MockCollectionRefInterface
		mockDocRef *mock_wrapper.MockDocRefInterface
		sensorInt SensorInterface
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockClient = mock_wrapper.NewMockDBClientInterface(mockCtrl)
		mockCollectionRef = mock_wrapper.NewMockCollectionRefInterface(mockCtrl)
		mockDocRef = mock_wrapper.NewMockDocRefInterface(mockCtrl)
		sensorInt = NewSensorImplementation(mockClient)
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("create dive", func() {
		It("is successful", func() {
			ctx := context.Background()
			inputSensor := InputSensor{
				SensorID: "sensor_id",
				Name: "test_sensor",
				Model: "test_model",
				Type: "test_type",
			}

			expectedSensor := Sensor{
				ID: "sensor_id",
				Name: "test_sensor",
				Model: "test_model",
				Type: "test_type",
				Status: SensorStatusActive,
			}
			mockClient.EXPECT().Collection(gomock.Eq("sensor")).Return(mockCollectionRef).Times(1)
			mockCollectionRef.EXPECT().Doc(inputSensor.SensorID).Return(mockDocRef).Times(1)
			mockDocRef.EXPECT().Set(gomock.Eq(ctx), gomock.Eq(expectedSensor)).Return(nil, nil).Times(1)

			sensor, err := sensorInt.Create(ctx, inputSensor)
			Expect(sensor).To(Equal(expectedSensor))
			Expect(err).To(BeNil())
		})

		It("doc ref set returns error", func() {
			ctx := context.Background()
			inputSensor := InputSensor{
				SensorID: "sensor_id",
				Name: "test_sensor",
				Model: "test_model",
				Type: "test_type",
			}

			expectedSensor := Sensor{
				ID: "sensor_id",
				Name: "test_sensor",
				Model: "test_model",
				Type: "test_type",
				Status: SensorStatusActive,
			}
			mockClient.EXPECT().Collection(gomock.Eq("sensor")).Return(mockCollectionRef).Times(1)
			mockCollectionRef.EXPECT().Doc(inputSensor.SensorID).Return(mockDocRef).Times(1)
			mockDocRef.EXPECT().Set(gomock.Eq(ctx), expectedSensor).Return(nil, errors.New("ka-boom")).Times(1)
			sensor, err := sensorInt.Create(ctx, inputSensor)

			Expect(sensor).To(Equal(Sensor{}))
			Expect(err).To(Equal(errors.New("ka-boom")))
		})
	})

	Context("get sensors by ids", func() {

	})
})
