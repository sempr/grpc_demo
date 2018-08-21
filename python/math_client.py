import grpc
import os
import math_pb2
import math_pb2_grpc


def call_sqrt(stub, value):
    try:
        response = stub.Sqrt(math_pb2.SqrtRequest(value=value))
        print(f"response: {response.value}")
    except grpc._channel._Rendezvous as e:
        print(e.code(), e.details())


def run(server):
    with grpc.insecure_channel(server) as channel:
        stub = math_pb2_grpc.MathStub(channel)

        call_sqrt(stub, 10)
        call_sqrt(stub, -10)


if __name__ == "__main__":
    server = os.getenv("GRPC_SERVER", "127.0.0.1:50000")
    run(server)
