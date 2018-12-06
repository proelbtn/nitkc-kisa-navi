import grpc


class LoggingInterceptor(grpc.ServerInterceptor):
    def intercept_service(self, continuation, handler_call_details):
        print(handler_call_details)
        return continuation(handler_call_details)
