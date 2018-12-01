from concurrent import futures
import grpc
import time
import food_pb2
import food_pb2_grpc

class FoodServicer(food_pb2_grpc.FoodServicer):
    def Search(self, request, context):
        return food_pb2.FoodResponse(name='Hello, world')

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    food_pb2_grpc.add_FoodServicer_to_server(FoodServicer(), server)
    server.add_insecure_port('localhost:50051')
    server.start()
    try:
        while True:
            time.sleep(86400)
    except KeyboardInterrupt:
        server.stop(0)

if __name__ == '__main__':
    serve()