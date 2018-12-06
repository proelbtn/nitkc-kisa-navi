from concurrent import futures
import grpc
import time
from shop import service_pb2, service_pb2_grpc
from shop import messages_pb2 as shop_messages
from interceptors import log

from grpc_reflection.v1alpha import reflection


class ShopServicer(service_pb2_grpc.ShopServicer):
    def Create(self, request, context):
        pass

    def Delete(self, request, context):
        pass

    def Get(self, request, context):
        pass

    def Search(self, request, context):
        pass


def serve():
    logging = log.LoggingInterceptor()
    server = grpc.server(futures.ThreadPoolExecutor(
        max_workers=10), interceptors=(logging, ))

    service_pb2_grpc.add_ShopServicer_to_server(
        ShopServicer(), server)

    SERVICE_NAMES = (
        service_pb2.DESCRIPTOR.services_by_name['Shop'].full_name,
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
