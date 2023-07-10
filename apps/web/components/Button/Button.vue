<template>
  <button @click="clickHandler">{{ value }}</button>
</template>

<script setup>
const value = ref('Button')

function clickHandler() {
  console.log(value.value)
}
</script>