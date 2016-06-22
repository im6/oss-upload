"use strict";
var _ = require('lodash'),
    request = require('request');

var privateFn = {
    urlPathMap: function(rawUrl, pathObj){
        var result = rawUrl;
        for(var onep in pathObj){
            if(pathObj.hasOwnProperty(onep)){
                result =  result.replace("{" + onep + "}", pathObj[onep]);
            }
        }
        return result;
    }
};

module.exports = {
    createReqJson: function(method, url, obj){
        var reqObj = {};
        for(var onep in obj){
            if(obj.hasOwnProperty(onep)){
                reqObj[onep] = obj[onep];
            }
        }

        // makesure body is stringified json
        if(typeof reqObj["body"] !== "undefined" && typeof reqObj["body"] != "string" ){
            reqObj["body"] = JSON.stringify(reqObj["body"]);
        }

        var result0 = {
            method: method,
            baseUrl: process.env.baseUrl,
            url: privateFn.urlPathMap(url, obj.pathVar),
            // if you want to add different default value in the obj, put it below
            timeout: 20000
        };

        var result = _.assign(result0, reqObj);

        return result;
    }
};