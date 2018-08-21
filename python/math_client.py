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


def call_stat(stub, num_from, num_to, step=1):
    request_iter = (math_pb2.StatRequest(value=i) for i in range(num_from, num_to, step))
    response = stub.Stat(request_iter)
    print(f"results: {response.sum} {response.count}")


def run(server):
    with grpc.insecure_channel(server) as channel:
        stub = math_pb2_grpc.MathStub(channel)

        call_sqrt(stub, 10)
        call_sqrt(stub, -10)

        call_stat(stub, 1, 3, 1)
        call_stat(stub, 1, 10, 2)


if __name__ == "__main__":
    server = os.getenv("GRPC_SERVER", "127.0.0.1:50000")
    run(server)
