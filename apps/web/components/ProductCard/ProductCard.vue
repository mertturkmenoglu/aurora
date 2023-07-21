<template>
  <nuxt-link :to="`/products/${product.id}`"
             class="hover:border hover:border-sky-600 p-4 flex flex-col rounded-lg hover:scale-[102%] transition-all ease-in duration-200 group">
    <div class="flex items-center justify-between">
      <span class="font-bold text-sm text-sky-600">{{ productMessage }}</span>
      <button>
        <HeartIcon class="w-7 h-6 text-gray-600 cursor-pointer"/>
      </button>
    </div>
    <img :src="product.images[0]?.url ?? ''" alt="" class="w-64 h-48 sm:w-64 md:h-96 object-cover mt-2 mx-auto">
    <div class="mt-2">
      <span class="font-bold text-green-600">{{ product.currentPrice }}$</span>
      <span v-if="product.currentPrice !== product.oldPrice" class="font-light line-through ml-2">
            {{ product.oldPrice }}$
        </span>
    </div>
    <span class="text-gray-600 text-sm mt-2">{{ product.category.name }}</span>
    <span class="font-bold mt-2 line-clamp-2 leading-4 h-8">{{ product.name }}</span>
    <span class="font-light text-sm mt-2">By {{ product.brand.name }}</span>

    <button
        class="hidden group-hover:flex px-4 py-2 text-white bg-sky-600 rounded-full text-center justify-center mt-4">
      Add to cart
    </button>
  </nuxt-link>

</template>

<script lang="ts" setup>
import {Product} from "~/utils/dto";
import {HeartIcon} from "@heroicons/vue/24/outline"
import {useProductMessage} from "~/composables/useProductMessage";

const {product} = defineProps<{
  product: Product
}>();

const productMessage = useProductMessage(product);
</script>