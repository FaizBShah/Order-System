package com.example.orderservicegrpcserver.service;

import com.example.proto.proto.ProductServiceGrpc;
import net.devh.boot.grpc.server.service.GrpcService;

@GrpcService
public class OrderGrpcService extends ProductServiceGrpc.ProductServiceImplBase {
}
