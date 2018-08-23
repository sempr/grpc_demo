var messages = require('./math_pb');
var services = require("./math_grpc_pb");

var grpc = require("grpc");

function calc_factor(input) {
    var i = 2;
    var origin = input;
    var res = [];
    while(i * i <= origin) {
        while (origin % i === 0) {
            res.push(i);
            origin /= i;
        }
        i ++;
    }
    if (origin > 1) {
        res.push(origin)
    }
    return res;
}

function sqrt(call, callback) {
    var val = call.request.getValue();
    var reply = new messages.SqrtResponse();
    if (val < 0) {
        var err = {
            "code": grpc.status.INVALID_ARGUMENT,
            "details": "不能为负数开方"
        }
        return callback(err);
    } else {
        reply.setValue(Math.sqrt(call.request.getValue()));
        return callback(null, reply);
    }
    callback(null, reply);
}

function stat(call, callback) {
    var sum = 0, count = 0;
    call.on("data", function(data) {
        sum += data | 0;
        count ++;
    })
    call.on("end", function() {
        var reply = new messages.StatResponse();
        reply.setSum(sum);
        reply.setCount(count);
        callback(null, reply);
    })
}

function factor(call, callback) {
    var val = call.request.getValue();
    var res = calc_factor(val);
    for (let i = 0; i != res.length; i ++) {
        var reply = new messages.FactorResponse();
        reply.setValue(res[i]);
        call.write(reply);
    }
    call.end();
}

function main() {
    var server = new grpc.Server();
    var funcs = {
        sqrt: sqrt,
        stat: stat,
        factor: factor
    };
    server.addService(services.MathService, funcs);
    server.bind("127.0.0.1:50000", grpc.ServerCredentials.createInsecure());
    server.start();
}

main();