from concurrent import futures
import grpc
import time
from food import service_pb2, service_pb2_grpc
from food import messages_pb2 as food_messages
from interceptors import log

from grpc_reflection.v1alpha import reflection


class FoodServicer(service_pb2_grpc.FoodServicer):
    def Search(self, request, context):
        print(request.genre)
        data = food_messages.FoodData(
            name=request.name, address='hogehoge', latitude=request.latitude, longitude=request.longitude, genre=request.genre)
        records = [food_messages.FoodRecord(id=1234, data=data)]
        return food_messages.FoodSearchResult(records=records)


def serve():
    logging = log.LoggingInterceptor()
    server = grpc.server(futures.ThreadPoolExecutor(
        max_workers=10), interceptors=(logging, ))

    service_pb2_grpc.add_FoodServicer_to_server(
        FoodServicer(), server)

    SERVICE_NAMES = (
        service_pb2.DESCRIPTOR.services_by_name['Food'].full_name,
        reflection.SERVICE_NAME
    )
    reflection.enable_server_reflection(SERVICE_NAMES, server)

    server.add_insecure_port('0.0.0.0:80')
    server.start()

    try:
        while True:
            time.sleep(1)
    except KeyboardInterrupt:
        server.stop(0)


if __name__ == '__main__':
    serve()
