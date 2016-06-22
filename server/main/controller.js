var path = require('path');
var appDir = path.dirname(require.main.filename);
var list = require('../request/list');

module.exports = {
    main: function(req, res, next){
        res.sendFile('/public/index.html',{ root: appDir });
    },

    getPoint1: function(req, res, next){
        var obj = {
            pathVar:{
                moduleName: 'react'
            }
        };

        list.getPoint1(obj).then(function(data){
            res.json(data);
        });

    }
};