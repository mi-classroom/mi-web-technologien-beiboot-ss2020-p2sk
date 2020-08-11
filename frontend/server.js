/*
nodejs server für das frontend
served via basic http
*/

var fs = require('fs'),
    http = require('http'),
    static = require('node-static');

const hostname = '0.0.0.0';
const port = 8081;

var file = new static.Server('./public');

http.createServer(function (request, response) {
    request.addListener('end', function () {
        file.serve(request, response);
    }).resume();
}).listen(port, hostname);