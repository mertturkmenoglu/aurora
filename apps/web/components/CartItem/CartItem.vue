<template>
  <div class="flex justify-between max-w-6xl items-center border border-gray-600 rounded-lg p-2">
    <nuxt-link :to="`/products/${product.id}`" class="flex items-center space-x-4">
      <img :src="image" alt="" class="w-32 h-48 object-cover mt-2 mx-auto" loading="lazy">
      <div class="mt-2 flex flex-col">
        <span class="font-bold mt-2 line-clamp-2 leading-4 h-8">{{ product.name }}</span>
        <span class="font-light text-sm mt-2">By {{ product.brand.name }}</span>
        <span class="font-bold text-green-600">{{ product.currentPrice }}$</span>
        <span v-if="product.currentPrice !== product.oldPrice" class="font-light line-through ml-2">
            {{ product.oldPrice }}$
        </span>
      </div>
    </nuxt-link>

    <div class="flex items-center space-x-16">

      <div class="flex items-center space-x-4">
        <button class="hover:bg-gray-200 rounded-full p-2 transition duration-200 ease-in-out">
          <MinusIcon class="text-gray-600 h-6 w-6"/>
        </button>

        <span class="border-b border-gray-600 px-4 py-2">
        {{ item.quantity }}
      </span>

        <button class="hover:bg-gray-200 rounded-full p-2 transition duration-200 ease-in-out">
          <PlusIcon class="text-gray-600 h-6 w-6"/>
        </button>
      </div>

      <button class="hover:bg-red-100 p-2 rounded-full transition duration-200 ease-in-out">
        <TrashIcon class="h-6 w-6 text-red-600"/>
      </button>

    </div>
  </div>
</template>

<script lang="ts" setup>
import {CartItem as TCartItem} from "~/utils/dto";
import {TrashIcon, MinusIcon, PlusIcon} from "@heroicons/vue/24/outline";

const {item} = defineProps<{
  item: TCartItem
}>()

const product = item.product
const image = product.images[0].url ?? ''
</script>