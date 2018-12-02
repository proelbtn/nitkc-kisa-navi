from concurrent import futures
import grpc
import time
from food import service_pb2, service_pb2_grpc
from food import messages_pb2 as food_messages
from interceptors import log


class FoodServicer(service_pb2_grpc.FoodServicer):
    def Search(self, request, context):
        return food_messages.FoodResponse(name=request.name)


def serve():
    print("food service")
    logging = log.LoggingInterceptor()
    server = grpc.server(futures.ThreadPoolExecutor(
        max_workers=10), interceptors=(logging, ))

    service_pb2_grpc.add_FoodServicer_to_server(
        FoodServicer(), server)

    server.add_insecure_port('localhost:30001')
    server.start()

    try:
        while True:
            time.sleep(1)
    except KeyboardInterrupt:
        server.stop(0)


if __name__ == '__main__':
    serve()
