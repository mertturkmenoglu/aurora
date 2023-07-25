<template>
  <div class="flex">
    <nuxt-link to="/">
      <HomeIcon class="w-6 h-6 text-sky-600"/>
    </nuxt-link>
    <div v-for="item in items" :key="item.name" class="flex items-center">
      <ChevronRightIcon class="w-4 h-4 text-gray-600 ml-1"/>
      <nuxt-link
          :to="item.href"
          class="text-gray-600 hover:text-sky-600 transition ease-in-out duration-200 ml-1"
      >
        {{ item.name }}
      </nuxt-link>
    </div>
  </div>
</template>

<script lang="ts" setup>
import {HomeIcon, ChevronRightIcon} from "@heroicons/vue/24/solid";

defineProps<{
  items: Array<{
    name: string;
    href: string;
  }>
}>()
</script>