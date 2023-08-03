<template>
  <Head>
    <title>Aurora / Home</title>
  </Head>

  <div class='container space-y-16 mx-auto'>
    <ClientOnly fallback-tag='span'>
      <Banner />
    </ClientOnly>

    <ProductCarousel
      :items='featuredProducts as Product[]'
      title='Featured products'
    />

    <ProductCarousel
      :items='popularProducts as Product[]'
      title='Popular products'
    />

    <ProductCarousel
      :items='newProducts as Product[]'
      title='New products'
    />

  </div>

  <div v-if='products' class='container gap-8 grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 mx-auto xl:px-32 my-16'>
    <div v-for='product in products' :key='product.id'>
      <ProductCard :hoverable='true' :product='product' />
    </div>
  </div>
</template>

<script lang='ts' setup>
import { ProductsDto, Product } from '~/utils/dto';
import { BASE_URL } from '~/utils/api';

const {
  data,
} = await useFetch<ProductsDto>(`${BASE_URL}/products/all?page=1&pageSize=90`);

const {
  data: featuredProductsData,
} = await useFetch<ProductsDto>(`${BASE_URL}/products/featured`);

const {
  data: popularProductsData,
} = await useFetch<ProductsDto>(`${BASE_URL}/products/popular`);

const {
  data: newProductsData,
} = await useFetch<ProductsDto>(`${BASE_URL}/products/new`);

const featuredProducts = featuredProductsData.value?.data ?? [];

const popularProducts = popularProductsData.value?.data ?? [];

const newProducts = newProductsData.value?.data ?? [];

const products: Product[] = data.value?.data || [];
</script>