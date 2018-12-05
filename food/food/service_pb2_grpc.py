# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

from food import messages_pb2 as food_dot_messages__pb2


class FoodStub(object):
  # missing associated documentation comment in .proto file
  pass

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.Create = channel.unary_unary(
        '/food.Food/Create',
        request_serializer=food_dot_messages__pb2.FoodCreateQuery.SerializeToString,
        response_deserializer=food_dot_messages__pb2.FoodCreateResult.FromString,
        )
    self.Delete = channel.unary_unary(
        '/food.Food/Delete',
        request_serializer=food_dot_messages__pb2.FoodDeleteQuery.SerializeToString,
        response_deserializer=food_dot_messages__pb2.FoodDeleteQuery.FromString,
        )
    self.Get = channel.unary_unary(
        '/food.Food/Get',
        request_serializer=food_dot_messages__pb2.FoodGetQuery.SerializeToString,
        response_deserializer=food_dot_messages__pb2.FoodGetResult.FromString,
        )
    self.Search = channel.unary_unary(
        '/food.Food/Search',
        request_serializer=food_dot_messages__pb2.FoodSearchQuery.SerializeToString,
        response_deserializer=food_dot_messages__pb2.FoodSearchResult.FromString,
        )


class FoodServicer(object):
  # missing associated documentation comment in .proto file
  pass

  def Create(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def Delete(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def Get(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def Search(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_FoodServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'Create': grpc.unary_unary_rpc_method_handler(
          servicer.Create,
          request_deserializer=food_dot_messages__pb2.FoodCreateQuery.FromString,
          response_serializer=food_dot_messages__pb2.FoodCreateResult.SerializeToString,
      ),
      'Delete': grpc.unary_unary_rpc_method_handler(
          servicer.Delete,
          request_deserializer=food_dot_messages__pb2.FoodDeleteQuery.FromString,
          response_serializer=food_dot_messages__pb2.FoodDeleteQuery.SerializeToString,
      ),
      'Get': grpc.unary_unary_rpc_method_handler(
          servicer.Get,
          request_deserializer=food_dot_messages__pb2.FoodGetQuery.FromString,
          response_serializer=food_dot_messages__pb2.FoodGetResult.SerializeToString,
      ),
      'Search': grpc.unary_unary_rpc_method_handler(
          servicer.Search,
          request_deserializer=food_dot_messages__pb2.FoodSearchQuery.FromString,
          response_serializer=food_dot_messages__pb2.FoodSearchResult.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'food.Food', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))
