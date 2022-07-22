from concurrent import futures
import api.proto.python.hello_pb2_grpc as hello_pb2_grpc
import api.proto.python.hello_pb2 as hello_pb2
import grpc
from grpc_reflection.v1alpha import reflection
from internal.gRPC.server.hello import Hello

# NewGrpcServer 创建异步grpc服务
def NewGrpcServer(config):
    # 创建RPC服务
    server = grpc.aio.server(futures.ThreadPoolExecutor(max_workers=5))
    hello_pb2_grpc.add_HelloServicer_to_server(Hello(), server)
    SERVER_NAME = (
        hello_pb2.DESCRIPTOR.services_by_name['Hello'].full_name,
        reflection.SERVICE_NAME,
    )
    reflection.enable_server_reflection(service_names=SERVER_NAME,server=server)
    server.add_insecure_port(config['host']+':'+config['port'])
    return server
