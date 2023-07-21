<template>
  <nuxt-link
      :class="clsx(
          {
          'hover:scale-[102%] transition-all ease-in duration-200 hover:border hover:border-sky-600': hoverable,
          },
          ' p-4 flex flex-col rounded-lg group'
      )"
      :to="`/products/${product.id}`"
  >
    <div class="flex items-center justify-between">
      <span class="font-bold text-sm text-sky-600">{{ productMessage }}</span>
      <button>
        <HeartIcon class="w-7 h-6 text-gray-600 cursor-pointer"/>
      </button>
    </div>
    <img :src="image" alt="" class="w-64 h-48 sm:w-64 md:h-96 object-cover mt-2 mx-auto" loading="lazy">
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
        :class="clsx(
            {
              'hidden group-hover:flex': hoverable,
              'flex': !hoverable,
            },
            'px-4 py-2 text-white bg-sky-600 rounded-full text-center justify-center mt-4'
        )">
      Add to cart
    </button>
  </nuxt-link>

</template>

<script lang="ts" setup>
import {Product} from "~/utils/dto";
import {HeartIcon} from "@heroicons/vue/24/outline"
import {useProductMessage} from "~/composables/useProductMessage";
import clsx from "clsx";

const {product, hoverable = true} = defineProps<{
  product: Product
  hoverable?: boolean
}>();

const productMessage = useProductMessage(product);

const image = computed(() => {
  if (!product.images.length) {
    return ''
  }

  return product.images[0]?.url ?? ''
})
</script>