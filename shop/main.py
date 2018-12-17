from concurrent import futures
import grpc
import time
from shop import service_pb2, service_pb2_grpc
from shop import messages_pb2 as shop_messages
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
    data = shop_messages.ShopData(name=res['name'], genre_id=res['genre_id'],
                                  address=res['address'], longitude=res['longitude'], latitude=res['latitude'], open=res['open'], close=res['close'])
    return shop_messages.ShopRecord(id=res['id'], data=data)


class ShopServicer(service_pb2_grpc.ShopServicer):
    def Create(self, request, context):
        try:
            with connection.cursor() as cursor:
                sql = "INSERT INTO Shops (name, genre_id, address, longitude, latitude, open, close) VALUES (%(name)s, %(genre_id)s, %(address)s, %(longitude)s, %(latitude)s, %(open)s, %(close)s)"
                affected = cursor.execute(sql, {
                    'name': request.data.name,
                    'genre_id': request.data.genre_id,
                    'address': request.data.address,
                    'longitude': request.data.longitude,
                    'latitude': request.data.latitude,
                    'open': request.data.open,
                    'close': request.data.close
                })
                connection.commit()
            return shop_messages.ShopCreateResult(success=affected == 1)
        except Exception as e:
            print(e)

    def Delete(self, request, context):
        try:
            with connection.cursor() as cursor:
                sql = "DELETE FROM Shops WHERE id = %(id)s"
                affected = cursor.execute(sql, {
                    'id': request.id
                })
                connection.commit()
            return shop_messages.ShopDeleteResult(success=affected == 1)
        except Exception as e:
            print(e)

    def Get(self, request, context):
        try:
            with connection.cursor() as cursor:
                sql = "SELECT * FROM Shops WHERE id = %(id)s"
                cursor.execute(sql, {'id': request.id})
                result = cursor.fetchone()
            return shop_messages.ShopGetResult(record=get_record(result))
        except Exception as e:
            print(e)

    def Search(self, request, context):
        try:
            with connection.cursor() as cursor:
                params = {}
                sql = "SELECT * FROM Shops WHERE 0 = 0"
                if request.name != '':
                    params['name'] = '%' + request.name + '%'
                    sql += " AND name LIKE %(name)s"
                if request.genre_id != 0:
                    params['genre_id'] = request.genre_id
                    sql += " AND genre_id = %(genre_id)s"
                cursor.execute(sql, params)
                result = cursor.fetchall()
                print(result)
            return shop_messages.ShopSearchResult(records=[get_record(row) for row in result])
        except Exception as e:
            print(e)

    def GetGenres(self, request, context):
        try:
            with connection.cursor() as cursor:
                sql = "SELECT * FROM ShopGenres"
                cursor.execute(sql)
                result = cursor.fetchall()
            return shop_messages.ShopGetGenresResult(genres=result)
        except Exception as e:
            print(e)


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
