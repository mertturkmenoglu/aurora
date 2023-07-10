<template>
  <div class="bg-black text-white py-2 flex items-center px-8 justify-center relative w-full">
    <span>Welcome to Aurora</span>
    <button class="absolute right-8" @click="$emit('close')">
      <XMarkIcon class="h-4 w-4 text-white"/>
    </button>
  </div>
</template>

<script lang="ts" setup>
import {XMarkIcon} from "@heroicons/vue/24/outline";

defineEmits<{
  close: []
}>()
</script>