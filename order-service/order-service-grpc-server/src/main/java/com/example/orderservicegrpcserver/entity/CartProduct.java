package com.example.orderservicegrpcserver.entity;

import jakarta.persistence.AttributeOverride;
import jakarta.persistence.AttributeOverrides;
import jakarta.persistence.Column;
import jakarta.persistence.Embeddable;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

@Embeddable
@Data
@NoArgsConstructor
@AllArgsConstructor
@Builder
@AttributeOverrides({
        @AttributeOverride(
                name = "productId",
                column = @Column(nullable = false)
        ),
        @AttributeOverride(
                name = "name",
                column = @Column(nullable = false)
        ),
        @AttributeOverride(
                name = "description",
                column = @Column(nullable = false)
        ),
        @AttributeOverride(
                name = "price",
                column = @Column(nullable = false)
        ),
        @AttributeOverride(
                name = "quantity",
                column = @Column(nullable = false)
        )
})
public class CartProduct {
    private Long productId;
    private String name;
    private String description;
    private Double price;
    private int quantity;
}
