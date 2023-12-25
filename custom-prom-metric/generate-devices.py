# Description: This script generates random data and sends it to the API

#!/usr/bin/env python3

import requests
from time import sleep

url = "http://localhost:8080/devices"

devices = 1000

for i in range(1, devices):
    payload = {
        "id": "device1" + str(i),
        "name": "device1" + str(i),
        "type": "router",
        "hostname": "device1" + str(i) + ".example.com",
        "ip": "192.168.2.1",
        "env": "developement",
    }
    headers = {
        "Content-Type": "application/json",
        "User-Agent": "insomnia/2023.5.8",
    }

    response = requests.request("POST", url, json=payload, headers=headers)
    print(
        "device added to the database, id: "
        + payload["id"]
        + ", type: "
        + payload["type"]
    )
    payload_2 = {
        "id": "device2" + str(i),
        "name": "device2" + str(i),
        "type": "switch",
        "hostname": "device2" + str(i) + ".example.com",
        "ip": "192.168.1.1",
        "env": "developement",
    }

    response = requests.request("POST", url, json=payload_2, headers=headers)
    print(
        "device added to the database, id: "
        + payload_2["id"]
        + ", type: "
        + payload_2["type"]
    )
    sleep(0.1)
