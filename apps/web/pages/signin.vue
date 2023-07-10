<template>
  <form class="mt-8 mx-auto max-w-sm">
    <h2 class="text-2xl font-medium">Sign in to Aurora</h2>

    <div class="mt-8 w-full">
      <label class="block text-sm font-medium text-gray-700 ml-1" for="email">Email</label>
      <input
          id="email"
          v-model="email"
          class="border border-black rounded py-1 px-1 w-full"
          placeholder="Email"
          type="email"
      />
    </div>

    <div class="mt-4 relative w-full">
      <label class="block text-sm font-medium text-gray-700 ml-1" for="password">Password</label>
      <input
          id="password"
          v-model="password"
          :type="showPassword ? 'text' : 'password'"
          class="border border-black rounded py-1 px-1 w-full"
          minlength="1"
          placeholder="Password"
      />
      <button class="absolute right-2 top-7" type="button" @click="showPassword = !showPassword">
        <EyeSlashIcon v-if="!showPassword" class="h-4 w-4 text-black"/>
        <EyeIcon v-else class="h-4 w-4 text-black"/>
      </button>
    </div>

    <div class="mt-16 w-full">
      <button class="bg-black text-white py-1 px-4 rounded w-full" type="submit">Sign in</button>
    </div>

    <div class="mt-8 flex flex-col space-y-2">
      <nuxt-link class="group" to="/signup">
        <span>Don't have an account?</span>
        <span class="ml-1 group-hover:underline font-semibold">Create one</span>
      </nuxt-link>

      <nuxt-link class="group" to="/forgot-password">
        <span>Forgot your password?</span>
        <span class="ml-1 group-hover:underline font-semibold">Reset</span>
      </nuxt-link>
    </div>
  </form>
</template>

<script lang="ts" setup>
import {EyeIcon, EyeSlashIcon} from "@heroicons/vue/24/outline";

const email = ref('')
const password = ref('')
const showPassword = ref(false)
</script>
