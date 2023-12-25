package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type Device struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Hostname string `json:"hostname"`
	IP       string `json:"ip"`
}

type DeviceList struct {
	Devices []Device `json:"devices"`
}

func StoreDevice(client *redis.Client, device Device, c *gin.Context) {
	log.Println("Storing device: ", device)
	var deviceJSON []byte
	deviceJSON, err := json.Marshal(device)
	if err != nil {
		log.Panic(err)
	}
	log.Println(string(deviceJSON))
	err = client.Set(context.Background(), device.ID, deviceJSON, 0).Err()
	if err != nil {
		log.Panic(err)
	}
	log.Println("Device stored")
}

func GetDevice(client *redis.Client, id string, c *gin.Context) Device {
	log.Println("Getting device: ", id)
	val, err := client.Get(context.Background(), id).Result()
	if err != nil {
		log.Panic(err)
	}
	log.Println("Device found: ", id)
	var device Device
	err = json.Unmarshal([]byte(val), &device)
	if err != nil {
		log.Panic(err)
	}
	log.Println(device)
	return device
}

func DeleteDevice(client *redis.Client, id string, c *gin.Context) {
	log.Println("Deleting device: ", id)
	err := client.Del(context.Background(), id).Err()
	if err != nil {
		log.Panic(err)
	}
	log.Println("Device deleted: ", id)
}

func GetDevices(client *redis.Client, w http.ResponseWriter, r *http.Request) DeviceList {
	log.Println("Getting all devices")
	val, err := client.Keys(context.Background(), "*").Result()
	if err != nil {
		log.Panic(err)
	}
	log.Println("Devices found: ", val)
	var deviceList DeviceList
	for _, id := range val {
		deviceList.Devices = append(deviceList.Devices, GetDevice(client, id, nil))
	}
	log.Println(deviceList)
	return deviceList
}

func ConnectToRedis() *redis.Client {
	if os.Getenv("REDIS_HOST") == "" {
		log.Panic("REDIS_HOST is not set")
	}
	redisHost := os.Getenv("REDIS_HOST")
	client := redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: "",
		DB:       0,
	})
	CheckConnected := client.Ping(context.Background())
	if CheckConnected.Err() != nil {
		log.Panic(CheckConnected.Err())
	}
	log.Println("Connected to Redis")
	return client
}
