<template>
  <button @click.prevent="onFavoriteClick">
    <HeartIconFilled v-if="isFavorite" class="w-6 h-6 text-sky-600 cursor-pointer"/>
    <HeartIconEmpty v-else class="w-6 h-6 text-gray-600 cursor-pointer"/>
  </button>
</template>

<script lang="ts" setup>
import {HeartIcon as HeartIconEmpty} from "@heroicons/vue/24/outline";
import {HeartIcon as HeartIconFilled} from "@heroicons/vue/24/solid";

const {data: fav} = await useAsyncData('fav', () => FavoriteManager.getAsyncInstance())
const isFavorite = ref(false);
const isValid = ref(true);

const {productId} = defineProps<{
  productId: string
}>();

async function onFavoriteClick() {
  if (!fav || !fav.value) {
    return;
  }

  if (isFavorite.value) {
    await fav.value.removeFromFavorite(productId)
  } else {
    await fav.value.addToFavorite(productId)
  }

  isValid.value = false;
}

watch([isValid, fav], () => {
  if (fav.value) {
    isFavorite.value = fav.value.isFavorite(productId)
  }
}, {immediate: true})
</script>