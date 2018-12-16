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
    except Exception as e:
        print(e)
        time.sleep(5)


def get_record(res):
    data = food_messages.FoodData(
        name=res['name'], price=res['price'], genre_id=res['genre_id'])
    return food_messages.FoodRecord(id=res['id'], data=data)


class FoodServicer(service_pb2_grpc.FoodServicer):
    def Create(self, request, context):
        try:
            with connection.cursor() as cursor:
                sql = "INSERT INTO Foods (name, genre_id, price) VALUES (%(name)s, %(genre_id)s, %(price)s)"
                affected = cursor.execute(sql, {
                    'name': request.data.name,
                    'genre_id': request.data.genre_id,
                    'price': request.data.price
                })
                connection.commit()
            return food_messages.FoodCreateResult(success=affected == 1)
        except Exception as e:
            print(e)

    def Delete(self, request, context):
        try:
            with connection.cursor() as cursor:
                sql = "DELETE FROM Foods WHERE id = %(id)s"
                affected = cursor.execute(sql, {
                    'id': request.id
                })
                connection.commit()
            return food_messages.FoodDeleteResult(success=affected == 1)
        except Exception as e:
            print(e)

    def Get(self, request, context):
        try:
            with connection.cursor() as cursor:
                sql = "SELECT * FROM Foods WHERE id = %(id)s"
                cursor.execute(sql, {'id': request.id})
                result = cursor.fetchone()
            return food_messages.FoodGetResult(record=get_record(result))
        except Exception as e:
            print(e)

    def Search(self, request, context):
        try:
            with connection.cursor() as cursor:
                params = {}
                sql = "SELECT * FROM Foods WHERE 0 = 0"
                if request.name != '':
                    params['name'] = '%' + request.name + '%'
                    sql += " AND name LIKE %(name)s"
                if request.genre_id != 0:
                    params['genre_id'] = request.genre_id
                    sql += " AND genre_id = %(genre_id)s"
                cursor.execute(sql, params)
                result = cursor.fetchall()
            return food_messages.FoodSearchResult(records=[get_record(row) for row in result])
        except Exception as e:
            print(e)

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
