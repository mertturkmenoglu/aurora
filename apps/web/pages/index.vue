<template>
  <div v-if="products" class="container gap-8 grid grid-cols-4 mx-auto px-32 my-16">
    <div v-for="product in [...products, ...products, ...products]" :key="product.id">
      <ProductCard :product="product"/>
    </div>
  </div>
</template>

<script lang="ts" setup>
import {ProductsDto, Product} from "~/utils/dto";

const {
  data,
} = await useFetch<ProductsDto>('http://localhost:5000/api/v1/products?categoryId=e2c46077-b90f-449d-a12b-d1280d3ee42b');

const products: Product[] = data.value?.data || [];
</script>