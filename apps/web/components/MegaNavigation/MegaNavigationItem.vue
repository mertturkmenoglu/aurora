<template>
  <Popover v-slot="{ open }">
    <PopoverButton
        :class="clsx(
           'group inline-flex items-center',
           'rounded-t-md py-2 px-2',
           'text-base font-medium text-midnight',
           'transition duration-150 ease-in-out hover:bg-gray-100',
           'focus:outline-none focus-visible:ring-2 focus-visible:ring-white focus-visible:ring-opacity-75',
       )"
    >
      <span>{{ item.name }}</span>
    </PopoverButton>

    <transition
        enter-active-class="transition duration-200 ease-out"
        enter-from-class="translate-y-1 opacity-0"
        enter-to-class="translate-y-0 opacity-100"
        leave-active-class="transition duration-150 ease-in"
        leave-from-class="translate-y-0 opacity-100"
        leave-to-class="translate-y-1 opacity-0"
    >
      <PopoverPanel
          class="fixed z-10 mt-3 mx-4 left-0 right-0 px-4"
      >
        <div
            class="overflow-hidden rounded-lg shadow-lg ring-1 ring-black ring-opacity-5"
        >
          <div class="relative p-8 bg-white flex w-full">
            <div class="flex-1">
              <a
                  v-for="l1 in item.items"
                  :key="l1.name"
                  :class="clsx(
                    '-m-3 flex items-center rounded-lg p-4', 'transition duration-150 ease-in-out',
                    'hover:bg-gray-50',
                    'focus:outline-none focus-visible:ring focus-visible:ring-sky-500 focus-visible:ring-opacity-50'
                )"
                  :href="`/category/${l1.id}`"
              >
                <div class="ml-4">
                  <p class="text-sm font-medium text-gray-900">
                    {{ l1.name }}
                  </p>
                </div>
              </a>
            </div>
            <div class="flex-1 hidden sm:flex">

            </div>
          </div>
          <div class="bg-gray-50 p-4">
            <nuxt-link
                :href="`/category/${item.id}`"
                class="flow-root rounded-md px-2 py-2 transition duration-150 ease-in-out hover:bg-gray-100 focus:outline-none focus-visible:ring focus-visible:ring-sky-500 focus-visible:ring-opacity-50"
            >
              <span class="block text-sm text-gray-500">
                See all {{ item.name }} products
              </span>
            </nuxt-link>
          </div>
        </div>
      </PopoverPanel>
    </transition>
  </Popover>
</template>
<script lang="ts" setup>
import {Popover, PopoverButton, PopoverPanel} from "@headlessui/vue";
import clsx from "clsx";
import {L0Item} from "./items";

defineProps<{
  item: L0Item;
}>()
</script>