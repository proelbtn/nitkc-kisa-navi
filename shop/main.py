from concurrent import futures
import grpc
import time
from food import service_pb2, service_pb2_grpc
from food import messages_pb2 as food_messages


class FoodServiceServicer(service_pb2_grpc.FoodServiceServicer):
    def Search(self, request, context):
        return food_messages.Food(name='Hello, world')


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    service_pb2_grpc.add_FoodServiceServicer_to_server(
        FoodServiceServicer(), server)
    server.add_insecure_port('0.0.0.0:30001')
    server.start()
    try:
        while True:
            time.sleep(86400)
    except KeyboardInterrupt:
        server.stop(0)


if __name__ == '__main__':
    serve()
