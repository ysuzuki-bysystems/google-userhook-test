#!/usr/bin/env node

import * as http from "node:http";

const port = Number.parseInt(process.env["PORT"] ?? "8000", 10);
const server = http.createServer((req, res) => {
  console.log(req.url, req.headers);

  res.writeHead(204);
  res.end();
});
server.listen(port);
