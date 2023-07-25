<template>
  <nuxt-link
      :to="isAuthenticated ? '/my-account' : '/signin'"
      class="transition duration-150 ease-in-out hover:bg-gray-100 rounded px-2 py-1 ml-4 flex items-end space-x-1">
    <UserIcon class="h-6 w-6 text-black"/>

    <span v-if="isAuthenticated && user" class="hidden sm:flex">My Account</span>
    <span v-else class="hidden sm:flex">{{ $t('header.signIn') }}</span>
  </nuxt-link>
</template>
<script lang="ts" setup>
import {UserIcon} from "@heroicons/vue/24/outline";

const {isAuthenticated, user} = useAuth()
</script>