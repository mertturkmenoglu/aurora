<template>
  <nav class="border-b border-midnight">
    <div class="w-full flex items-center space-x-16">
      <PopoverGroup v-for="l0 in items.items" class="flex relative">
        <MegaNavigationItem :item="l0"/>
      </PopoverGroup>
    </div>
  </nav>
</template>

<script lang="ts" setup>
import {MegaNavigationItems} from './items'
import {PopoverGroup} from "@headlessui/vue";

const items = ref<MegaNavigationItems>({
  items: []
})

onMounted(() => {
  const serialized = sessionStorage.getItem("megaNavigation")

  if (serialized) {
    items.value = JSON.parse(serialized)
    return
  }

  api<{
    data: MegaNavigationItems
  }>("/categories")
      .then(({data}) => {
        items.value = data
        sessionStorage.setItem("megaNavigation", JSON.stringify(data))
      })
})
</script>