<template>
  <div v-if="isAuthenticated && user" class="container mx-auto mt-16">
    <h2 class="text-2xl font-bold">Your Shopping Cart</h2>
    <div v-if="cart" class="mt-16">
      <div v-for="item in cart.items">
        <CartItem :item="item"/>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import {Cart, CartDto} from "~/utils/dto";

const {isAuthenticated, user} = useAuth()

const {data} = await useFetch<CartDto>(`${BASE_URL}/cart`, {
  headers: {
    'x-access-token': localStorage.getItem('accessToken') || '',
    'x-refresh-token': localStorage.getItem('refreshToken') || '',
  },
})

const cart: Cart | undefined = data.value?.data
</script>