<template>
  <header>
    <Sliver v-if="showSliver" @close="showSliver = false"/>
    <nav class="py-4 flex justify-between items-center mx-8">
      <div class="flex items-center">
        <nuxt-link class="flex items-center" to="/">
          <CloudIcon class="h-8 w-8 text-sky-600"/>
          <span class="text-xl text-sky-600 ml-2">Aurora</span>
        </nuxt-link>
      </div>
      <div class="flex items-end">
        <input
            v-model="searchTerm"
            :placeholder="$t('header.search')"
            class="border-b border-b-black rounded-none py-1 px-1 pr-6 hidden sm:flex"
            type="text"
        />
        <nuxt-link :to="`/search?q=${searchTerm}`" class="z-10">
          <MagnifyingGlassIcon class="h-6 w-6 sm:h-4 sm:w-4 -ml-4 sm:mb-2 text-black"/>
        </nuxt-link>

        <ClientOnly fallback="" fallback-tag="span">
          <HeaderUserSection/>
        </ClientOnly>

        <nuxt-link
            class="flex items-end space-x-1 ml-4 transition duration-150 ease-in-out hover:bg-gray-100 rounded px-2 py-1"
            to="/cart">
          <ShoppingBagIcon class="h-6 w-6 text-black"/>
          <span class="hidden sm:flex">{{ $t('header.cart') }}</span>
        </nuxt-link>
      </div>
    </nav>
    <ClientOnly fallback="" fallback-tag="span">
      <MegaNavigation class="my-4 mx-8"/>
    </ClientOnly>
  </header>
</template>

<script lang="ts" setup>
import {MagnifyingGlassIcon, ShoppingBagIcon} from "@heroicons/vue/24/outline";
import {CloudIcon} from "@heroicons/vue/24/solid";

const showSliver = ref(true)
const searchTerm = ref('')
</script>