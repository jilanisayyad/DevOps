const express = require("express");
const { spawn } = require("child_process");
const bodyParser = require("body-parser");

const app = express();
const port = 5000;

app.use(bodyParser.json());

function getPodsJSON(callback) {
  const cmd = spawn("kubectl", ["get", "pods", "-o", "json"]);
  let stdout = "";
  let stderr = "";

  cmd.stdout.on("data", (data) => {
    stdout += data;
  });

  cmd.stderr.on("data", (data) => {
    stderr += data;
  });

  cmd.on("close", (code) => {
    if (code !== 0) {
      console.error(`Error running kubectl command: ${stderr}`);
      callback(null, stderr);
    } else {
      try {
        const podsJSON = JSON.parse(stdout);
        callback(podsJSON, null);
      } catch (err) {
        console.error(`Error parsing JSON: ${err}`);
        callback(null, err);
      }
    }
  });
}

app.get("/api/v1/pods", (req, res) => {
  getPodsJSON((podsJSON, err) => {
    if (err) {
      res.status(500).json({ error: "Error getting pods in JSON format" });
    } else {
      const pods = podsJSON.items;
      res.json(pods);
    }
  });
});

app.listen(port, () => {
  console.log(`Server is running on port ${port}`);
});
