package com.example.orderservicegrpcserver.service;

import com.example.orderservicegrpcserver.entity.Cart;
import com.example.orderservicegrpcserver.entity.CartProduct;
import com.example.orderservicegrpcserver.entity.Order;
import com.example.orderservicegrpcserver.repository.OrderRepository;
import com.example.proto.proto.Product;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.stream.Collectors;

@Service
public class OrderService {

    @Autowired
    private OrderRepository orderRepository;

    public Order createOrder(Long userId, List<Product> products) {
        List<CartProduct> cartProducts = products.stream()
                .map((product) -> CartProduct.builder()
                        .productId(product.getId())
                        .name(product.getName())
                        .description(product.getDescription())
                        .price(product.getPrice())
                        .quantity(product.getQuantity())
                        .build())
                .collect(Collectors.toList());

        Double totalPrice = products.stream()
                .mapToDouble((product -> product.getPrice() * product.getQuantity()))
                .sum();

        Cart cart = Cart.builder()
                .cartProducts(cartProducts)
                .totalPrice(totalPrice)
                .build();

        Order order = Order.builder()
                .userId(userId)
                .cart(cart)
                .build();

        return orderRepository.save(order);
    }

    public List<Order> getAllOrdersByUserId(Long userId) {
        return orderRepository.findByUserId(userId);
    }

}
