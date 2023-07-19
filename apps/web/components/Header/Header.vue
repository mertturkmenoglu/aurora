<template>
  <header>
    <Sliver v-if="showSliver" @close="showSliver = false"/>
    <nav class="py-4 flex justify-between items-center mx-8">
      <div class="flex items-center">
        <nuxt-link class="flex items-center" to="/">
          <CloudIcon class="h-8 w-8 text-black"/>
          <span class="text-xl text-black ml-2">Aurora</span>
        </nuxt-link>
      </div>
      <div class="flex items-end">
        <input
            v-model="searchTerm"
            :placeholder="$t('header.search')"
            class="border-b border-b-black rounded-none py-1 px-1 pr-6"
            type="text"
        />
        <nuxt-link :to="`/search?t=${searchTerm}`">
          <MagnifyingGlassIcon class="h-4 w-4 -ml-4 mb-2 text-black"/>
        </nuxt-link>

        <nuxt-link class="hover:underline ml-4 flex items-end space-x-1" to="/signin">
          <UserIcon class="h-6 w-6 text-black"/>
          <span>{{ $t('header.signIn') }}</span>
        </nuxt-link>

        <nuxt-link class="flex items-end space-x-1 ml-4 hover:underline" to="/cart">
          <ShoppingBagIcon class="h-6 w-6 text-black"/>
          <span>{{ $t('header.cart') }}</span>
        </nuxt-link>
      </div>
    </nav>
    <MegaNavigation class="my-4 mx-8"/>
  </header>
</template>

<script lang="ts" setup>
import {CloudIcon, MagnifyingGlassIcon, ShoppingBagIcon, UserIcon} from "@heroicons/vue/24/outline";

const showSliver = ref(true)
const searchTerm = ref('')
</script>