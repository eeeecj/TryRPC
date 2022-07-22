import api.proto.python.hello_pb2_grpc as hello_pb2_grpc
import api.proto.python.hello_pb2 as hello_pb2
import grpc

# Hello定义grpc服务
class Hello(hello_pb2_grpc.HelloServicer):
    async def Hello(self, request: hello_pb2.HelloResponse, context: grpc.aio.ServicerContext) -> hello_pb2.HelloResponse:
        return hello_pb2.HelloResponse(output="hello "+request.input+" Test")