package com.example.orderservicegrpcserver.controller;

import com.example.orderservicegrpcserver.entity.Order;
import com.example.orderservicegrpcserver.service.OrderService;
import com.example.proto.order.*;
import io.grpc.stub.StreamObserver;
import net.devh.boot.grpc.server.service.GrpcService;
import org.springframework.beans.factory.annotation.Autowired;

import java.util.List;
import java.util.stream.Collectors;

@GrpcService
public class OrderGrpcController extends OrderServiceGrpc.OrderServiceImplBase {

    @Autowired
    private OrderService orderService;

    @Override
    public void createOrder(CreateOrderRequest request, StreamObserver<CreateOrderResponse> responseObserver) {
        Order order = orderService.createOrder(request.getUserId(), request.getCart().getProductsList());

        CreateOrderResponse response = orderService.convertOrderToCreateOrderResponse(order);

        responseObserver.onNext(response);
        responseObserver.onCompleted();
    }

    @Override
    public void getAllOrdersByUserId(GetAllOrdersByUserIdRequest request, StreamObserver<GetAllOrdersByUserIdResponse> responseObserver) {
        List<Order> orders = orderService.getAllOrdersByUserId(request.getUserId());

        List<CreateOrderResponse> userOrders = orders.stream()
                .map((order) -> orderService.convertOrderToCreateOrderResponse(order))
                .toList();

        GetAllOrdersByUserIdResponse response = GetAllOrdersByUserIdResponse.newBuilder()
                .addAllOrders(userOrders)
                .build();

        responseObserver.onNext(response);
        responseObserver.onCompleted();
    }
}
