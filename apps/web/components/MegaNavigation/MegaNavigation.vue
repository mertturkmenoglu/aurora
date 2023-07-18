<template>
  <nav>
    <div class="w-full flex items-center justify-between">
      <div v-for="l0 in megaNavigationItems.items" class="flex">
        <MegaNavigationItem :item="l0"/>
      </div>
    </div>
  </nav>
</template>

<script lang="ts" setup>
import {megaNavigationItems} from './items'
</script>