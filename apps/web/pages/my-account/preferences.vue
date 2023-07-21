<template>
  <div v-if="isAuthenticated && user" class="container mx-auto my-16">
    <nuxt-link class="flex items-center text-sky-600 text-sm hover:underline" to="/my-account">
      <ChevronLeftIcon class="w-6 h-6 mr-2"/>
      <span class="font-semibold">Back to My Account</span>
    </nuxt-link>

    <div class="mt-8 mb-64">
      <h2 class="text-2xl text-sky-600 font-bold">Your Preferences</h2>

      <div class="mt-4 space-y-2">
        <div class="flex space-x-2">
          <div class="text-black font-bold">Email:</div>
          <div>{{ user.adPreference.email }}</div>
        </div>

        <div class="flex space-x-2">
          <div class="text-black font-bold">SMS:</div>
          <div>{{ user.adPreference.sms }}</div>
        </div>

        <div class="flex space-x-2">
          <div class="text-black font-bold">Phone:</div>
          <div>{{ user.adPreference.phone }}</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import {ChevronLeftIcon} from "@heroicons/vue/24/outline";

const {isAuthenticated, user} = useAuth()
</script>