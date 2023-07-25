<template>
  <Head>
    <title>Aurora / Categories</title>
  </Head>
  <div v-if="products" class="container gap-8 grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 mx-auto xl:px-32 my-16">
    <div v-for="product in products" :key="product.id">
      <ProductCard :product="product"/>
    </div>
  </div>
</template>

<script lang="ts" setup>
import {ProductsDto, Product} from "~/utils/dto";

const route = useRoute();

const {
  data,
} = await useFetch<ProductsDto>(`${BASE_URL}/products?categoryId=${route.params.id}`);

const products: Product[] = data.value?.data || [];
</script>