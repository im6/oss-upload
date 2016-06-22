"use strict";
var util = require("./converter"),
    requester = require("./initiator");

module.exports = {
    getPoint1: function(obj){
        var obj2 = util.createReqJson("GET", "/{moduleName}", obj);
        return requester.one(obj2);
    }
};
