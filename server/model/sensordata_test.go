package model_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/project-hermes/hermes-app/server/model"
)

var _ = Describe("SensorData Tests", func() {
	var (
		sensorData SensorData
	)

	AfterEach(func() {
		sensorData = SensorData{}
	})

	Context("valid map firestore data to dive", func() {
		BeforeEach(func() {
			data := map[string]interface{}{
				"depth":       float64(789.987),
				"rawPressure": int64(12345),
				"rawTemp":     int64(23456),
				"temp":        float64(123.45),
				"time":        int64(34567),
			}

			sensorData = MapSensorData(data)
		})

		AfterEach(func() {
			sensorData = SensorData{}
		})

		It("depth as float64", func() {
			Expect(sensorData.Depth).To(Equal(789.987))
		})

		It("raw pressure as int64", func() {
			Expect(sensorData.RawPressure).To(Equal(12345))
		})

		It("raw temp as int64", func() {
			Expect(sensorData.RawTemp).To(Equal(23456))
		})

		It("temp as float64", func() {
			Expect(sensorData.Temp).To(Equal(123.45))
		})

		It("time as int64", func() {
			Expect(sensorData.Time).To(Equal(34567))
		})
	})

	Context("Float types unknown", func() {
		var (
			data map[string]interface{}
		)

		BeforeEach(func() {
			data = map[string]interface{}{
				"rawPressure": int64(12345),
				"rawTemp":     int64(23456),
				"time":        int64(34567),
			}
		})

		AfterEach(func() {
			sensorData = SensorData{}
			data = nil
		})

		It("temp as float64", func() {
			data["temp"] = float64(123.45)
			sensorData = MapSensorData(data)
			Expect(sensorData.Temp).To(Equal(123.45))
		})

		It("temp as int64", func() {
			data["temp"] = int64(12345)
			sensorData = MapSensorData(data)
			Expect(sensorData.Temp).To(Equal(float64(12345)))
		})

		It("temp as int", func() {
			data["temp"] = int(12345)
			sensorData = MapSensorData(data)
			Expect(sensorData.Temp).To(Equal(float64(12345)))
		})

		It("temp as string is zero", func() {
			data["temp"] = "123.45"
			sensorData = MapSensorData(data)
			Expect(sensorData.Temp).To(Equal(0.0))
		})

		It("depth as float64", func() {
			data["depth"] = float64(123.45)
			sensorData = MapSensorData(data)
			Expect(sensorData.Depth).To(Equal(123.45))
		})

		It("depth as int64", func() {
			data["depth"] = int64(12345)
			sensorData = MapSensorData(data)
			Expect(sensorData.Depth).To(Equal(float64(12345)))
		})

		It("depth as int", func() {
			data["depth"] = int(12345)
			sensorData = MapSensorData(data)
			Expect(sensorData.Depth).To(Equal(float64(12345)))
		})

		It("depth as string is zero", func() {
			data["depth"] = "123.45"
			sensorData = MapSensorData(data)
			Expect(sensorData.Depth).To(Equal(0.0))
		})
	})
})
