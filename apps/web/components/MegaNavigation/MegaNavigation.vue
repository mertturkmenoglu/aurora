<template>
  <nav class="border-b border-midnight">
    <div class="w-full flex items-center justify-between">
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

const serialized = sessionStorage.getItem("megaNavigation")

if (serialized) {
  items.value = JSON.parse(serialized)
} else {
  const {data} = await useFetch<{ data: MegaNavigationItems }>("http://localhost:5000/api/v1/categories")

  if (data.value) {
    items.value = data.value.data
    sessionStorage.setItem("megaNavigation", JSON.stringify(data.value.data))
  }
}


</script>