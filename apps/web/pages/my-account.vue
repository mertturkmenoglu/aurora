<template>
  <div v-if="isAuthenticated && user" class="mx-auto my-32">
    <h1 class="text-4xl font-bold text-gray-900">
      Welcome, <span class="text-sky-600">{{ user.fullName }}</span>
    </h1>

    <div class="mt-8">
      <h2 class="text-xl font-medium">Your account details</h2>
      <p class="text-gray-600">Email: {{ user.email }}</p>
      <p class="text-gray-600">Phone: {{ user.phone }}</p>
      <div>Ad Preferences:</div>
      <p class="text-gray-600 ml-4">Email: {{ user.adPreference.email }}</p>
      <p class="text-gray-600 ml-4">Phone: {{ user.adPreference.phone }}</p>
      <p class="text-gray-600 ml-4">Sms: {{ user.adPreference.sms }}</p>
      <div>Addresses:</div>
      <div v-for="address in user.addresses">
        <p class="text-gray-600 ml-4">Phone: {{ address.phone }}</p>
        <p class="text-gray-600 ml-4">City: {{ address.city }}</p>
        <p class="text-gray-600 ml-4">State: {{ address.state }}</p>
        <p class="text-gray-600 ml-4">Zip: {{ address.zipCode }}</p>
      </div>
    </div>

    <button class="mt-8 bg-sky-600 text-white px-4 py-2 rounded-full" @click="logout">Logout</button>
  </div>

</template>
<script lang="ts" setup>
const {isAuthenticated, user} = useAuth()

function logout() {
  if (localStorage) {
    localStorage.removeItem('accessToken')
    localStorage.removeItem('refreshToken')
    window.location.href = '/'
  }

}
</script>