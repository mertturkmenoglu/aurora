<template>
  <div v-if='cf' class='mx-auto flex flex-col items-center gap-8'>
    <div v-for='item in cf' :key='item.title'>
      <nuxt-link :href='item.href'>
        <img
          :alt='item.image.fields.description'
          :src='item.image.fields.file.url'
        />
      </nuxt-link>
    </div>
  </div>

</template>

<script lang='ts' setup>
import { CfBanner } from '~/utils/dto';

const cf = ref<CfBanner[]>();

onMounted(() => {
  const fn = async () => {
    const res = await getHomeBanners();
    const arr = res.items as unknown as { fields: CfBanner }[];
    cf.value = arr.filter((item) => item.fields.enabledPages.includes('home')).map((item) => item.fields);
  };

  fn();
});
</script>