<template>
  <div v-if="isAuthenticated && user" class="container mx-auto my-16">
    <nuxt-link class="flex items-center text-sky-600 text-sm hover:underline" to="/my-account">
      <ChevronLeftIcon class="w-6 h-6 mr-2"/>
      <span class="font-semibold">Back to My Account</span>
    </nuxt-link>

    <div class="mt-8 mb-64">
      <h2 class="text-2xl text-sky-600 font-bold">Your Addresses</h2>

      <div v-for="address in user.addresses" v-if="user.addresses.length > 0">
        <div class="mt-4 space-y-2">
          <pre class="flex space-x-2">
            {{ JSON.stringify(address, null, 2) }}
          </pre>
        </div>
      </div>
      <div v-else class="mt-8">
        You don't have any addresses
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import {ChevronLeftIcon} from "@heroicons/vue/24/outline";

const {isAuthenticated, user} = useAuth()
</script>