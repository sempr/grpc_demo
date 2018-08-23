// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var math_pb = require('./math_pb.js');

function serialize_FactorRequest(arg) {
  if (!(arg instanceof math_pb.FactorRequest)) {
    throw new Error('Expected argument of type FactorRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_FactorRequest(buffer_arg) {
  return math_pb.FactorRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_FactorResponse(arg) {
  if (!(arg instanceof math_pb.FactorResponse)) {
    throw new Error('Expected argument of type FactorResponse');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_FactorResponse(buffer_arg) {
  return math_pb.FactorResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_SqrtRequest(arg) {
  if (!(arg instanceof math_pb.SqrtRequest)) {
    throw new Error('Expected argument of type SqrtRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_SqrtRequest(buffer_arg) {
  return math_pb.SqrtRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_SqrtResponse(arg) {
  if (!(arg instanceof math_pb.SqrtResponse)) {
    throw new Error('Expected argument of type SqrtResponse');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_SqrtResponse(buffer_arg) {
  return math_pb.SqrtResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_StatRequest(arg) {
  if (!(arg instanceof math_pb.StatRequest)) {
    throw new Error('Expected argument of type StatRequest');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_StatRequest(buffer_arg) {
  return math_pb.StatRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_StatResponse(arg) {
  if (!(arg instanceof math_pb.StatResponse)) {
    throw new Error('Expected argument of type StatResponse');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_StatResponse(buffer_arg) {
  return math_pb.StatResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var MathService = exports.MathService = {
  sqrt: {
    path: '/Math/Sqrt',
    requestStream: false,
    responseStream: false,
    requestType: math_pb.SqrtRequest,
    responseType: math_pb.SqrtResponse,
    requestSerialize: serialize_SqrtRequest,
    requestDeserialize: deserialize_SqrtRequest,
    responseSerialize: serialize_SqrtResponse,
    responseDeserialize: deserialize_SqrtResponse,
  },
  stat: {
    path: '/Math/Stat',
    requestStream: true,
    responseStream: false,
    requestType: math_pb.StatRequest,
    responseType: math_pb.StatResponse,
    requestSerialize: serialize_StatRequest,
    requestDeserialize: deserialize_StatRequest,
    responseSerialize: serialize_StatResponse,
    responseDeserialize: deserialize_StatResponse,
  },
  factor: {
    path: '/Math/Factor',
    requestStream: false,
    responseStream: true,
    requestType: math_pb.FactorRequest,
    responseType: math_pb.FactorResponse,
    requestSerialize: serialize_FactorRequest,
    requestDeserialize: deserialize_FactorRequest,
    responseSerialize: serialize_FactorResponse,
    responseDeserialize: deserialize_FactorResponse,
  },
};

exports.MathClient = grpc.makeGenericClientConstructor(MathService);
