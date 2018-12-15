from concurrent import futures
import grpc
import time
from food import service_pb2, service_pb2_grpc
from food import messages_pb2 as food_messages
from interceptors import log

from grpc_reflection.v1alpha import reflection
import pymysql


config = {
    'host': 'db',
    'user': 'root',
    'password': 'password',
    'db': 'eve-navi',
    'charset': 'utf8mb4',
    'cursorclass': pymysql.cursors.DictCursor
}

while True:
    try:
        connection = pymysql.connect(**config)
        break
    except:
        time.sleep(1)


class FoodServicer(service_pb2_grpc.FoodServicer):
    def Create(self, request, context):
        return food_messages.FoodCreateResult(success=True)

    def Delete(self, request, context):
        return food_messages.FoodDeleteResult(success=True)

    def Get(self, request, context):
        pass

    def Search(self, request, context):
        pass

    def GetGenres(self, request, context):
        with connection.cursor() as cursor:
            sql = "SELECT * FROM FoodGenres"
            cursor.execute(sql)
            result = cursor.fetchall()
        return food_messages.FoodGetGenresResult(genres=result)


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
