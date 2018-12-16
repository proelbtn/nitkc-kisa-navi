from concurrent import futures
import grpc
import time
from souvenir import service_pb2, service_pb2_grpc
from souvenir import messages_pb2 as souvenir_messages
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


def get_record(res):
    data = souvenir_messages.SouvenirData(
        name=res['name'], genre_id=res['genre_id'], price=res['price'])
    return souvenir_messages.SouvenirRecord(id=res['id'], data=data)


class SouvenirServicer(service_pb2_grpc.SouvenirServicer):
    def Create(self, request, context):
        try:
            with connection.cursor() as cursor:
                sql = "INSERT INTO Souvenirs (name, genre_id, price) VALUES (%(name)s, %(genre_id)s, %(price)s)"
                affected = cursor.execute(sql, {
                    'name': request.data.name,
                    'genre_id': request.data.genre_id,
                    'price': request.data.price,
                })
                connection.commit()
            return souvenir_messages.SouvenirCreateResult(success=affected == 1)
        except Exception as e:
            print(e)

    def Delete(self, request, context):
        try:
            with connection.cursor() as cursor:
                sql = "DELETE FROM Souvenirs WHERE id = %(id)s"
                affected = cursor.execute(sql, {
                    'id': request.id
                })
                connection.commit()
            return souvenir_messages.SouvenirDeleteResult(success=affected == 1)
        except Exception as e:
            print(e)

    def Get(self, request, context):
        try:
            with connection.cursor() as cursor:
                sql = "SELECT * FROM Souvenirs WHERE id = %(id)s"
                result = cursor.execute(sql, {
                    'id': request.id
                })
                connection.commit()
            return souvenir_messages.SouvenirDeleteResult(record=get_record(result))
        except Exception as e:
            print(e)

    def Search(self, request, context):
        try:
            with connection.cursor() as cursor:
                params = {}
                sql = "SELECT * FROM Souvenirs WHERE 0 = 0"
                if request.name != '':
                    params['name'] = '%' + request.name + '%'
                    sql += " AND name LIKE %(name)s"
                if request.genre_id != 0:
                    params['genre_id'] = request.genre_id
                    sql += " AND genre_id = %(genre_id)s"
                cursor.execute(sql, params)
                result = cursor.fetchall()
            return souvenir_messages.SouvenirSearchResult(records=[get_record(row) for row in result])
        except Exception as e:
            print(e)

    def GetGenres(self, request, context):
        with connection.cursor() as cursor:
            sql = "SELECT * FROM SouvenirGenres"
            cursor.execute(sql)
            result = cursor.fetchall()
        return souvenir_messages.SouvenirGetGenresResult(genres=result)


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
