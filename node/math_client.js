var messages = require('./math_pb');
var services = require("./math_grpc_pb");

var grpc = require("grpc");

var client = new services.MathClient("127.0.0.1:50000", grpc.credentials.createInsecure());

var callSqrt = function(input) {
    var request = new messages.SqrtRequest();
    request.setValue(input);
    var promise = new Promise(function(resolve, reject) {
        client.sqrt(request, function(err, response) {
            if (!err) {
                resolve(response.getValue());
            } else {
                console.error(err.message);
                reject(err);
            }
        })
    });
    return promise;
}

var callStat = function(input) {
    var promise = new Promise(function(resolve, reject) {
        var call = client.stat(function(error, stats) {
            if (error) {
                console.error(error);
                reject(error);
            }
            var tmpObj = {
                sum: stats.getSum(),
                count: stats.getCount()
            }
            resolve(tmpObj);
        });
        for (let i = 0; i != input.length; i ++) {
            var request = new messages.StatRequest();
            request.setValue(input[i]);
            call.write(request);
        }
        call.end();
    });
    return promise;
}

var callFactor = function(input) {
    var request = new messages.FactorRequest();
    request.setValue(input);
    var promise = new Promise(function(resolve, reject) {
        var call = client.factor(request);
        var res = [];
        call.on("data", function(data) {
            res.push(data.getValue());
        });
        call.on("end", function() {
            resolve(res);
        });
        call.on("status", function(status) {
            // 错误处理在此处，需要分析status对象
            // if (status.code) {
                // reject(status);
            // }
        });
    })
    return promise;
}

function main() {
    callSqrt(100).then(function(res){ 
        console.log(res); 
    }, function() {

    });

    callSqrt(-1).then(function(res) {
        console.log(res);
    }, function(){

    });

    callStat([1,2,3,4]).then(function(res) {
        console.log(`sum=${res.sum}, count=${res.count}`);
    }, function() {

    });

    callFactor(100).then(function(res) {
        console.log(res);
    }, function() {

    });
}

main();