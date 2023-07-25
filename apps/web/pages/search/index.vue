<template>
  <div v-if="q && data && results" class="container mx-auto">
    <h2 class="font-bold text-2xl mt-16">Search results for "{{ q }}"</h2>
    <div class="gap-8 grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 my-16">
      <div v-for="product in results" :key="product.id">
        <ProductCard :hoverable="true" :product="product"/>
      </div>
    </div>
  </div>
  <div v-else-if="pending">Loading...</div>
  <div v-else>
    <h2 class="text-center font-bold text-2xl mt-16">No search results</h2>
  </div>
</template>

<script lang="ts" setup>
import {ProductsDto} from "~/utils/dto";

const route = useRoute()
const searchTerm = route.query.q;
const q = ref('')

onMounted(() => {
  if (!searchTerm) {
    return
  }

  if (typeof searchTerm === 'string') {
    q.value = searchTerm
    return
  }
})

const {data, pending, error} = await useFetch<ProductsDto>(`${BASE_URL}/search?q=${searchTerm}`)
const results = data.value?.data || []
</script>
