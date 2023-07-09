<template>
  <header>
    <div v-if="showSliver" class="bg-black text-white py-2 flex items-center px-8 justify-center relative">
      <span class="">Welcome to Aurora</span>
      <button @click="showSliver = false" class="absolute right-8">
        <XMarkIcon class="h-4 w-4 text-white"/>
      </button>
    </div>
    <nav class="py-4 flex justify-between items-center">
      <div class="flex items-center">
        <nuxt-link to="/" class="flex items-center">
          <CloudIcon class="h-8 w-8 text-black"/>
          <span class="text-xl text-black ml-2">Aurora</span>
        </nuxt-link>
      </div>
      <div class="flex items-end">
        <input v-model="searchTerm" type="text" class="border-b border-b-black rounded-none py-1 px-1" placeholder="Search"/>
        <nuxt-link :to="`/search?t=${searchTerm}`">
          <MagnifyingGlassIcon class="h-4 w-4 -ml-4 mb-2 text-black"/>
        </nuxt-link>

        <nuxt-link to="/signin" class="hover:underline ml-4 flex items-end space-x-1">
          <UserIcon class="h-6 w-6 text-black" />
          <span>Sign in</span>
        </nuxt-link>

        <nuxt-link to="/cart" class="flex items-end space-x-1 ml-4 hover:underline">
          <ShoppingBagIcon class="h-6 w-6 text-black"/>
          <span>Cart</span>
        </nuxt-link>
      </div>

    </nav>
  </header>
</template>
<script setup lang="ts">
import {XMarkIcon, CloudIcon, ShoppingBagIcon, MagnifyingGlassIcon, UserIcon} from '@heroicons/vue/24/outline'

const showSliver = ref(true)
const searchTerm = ref('')
</script>