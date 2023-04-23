from kubernetes import client, config, watch
import websockets
import asyncio
import json
import logging

async def send_events(websocket):
    config.load_incluster_config()
    v1 = client.CoreV1Api()
    w = watch.Watch()
    for event in w.stream(v1.list_pod_for_all_namespaces):
        cus_event={
            "event_type": event['type'],
            "object": {
                "resource": event['object'].metadata.name,
                "namespace": event['object'].metadata.namespace,
        }
        }
        logging.info("Event: %s %s %s" % (cus_event['event_type'], cus_event['object']['namespace'], cus_event['object']['resource']))
        await websocket.send(json.dumps(cus_event))


async def main(websocket, path):
    await send_events(websocket)

start_server = websockets.serve(main, "0.0.0.0", 8080)
print("Server starting on port 8080")
asyncio.get_event_loop().run_until_complete(start_server)
print("Server started on port 8080")
asyncio.get_event_loop().run_forever()


