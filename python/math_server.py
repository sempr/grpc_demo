from concurrent import futures
import time
import os
import logging
import grpc
import math
import math_pb2
import math_pb2_grpc


def factor(val: int):
    i = 2
    while i*i < val:
        while val % i == 0:
            val //= i
            yield i
        i += 1
    if val > 1:
        yield val


class Mather(math_pb2_grpc.MathServicer):
    def Sqrt(self, request, context):
        value = request.value
        if value < 0:
            context.set_code(grpc.StatusCode.INVALID_ARGUMENT)
            context.set_details("负数不能开平方根".encode("utf8"))
            return math_pb2.SqrtResponse(value=0)
        return math_pb2.SqrtResponse(value=math.sqrt(value))

    def Stat(self, request_iterator, context):
        summ, count = 0, 0
        for stat in request_iterator:
            summ += stat.value
            count += 1
        return math_pb2.StatResponse(sum=summ, count=count)

    def Factor(self, request, context):
        value = request.value
        print(value)
        for val in factor(value):
            print(val)
            yield math_pb2.FactorResponse(value=val)


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
