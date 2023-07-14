<template>
  <form class="mt-8 mx-auto max-w-sm">
    <h2 class="text-2xl font-medium">Sign in to Aurora</h2>

    <div class="mt-8 w-full">
      <label for="email" class="block text-sm font-medium text-gray-700 ml-1">Email</label>
      <input
          v-model="email"
          id="email"
          type="email"
          class="border border-black rounded py-1 px-1 w-full"
          placeholder="Email"
      />
    </div>

    <div class="mt-4 relative w-full">
      <label for="password" class="block text-sm font-medium text-gray-700 ml-1">Password</label>
      <input
          v-model="password"
          id="password"
          :type="showPassword ? 'text' : 'password'"
          class="border border-black rounded py-1 px-1 w-full"
          minlength="1"
          placeholder="Password"
      />
      <button @click="showPassword = !showPassword" type="button" class="absolute right-2 top-7">
        <EyeSlashIcon v-if="!showPassword" class="h-4 w-4 text-black" />
        <EyeIcon v-else class="h-4 w-4 text-black" />
      </button>
    </div>

    <div class="mt-16 w-full">
      <button type="submit" class="bg-black text-white py-2 px-4 rounded-lg w-full">Sign in</button>
    </div>

    <div class="mt-8 flex flex-col space-y-2">
      <nuxt-link to="/signup" class="group">
        <span>Don't have an account?</span>
        <span class="ml-1 group-hover:underline font-semibold">Create one</span>
      </nuxt-link>

      <nuxt-link to="/forgot-password" class="group">
        <span>Forgot your password?</span>
        <span class="ml-1 group-hover:underline font-semibold">Reset</span>
      </nuxt-link>
    </div>
  </form>
</template>

<script setup lang="ts">
import { EyeIcon, EyeSlashIcon} from "@heroicons/vue/24/outline";

const email = ref('')
const password = ref('')
const showPassword = ref(false)
</script>
