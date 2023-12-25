# Description: This script deletes all the devices from the database

#!/usr/bin/env python3

import requests
from time import sleep

for i in range(1, 1000):
    payload = {"id": "device1" + str(i)}
    headers = {"Content-Type": "application/json"}

    routerURL = "http://localhost:8080/devices/" + "device1" + str(i)
    response = requests.request("DELETE", routerURL, json=payload, headers=headers)
    print("device removed to the database, id: " + payload["id"])

    switchURL = "http://localhost:8080/devices/" + "device2" + str(i)
    payload_2 = {"id": "device2" + str(i)}
    response = requests.request("DELETE", switchURL, json=payload_2, headers=headers)
    print("device removed to the database, id: " + payload_2["id"])
print("Removed all devices from the database")
