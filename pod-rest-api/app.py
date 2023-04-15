from flask import Flask, jsonify
import subprocess
import json

app = Flask(__name__)


def get_pods_json():
    try:
        process = subprocess.Popen(
            ["kubectl", "get", "pods", "-o", "json"], stdout=subprocess.PIPE
        )
        stdout, stderr = process.communicate()
        if process.returncode != 0:
            print(f"Error running kubectl command: {stderr}")
            return None
        pods_json = json.loads(stdout)
        return pods_json
    except Exception as e:
        print(f"Error getting pods in JSON format: {e}")
        return None


@app.route("/api/v1/pods")
def get_pods():
    pods_json = get_pods_json()
    if pods_json is None:
        return jsonify({"error": "Error getting pods in JSON format"}), 500
    pods = pods_json["items"]
    return jsonify(pods), 200


if __name__ == "__main__":
    app.run(port=5000, debug=True)
