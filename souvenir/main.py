from concurrent import futures
import grpc
import time
from souvenir import service_pb2, service_pb2_grpc
from souvenir import messages_pb2 as souvenir_messages
from interceptors import log

from grpc_reflection.v1alpha import reflection


class SouvenirServicer(service_pb2_grpc.SouvenirServicer):
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

    service_pb2_grpc.add_SouvenirServicer_to_server(
        SouvenirServicer(), server)

    SERVICE_NAMES = (
        service_pb2.DESCRIPTOR.services_by_name['Souvenir'].full_name,
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
