var express = require('express'),
    router = express.Router(),
    ctr = require("./controller");


router.get('/home',ctr.main);
router.get('/getPoint1',ctr.getPoint1);

module.exports = router;