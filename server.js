var path = require('path');
var express = require('express');

var app = express();
var server_port = process.env.OPENSHIFT_NODEJS_PORT || 8080;
var server_ip_address = process.env.OPENSHIFT_NODEJS_IP || '127.0.0.1';

//var env = process.env.ENVIRONMENT;
var env = 'dev';

app.set('x-powered-by', false);
app.use(express.static(path.join(__dirname, 'public')));

app.use('/', require('./server/main/route'));

app.listen(server_port, server_ip_address, function () {
    console.log( "Listening on " + server_ip_address + ", server_port " + server_port + '...' );
});