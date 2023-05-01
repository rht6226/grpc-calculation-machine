# gRPC client-server machine

The project is a client-server application that communicates via gRPC. The client will stream a list of instructions to the server, and the server will stream back the results of executing those instructions.

The client application will be responsible for gathering a list of instructions from some source (such as a file or user input) and streaming them to the server. The server application will receive the instructions, execute them, and stream back the results to the client.

To facilitate this communication, the project will define a set of protobuf messages for the client and server to use. These messages will include a request message for the client to send instructions to the server and a response message for the server to send results back to the client. The project will also define gRPC service definitions that specify how the client and server can interact with each other.

# instructions to run

    $ make gen
    $ make mock
    $ make test
    $ make run-server
    $ make run-client