<template>
  <form class="mt-8 mx-auto max-w-sm space-y-4">
    <h2 class="text-2xl font-medium">Sign in to Aurora</h2>
    <div class="w-full space-y-1">
      <label for="name" class="block text-sm font-medium text-gray-700 ml-1">Name</label>
      <input
          v-model="name"
          id="name"
          type="name"
          class="border border-black rounded py-1 px-1 w-full"
          placeholder="Name"
      />
    </div>

    <div class="w-full space-y-1">
      <label for="lastName" class="block text-sm font-medium text-gray-700 ml-1">Last Name</label>
      <input
          v-model="lastName"
          id="lastName"
          type="lastName"
          class="border border-black rounded py-1 px-1 w-full"
          placeholder="Last Name"
      />
    </div>

    <div class="w-full space-y-1">
      <label for="email" class="block text-sm font-medium text-gray-700 ml-1">Email</label>
      <input
          v-model="email"
          id="email"
          type="email"
          class="border border-black rounded py-1 px-1 w-full"
          placeholder="Email"
      />
    </div>

    <div class="relative w-full space-y-1">
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

    <div class="relative w-full space-y-1">
      <label for="repeatPassword" class="block text-sm font-medium text-gray-700 ml-1">Password</label>
      <input
          v-model="repeatPassword"
          id="repeatPassword"
          :type="showRepeatPassword ? 'text' : 'password'"
          class="border border-black rounded py-1 px-1 w-full"
          minlength="1"
          placeholder="Repeat Password"
      />
      <button @click="showRepeatPassword = !showRepeatPassword" type="button" class="absolute right-2 top-7">
        <EyeSlashIcon v-if="!showRepeatPassword" class="h-4 w-4 text-black" />
        <EyeIcon v-else class="h-4 w-4 text-black" />
      </button>
    </div>

    <div class="relative w-full space-y-1">
      <FormKit type="checkbox" label="Lorem ipsum" />
      </div>

    <div class="mt-16 w-full">
      <button type="submit" class="bg-black text-white py-2 px-4 rounded-lg w-full">Sign up</button>
    </div>

    <div class="mt-8 flex flex-col space-y-2">
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
const name = ref('')
const lastName = ref('')
const repeatPassword = ref('')

const password = ref('')
const showPassword = ref(false)
const showRepeatPassword = ref(false)
</script>
