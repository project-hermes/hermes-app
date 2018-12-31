package model_test

import (
	"context"
	"errors"
	"time"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/project-hermes/hermes-app/server/model"
	"github.com/project-hermes/hermes-app/server/wrapper/mock"
	pblatlng "google.golang.org/genproto/googleapis/type/latlng"
)

var _ = Describe("Dive Tests", func() {
	var (
		dive       Dive

		startTime  time.Time
		endTime    time.Time
	)

	AfterEach(func() {
		dive = Dive{}
	})

	Describe("dive implementation", func() {
		var (
			mockCtrl   *gomock.Controller
			mockClient *mock_wrapper.MockDBClientInterface
			mockCollectionRef *mock_wrapper.MockCollectionRefInterface
			mockDocRef *mock_wrapper.MockDocRefInterface
			diveInt DiveInterface
		)

		BeforeEach(func() {
			mockCtrl = gomock.NewController(GinkgoT())
			mockClient = mock_wrapper.NewMockDBClientInterface(mockCtrl)
			mockCollectionRef = mock_wrapper.NewMockCollectionRefInterface(mockCtrl)
			mockDocRef = mock_wrapper.NewMockDocRefInterface(mockCtrl)
			diveInt = NewDiveImplementation(mockClient)
		})

		AfterEach(func() {
			mockCtrl.Finish()
		})

		Context("create dive", func() {
			It("is successful returns nil error", func() {
				ctx := context.Background()
				dive := Dive{
					SensorID: "test_sensor",
				}
				mockClient.EXPECT().Collection(gomock.Eq("dive")).Return(mockCollectionRef).Times(1)
				mockCollectionRef.EXPECT().NewDoc().Return(mockDocRef).Times(1)
				mockDocRef.EXPECT().Set(gomock.Eq(ctx), gomock.Eq(dive)).Return(nil, nil).Times(1)
				err := diveInt.Create(ctx, dive)
				Expect(err).To(BeNil())
			})

			It("doc ref set returns error", func() {
				ctx := context.Background()
				dive := Dive{
					SensorID: "test_sensor",
				}
				mockClient.EXPECT().Collection(gomock.Eq("dive")).Return(mockCollectionRef).Times(1)
				mockCollectionRef.EXPECT().NewDoc().Return(mockDocRef).Times(1)
				mockDocRef.EXPECT().Set(gomock.Eq(ctx), dive).Return(nil, errors.New("ka-boom")).Times(1)
				err := diveInt.Create(ctx, dive)
				Expect(err).To(Equal(errors.New("ka-boom")))
			})
		})
	})

	Context("valid map firestore data to dive", func() {
		BeforeEach(func() {

			endTime = time.Now()
			startTime = endTime.Add(-1 * time.Hour)
			data := map[string]interface{}{
				"sensorId":  "test123",
				"startTime": startTime,
				"endTime":   endTime,
				"startPoint": &pblatlng.LatLng{
					Latitude:  -12.345,
					Longitude: 32.45,
				},
				"endPoint": &pblatlng.LatLng{
					Latitude:  45.67,
					Longitude: -32.45,
				},
				"sensorData": []interface{}{
					map[string]interface{}{
						"depth":       float64(789.987),
						"rawPressure": int64(12345),
						"rawTemp":     int64(23456),
						"temp":        float64(123.45),
						"time":        int64(34567),
					},
				},
			}

			dive = MapDive(data)
		})

		AfterEach(func() {
			dive = Dive{}
			startTime = time.Time{}
			endTime = time.Time{}
		})

		It("sensorId to sensorId", func() {
			Expect(dive.SensorID).To(Equal("test123"))
		})

		It("startPoint LatLng to GeoPoint", func() {
			expected := GeoPoint{Lat: -12.345, Long: 32.45}
			Expect(dive.StartPoint).To(Equal(expected))
		})

		It("endPoint LatLng to GeoPoint", func() {
			expected := GeoPoint{Lat: 45.67, Long: -32.45}
			Expect(dive.EndPoint).To(Equal(expected))
		})

		It("startTime to startTime", func() {
			Expect(dive.StartTime).To(Equal(startTime))
		})

		It("endTime to endTime", func() {
			Expect(dive.EndTime).To(Equal(endTime))
		})

		It("sensorData is parsed and added", func() {
			// skipped
		})
	})
})
