from concurrent import futures
import time
import os
import logging
import grpc
import math
import math_pb2
import math_pb2_grpc


class Mather(math_pb2_grpc.MathServicer):
    def Sqrt(self, request, context):
        return math_pb2.SqrtResponse(value=math.sqrt(request.value))


def serve(bind="0.0.0.0:50000"):
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    math_pb2_grpc.add_MathServicer_to_server(Mather(), server)
    logging.warning(f"start grpc server on {bind}")
    server.add_insecure_port(bind)
    server.start()
    try:
        while True:
            time.sleep(3600)
    except KeyboardInterrupt:
        server.stop(0)


if __name__ == "__main__":
    bind = os.getenv("GRPC_BIND", "0.0.0.0:50000")
    serve(bind)