<template>
  <Head>
    <title>Aurora / Sign In</title>
  </Head>
  <form class="my-32 mx-auto max-w-sm" @submit.prevent="onFormSubmit">
    <h2 class="text-2xl font-medium">Sign in to <span class="text-sky-600">Aurora</span></h2>

    <div class="mt-8 w-full">
      <label class="block text-sm font-medium text-gray-700 ml-1" for="email">Email</label>
      <input
          id="email"
          v-model="email"
          class="border-b border-gray-400 py-1 px-1 w-full"
          placeholder="Email"
          type="email"
      />
    </div>

    <div class="mt-8 relative w-full">
      <label class="block text-sm font-medium text-gray-700 ml-1" for="password">Password</label>
      <input
          id="password"
          v-model="password"
          :type="showPassword ? 'text' : 'password'"
          class="border-b border-gray-400 py-1 px-1 w-full"
          minlength="1"
          placeholder="Password"
      />
      <button class="absolute right-2 top-7" type="button" @click="showPassword = !showPassword">
        <EyeSlashIcon v-if="!showPassword" class="h-4 w-4 text-black"/>
        <EyeIcon v-else class="h-4 w-4 text-black"/>
      </button>
    </div>

    <div class="mt-8 w-full">
      <button class="bg-sky-600 text-white py-2 px-4 rounded-full w-full" type="submit">Sign in</button>
    </div>

    <div class="mt-8 flex flex-col space-y-2 text-gray-600">
      <nuxt-link class="group" to="/signup">
        <span>Don't have an account?</span>
        <span class="ml-1 group-hover:underline font-semibold text-sky-600">Create one</span>
      </nuxt-link>

      <nuxt-link class="group" to="/forgot-password">
        <span>Forgot your password?</span>
        <span class="ml-1 group-hover:underline font-semibold text-sky-600">Reset</span>
      </nuxt-link>
    </div>
  </form>
</template>

<script lang="ts" setup>
import {EyeIcon, EyeSlashIcon} from "@heroicons/vue/24/outline";
import {api} from "~/utils/api";
import type {LoginResponseDto} from "~/utils/dto";
import {FetchError} from "ofetch";

const email = ref('')
const password = ref('')
const showPassword = ref(false)

const onFormSubmit = async () => {
  try {
    const {data} = await api<LoginResponseDto>('/auth/login', {
      method: 'POST',
      body: {
        email: email.value,
        password: password.value
      }
    })

    localStorage.setItem("accessToken", data.accessToken)
    localStorage.setItem("refreshToken", data.refreshToken)

    window.location.href = "/"
  } catch (e) {
    if (e instanceof FetchError) {
      console.error(e.response?._data.error)
    }
  }
}
</script>
