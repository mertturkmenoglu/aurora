<template>
  <nav class="border-b border-midnight">
    <div class="w-full flex items-center justify-between">
      <PopoverGroup v-for="l0 in megaNavigationItems.items" class="flex relative">
        <MegaNavigationItem :item="l0"/>
      </PopoverGroup>
    </div>
  </nav>
</template>

<script lang="ts" setup>
import {megaNavigationItems} from './items'
import {PopoverGroup} from "@headlessui/vue";
</script>