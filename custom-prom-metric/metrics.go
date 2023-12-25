package main

import (
	"log"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var TotalDevices = promauto.NewGauge(prometheus.GaugeOpts{
	Name:        "total_devices",
	ConstLabels: TotalDevicesLabels,
	Help:        "Total number of devices",
})

var TotalDevicesLabels = prometheus.Labels{
	"app": "device-manager",
	"env": os.Getenv("ENVIRONMENT"),
}

func GetDeviceCounts() {
	if os.Getenv("ENVIRONMENT") == "" {
		log.Fatal("ENVIRONMENT environment variable not set")
	}
	go func() {
		for {
			client := ConnectToRedis()
			Devices := GetDevices(client, nil, nil)
			TotalDevices.Set(float64(len(Devices.Devices)))
			time.Sleep(5 * time.Second)
			log.Println("Total devices: ", len(Devices.Devices))
		}
	}()
}

var devicesByType = promauto.NewGaugeVec(prometheus.GaugeOpts{
	Name: "devices_by_type",
	Help: "Number of devices by type",
}, []string{"type"})

func GetDeviceCountsByType() {
	if os.Getenv("ENVIRONMENT") == "" {
		log.Fatal("ENVIRONMENT environment variable not set")
	}
	go func() {
		for {
			client := ConnectToRedis()
			Devices := GetDevices(client, nil, nil)
			DevicesByType := make(map[string]int)
			for _, device := range Devices.Devices {
				DevicesByType[device.Type]++
			}
			for deviceType, count := range DevicesByType {
				devicesByType.WithLabelValues(deviceType).Set(float64(count))
				log.Println("Devices of type", deviceType, ":", count)
			}
			time.Sleep(5 * time.Second)
		}
	}()
}

func GetMetrics() {
	GetDeviceCounts()
	GetDeviceCountsByType()
}
