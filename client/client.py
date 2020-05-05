import logging

import grpc

import helloworld_pb2
import helloworld_pb2_grpc


def run():
    # NOTE(gRPC Python Team): .close() is possible on a channel and should be
    # used in circumstances in which the with statement does not fit the needs
    # of the code

    with grpc.insecure_channel('localhost:50051') as channel:
        stub = helloworld_pb2_grpc.GreeterStub(channel)

        message = helloworld_pb2.HelloRequest(name='m0ai')
        response = stub.SayHello(message)
        print("Greeter client received: " + response.message)

        message = helloworld_pb2.HelloRequest(name='m0ai')
        response = stub.SayHelloAgain(message)
        print("Greeter client received: " + response.message)


if __name__ == '__main__':
    logging.basicConfig()
    run()
