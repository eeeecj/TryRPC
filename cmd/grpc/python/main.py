import sys
sys.path.append(".")
import asyncio 
from internal.gRPC.options.option import Config
from internal.gRPC.app import App

def main():
    config=Config("./config/grpc/grpc-server.ini")
    app=App(config)
    loop=asyncio.get_event_loop()
    # 优雅关闭
    try:
        loop.run_until_complete(app.Run())
    finally:
        loop.run_until_complete(app.Shutdown)
        loop.close()

if __name__ == '__main__':
    main()
