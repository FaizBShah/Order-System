package com.example.orderservicegrpcserver.entity;

import jakarta.persistence.*;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.util.List;

@Entity
@Table(name = "carts")
@Data
@NoArgsConstructor
@AllArgsConstructor
@Builder
public class Cart {

    @Id
    @SequenceGenerator(
            name = "carts_sequence",
            sequenceName = "carts_sequence",
            allocationSize = 1
    )
    @GeneratedValue(
            strategy = GenerationType.SEQUENCE,
            generator = "carts_sequence"
    )
    private Long id;

    @ElementCollection(fetch = FetchType.EAGER)
    @CollectionTable(name = "cart_products")
    private List<CartProduct> cartProducts;

    @Column(nullable = false)
    private Double totalPrice;

}
