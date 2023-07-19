<template>
  <form class="my-32 mx-auto max-w-sm space-y-4">
    <h2 class="text-2xl font-medium">Create an <span class="text-sky-600">Aurora</span> account</h2>
    <div class="w-full space-y-1">
      <label class="block text-sm font-medium text-gray-700 ml-1" for="name">Name</label>
      <input
          id="name"
          v-model="name"
          class="border border-black rounded py-1 px-1 w-full"
          placeholder="Name"
          type="text"
      />
    </div>

    <div class="w-full space-y-1">
      <label class="block text-sm font-medium text-gray-700 ml-1" for="lastName">Last Name</label>
      <input
          id="lastName"
          v-model="lastName"
          class="border border-black rounded py-1 px-1 w-full"
          placeholder="Last Name"
          type="text"
      />
    </div>

    <div class="w-full space-y-1">
      <label class="block text-sm font-medium text-gray-700 ml-1" for="email">Email</label>
      <input
          id="email"
          v-model="email"
          class="border border-black rounded py-1 px-1 w-full"
          placeholder="Email"
          type="email"
      />
    </div>

    <div class="relative w-full space-y-1">
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

    <div class="relative w-full space-y-1">
      <label class="block text-sm font-medium text-gray-700 ml-1" for="repeatPassword">Password</label>
      <input
          id="repeatPassword"
          v-model="repeatPassword"
          :type="showRepeatPassword ? 'text' : 'password'"
          class="border border-black rounded py-1 px-1 w-full"
          minlength="1"
          placeholder="Repeat Password"
      />
      <button class="absolute right-2 top-7" type="button" @click="showRepeatPassword = !showRepeatPassword">
        <EyeSlashIcon v-if="!showRepeatPassword" class="h-4 w-4 text-black"/>
        <EyeIcon v-else class="h-4 w-4 text-black"/>
      </button>
    </div>

    <div class="relative w-full space-y-1">
      <FormKit label="Send me marketing emails" type="checkbox"/>
    </div>

    <div class="mt-16 w-full">
      <button class="bg-sky-600 text-white py-2 px-4 rounded-lg w-full" type="submit">Sign up</button>
    </div>

    <div class="mt-8 flex flex-col space-y-2">
      <nuxt-link class="group" to="/forgot-password">
        <span>Forgot your password?</span>
        <span class="ml-1 group-hover:underline font-semibold text-sky-600">Reset</span>
      </nuxt-link>
    </div>
  </form>
</template>

<script lang="ts" setup>
import {EyeIcon, EyeSlashIcon} from "@heroicons/vue/24/outline";

const email = ref('')
const name = ref('')
const lastName = ref('')
const repeatPassword = ref('')

const password = ref('')
const showPassword = ref(false)
const showRepeatPassword = ref(false)
</script>
