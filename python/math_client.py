import grpc
import os
import math_pb2
import math_pb2_grpc


def run(server):
    with grpc.insecure_channel(server) as channel:
        stub = math_pb2_grpc.MathStub(channel)
        response = stub.Sqrt(math_pb2.SqrtRequest(value=3.0))
        print(f"response: {response.value}")


if __name__ == "__main__":
    server = os.getenv("GRPC_SERVER", "127.0.0.1:50000")
    run(server)
