<template>
  <div v-if="isAuthenticated && user" class="mx-auto my-32 container">
    <h1 class="text-4xl font-bold text-gray-900">
      Welcome, <span class="text-sky-600">{{ user.fullName }}</span>
    </h1>

    <div class="mt-32 grid grid-cols-2 lg:grid-cols-4 gap-8">
      <div v-for="{ link, icon: Icon, title } in items">
        <nuxt-link :to="link" class="flex">
          <div
              class="p-8 flex flex-col items-center group w-full bg-white text-sky-600 border border-sky-600 rounded-lg hover:bg-sky-600 hover:text-white transition ease-in duration-200">
            <component :is="Icon" class="h-8 w-8"/>
            <div class="font-semibold mt-4 text-center line-clamp-1">{{ title }}</div>
          </div>
        </nuxt-link>
      </div>

      <button class="flex w-full" @click.prevent="logout">
        <div
            class="p-8 flex flex-col items-center group w-full bg-white text-sky-600 border border-sky-600 rounded-lg hover:bg-sky-600 hover:text-white transition ease-in duration-200">
          <ArrowLeftOnRectangleIcon class="h-8 w-8"/>
          <div class="font-semibold mt-4 text-center line-clamp-1">Logout</div>
        </div>
      </button>
    </div>
  </div>
</template>

<script lang="ts" setup>
import {
  IdentificationIcon,
  Cog6ToothIcon,
  MapPinIcon,
  TruckIcon,
  ArrowLeftOnRectangleIcon,
  BanknotesIcon,
  HeartIcon,
  LockClosedIcon,
} from "@heroicons/vue/24/outline";

const {isAuthenticated, user} = useAuth()

const items = [
  {
    title: 'Overview',
    icon: IdentificationIcon,
    link: '/my-account/overview',
    isButton: false,
  },
  {
    title: 'Preferences',
    icon: Cog6ToothIcon,
    link: '/my-account/preferences',
    isButton: false,
  },
  {
    title: 'Addresses',
    icon: MapPinIcon,
    link: '/my-account/addresses',
    isButton: false,
  },
  {
    title: 'Orders',
    icon: TruckIcon,
    link: '/my-account/orders',
    isButton: false,
  },
  {
    title: 'Billing',
    icon: BanknotesIcon,
    link: '/my-account/billing',
    isButton: false,
  },
  {
    title: 'Favorites',
    icon: HeartIcon,
    link: '/favorites',
    isButton: false,
  },
  {
    title: "Login & Security",
    icon: LockClosedIcon,
    link: "/my-account/security",
    isButton: false,
  },
]

function logout() {
  if (localStorage) {
    localStorage.removeItem('accessToken')
    localStorage.removeItem('refreshToken')
    window.location.href = '/'
  }

}
</script>