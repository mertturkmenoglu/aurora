<template>
  <Head>
    <title>Aurora / Home</title>
  </Head>

  <div class='container space-y-16 mx-auto'>
    <ClientOnly fallback-tag='span'>
      <Banner />
    </ClientOnly>

    <ProductCarousel :items='featuredProducts as Product[]' title='Featured products' />

    <ProductCarousel :items='popularProducts as Product[]' title='Popular products' />

    <ProductCarousel :items='newProducts as Product[]' title='New products' />

    <ProductCarousel :items='saleProducts as Product[]' title='Sale products' />

  </div>

  <div v-if='products' class='container gap-8 grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 mx-auto xl:px-32 my-16'>
    <div v-for='product in products' :key='product.id'>
      <ProductCard :hoverable='true' :product='product' />
    </div>
  </div>
</template>

<script lang='ts' setup>
import { ProductsDto, Product, HomeAggregationDto } from '~/utils/dto';
import { BASE_URL } from '~/utils/api';

const {
  data,
} = await useFetch<ProductsDto>(`${BASE_URL}/products/all?page=1&pageSize=90`);

const { data: aggregatedData } = await useFetch<HomeAggregationDto>(`${BASE_URL}/aggregations/home`);

const featuredProducts = aggregatedData?.value?.data?.featured ?? [];
const popularProducts = aggregatedData?.value?.data?.popular ?? [];
const newProducts = aggregatedData?.value?.data?.new ?? [];
const saleProducts = aggregatedData?.value?.data?.sale ?? [];

const products: Product[] = data.value?.data || [];
</script>
