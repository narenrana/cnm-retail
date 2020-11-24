var express = require("express");
var app = express();
const path = require("path");
var proxy = require("express-http-proxy");
const numCPUs = require("os").cpus().length;
const cluster = require("cluster");

if (cluster.isMaster) {
  console.log(`Master ${process.pid} is running`);
  // Fork workers.
  for (let i = 0; i < numCPUs; i++) {
    cluster.fork();
  }

  cluster.on("exit", (worker, code, signal) => {
    console.log(`worker ${worker.process.pid} died`);
  });
} else {
  // Workers can share any TCP connection
  // In this case it is an HTTP server
  // app.use(
  //   "/api",
  //   proxy("http://localhost:8080", {
  //     proxyReqOptDecorator: function (proxyReqOpts, srcReq) {
  //       return proxyReqOpts;
  //     },
  //   })
  // );

  app.use(express.static("build"));
  app.get("/", (req, res) => {
    res.sendFile(path.join(__dirname, "build", "index.html"));
  });
  app.get("*", (req, res) =>
    res.sendFile(path.resolve("", "build", "index.html"))
  );

  app.listen(3000);

  console.log(`Worker ${process.pid} started `);
}
