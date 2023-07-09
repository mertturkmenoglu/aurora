<template>
  <div class="bg-black text-white py-2 flex items-center px-8 justify-center relative w-full">
    <span>Welcome to Aurora</span>
    <button @click="$emit('close')" class="absolute right-8">
      <XMarkIcon class="h-4 w-4 text-white"/>
    </button>
  </div>
</template>

<script setup lang="ts">
import {XMarkIcon} from "@heroicons/vue/24/outline";

defineEmits<{
  close: []
}>()
</script>