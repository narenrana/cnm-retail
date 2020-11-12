var express = require("express");
var app = express();
const path = require("path");
const { proxy } = require("http-proxy-middleware");

module.exports = function (app) {
  app.use(proxy("/api", { target: "http://localhost:8080" }));
  // whatever that http-proxy-middleware accepts
};

//setting middleware
//app.use(express.static(__dirname + "build")); //Serves resources from public folder
app.use(express.static("build"));
app.get("/", (req, res) => {
  res.sendFile(path.join(__dirname, "build", "index.html"));
});
var server = app.listen(3000);
