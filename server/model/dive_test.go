package model_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/project-hermes/hermes-app/server/model"
	pblatlng "google.golang.org/genproto/googleapis/type/latlng"
)

var _ = Describe("Dive Tests", func() {
	var (
		dive      Dive
		startTime time.Time
		endTime   time.Time
	)

	AfterEach(func() {
		dive = Dive{}
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
