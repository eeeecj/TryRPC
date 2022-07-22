import asyncio
import sys
from internal.gRPC.options.option import Config
from internal.gRPC.server.server import NewGrpcServer

class App():
    def __init__(self,config:Config) -> None:
        self.config=config
        self.Server=NewGrpcServer(self.config.config['grpc'])
        

    async def Run(self):
        await self.Server.start()
        self.Shutdown=self.stratShutdown()
        await self.Server.wait_for_termination()
        

    async def stratShutdown(self):
        await self.Server.stop(5)