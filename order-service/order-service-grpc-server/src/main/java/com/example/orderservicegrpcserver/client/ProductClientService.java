package com.example.orderservicegrpcserver.client;

import com.example.proto.order.Product;
import com.example.proto.product.ProductServiceGrpc;
import com.example.proto.product.UpdateProduct;
import com.example.proto.product.UpdateProductRequest;
import com.example.proto.product.UpdateProductResponse;
import net.devh.boot.grpc.client.inject.GrpcClient;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class ProductClientService {

    @GrpcClient("product")
    private ProductServiceGrpc.ProductServiceBlockingStub productServiceStub;

    public void updateProducts(List<Product> products) {
        List<UpdateProduct> updateProducts = products.stream()
                .map((product -> UpdateProduct.newBuilder()
                        .setId(product.getId())
                        .setQuantity(product.getQuantity())
                        .build()))
                .toList();

        UpdateProductRequest request = UpdateProductRequest.newBuilder()
                .addAllProducts(updateProducts)
                .build();

        UpdateProductResponse response = productServiceStub.updateProducts(request);

        if (response.getResponseCase() == UpdateProductResponse.ResponseCase.ERRORRESPONSE) {
            throw new IllegalStateException("Failed to update product");
        }
    }

}
