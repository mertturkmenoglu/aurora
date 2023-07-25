<template>
  <div v-if="items">
    <h2 class="font-bold text-xl text-black">
      {{ title }}
    </h2>
    <swiper
        :modules="swiperModules"
        :slides-per-view="5"
        :space-between="50"
        class="mt-8"
        navigation
    >
      <swiper-slide v-for="item in items">
        <ProductCard
            :hoverable="false"
            :product="item"
        />
      </swiper-slide>
    </swiper>
  </div>
</template>

<script lang="ts" setup>
import {Swiper, SwiperSlide} from "swiper/vue";
import {Navigation, Pagination, Scrollbar, A11y} from 'swiper/modules';

import 'swiper/css';
import 'swiper/css/navigation';
import 'swiper/css/pagination';
import 'swiper/css/scrollbar';
import {Product} from "~/utils/dto";

const {items, title} = defineProps<{
  items: Product[];
  title: string;
}>()

const swiperModules = [Navigation, Pagination, Scrollbar, A11y]

</script>
