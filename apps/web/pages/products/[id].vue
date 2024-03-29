<template>
  <Head>
    <title>Aurora / {{ product?.name }}</title>
  </Head>
  <div v-if='product && v' class='container mx-auto my-8'>
    <Breadcrumbs :items='breadcrumbLinks' />

    <div class='mt-8 w-full'>
      <div class='grid grid-cols-9 w-full'>
        <div class='col-span-6 p-8'>
          <!-- Image -->
          <swiper
            :modules='modules'
            :navigation='true'
            :spaceBetween='10'
            :style="{
                '--swiper-navigation-color': '#000',
                '--swiper-pagination-color': '#000',
              }"
            :thumbs='{ swiper: thumbsSwiper }'
            class='h-[48rem] w-[48rem] object-contain'
          >
            <swiper-slide v-for='image in product.variants.map((it) => it.image)'>
              <img :src='image.url' alt='' class='w-full h-full' />
            </swiper-slide>
          </swiper>
          <swiper
            :freeMode='true'
            :modules='thumbsModules'
            :scrollbar='{
                enabled: true,
                draggable: true
              }'
            :slidesPerView='5'
            :spaceBetween='50'
            :watchSlidesProgress='true'
            class=' object-cover mt-4 w-full'
            @swiper='setThumbsSwiper'
          >
            <swiper-slide v-for='image in product.variants.map((it) => it.image)'>
              <img :src='image.url' alt='' loading='lazy' />
            </swiper-slide>
          </swiper>

        </div>
        <div class='col-span-3 rounded-lg shadow-md flex flex-col p-4 bg-white h-min'>
          <!-- Message and fav -->
          <div class='flex items-center justify-between'>
            <span class='font-bold text-sm text-sky-600'>{{ productMessage }}</span>
            <ClientOnly fallback='' fallback-tag='span'>
              <FavoriteButton :productId='product.id' class='z-10' />
            </ClientOnly>
          </div>

          <!-- Brand -->
          <nuxt-link :to='`/products?brandId=${product.brand.id}`' class='mt-2'>
            <span class='underline text-sm mt-2 text-gray-600'>{{ product.brand.name }}</span>
          </nuxt-link>

          <!-- Rating -->
          <div class='mt-4 flex items-center space-x-1'>
            <template v-for='_ in reviewFilledStarsCount'>
              <FilledStarIcon class='w-3 h-3 text-sky-600' />
            </template>
            <template v-for='_ in reviewEmptyStarsCount'>
              <EmptyStarIcon class='w-4 h-4 text-gray-600' />
            </template>
            <span class='text-gray-600 text-xs ml-2'>(4.8)</span>
            <span class='text-gray-600 text-xs ml-2'>4238 reviews</span>
          </div>

          <!-- Name -->
          <h2 class='font-bold text-xl mt-4 text-midnight'>
            {{ product.name }}
          </h2>

          <!-- Price -->
          <div class='mt-4'>
            <span
              :class="clsx(
                    'font-bold text-green-600',
                    {
                        'text-3xl': v.currentPrice !== v.oldPrice,
                        'text-xl': v.currentPrice === v.oldPrice,
                    }
                )"
            >
              {{ v.currentPrice }}$
            </span>
            <span v-if='v.currentPrice !== v.oldPrice' class='font-light line-through ml-2'>
              {{ v.oldPrice }}$
            </span>
          </div>

          <div v-if='isCriticalStock'>
            <span class='text-red-600 text-sm mt-2'>Only {{ v.inventory }} left in stock</span>
          </div>

          <!-- Style -->
          <div v-if='styles.length > 0'>
            <div class='mt-4'>Styles:</div>
            <div class='mt-2 grid grid-cols-3 gap-4'>
              <div v-for='(style, idx) in styles'>
                <button :class="clsx(
                            'border border-sky-600 rounded py-1 px-2 text-gray-600 col-span-1 w-full',
                            {
                                'bg-sky-600 text-white': styleIndex === idx,
                                'hover:bg-sky-600 hover:text-white transition ease-in duration-200': styleIndex !== idx,
                            }
                        )" @click='styleIndex = idx'>
                  {{ style.name }}
                </button>
              </div>
            </div>
          </div>

          <!-- Size -->
          <div v-if='sizes.length > 0'>
            <div class='mt-4'>Sizes:</div>
            <div class='mt-2 grid grid-cols-4 gap-4'>
              <div v-for='(size, idx) in sizes'>
                <button :class="clsx(
                            'border border-sky-600 rounded py-1 px-2 text-gray-600 col-span-1 w-full',
                            {
                                'bg-sky-600 text-white': sizeIndex === idx,
                                'hover:bg-sky-600 hover:text-white transition ease-in duration-200': sizeIndex !== idx,
                            }
                        )" @click='sizeIndex = idx'>
                  {{ size.name }}
                </button>
              </div>
            </div>
          </div>

          <!-- Shipping -->
          <div class='mt-8'>
            <div class='flex items-center space-x-1'>
              <TruckIcon class='w-6 h-6 text-midnight' />
              <span class='font-bold text-midnight'>Shipping</span>
            </div>

            <p class='mt-2 text-gray-600'>
              <span v-if='isFreeShipping'>
                Free Shipping
              </span>
              <span v-else>
                Shipping: <span class='font-bold text-black'>{{ v.shippingPrice }}$</span>
              </span>
            </p>
            <p class='mt-2 text-gray-600 text-sm'>
              Get the item in {{ v.shippingTime }} with {{ v.shippingType }} shipping
            </p>
          </div>

          <!-- Address -->
          <div class='mt-8'>
            <div class='flex items-center space-x-1'>
              <MapPinIcon class='w-6 h-6 text-midnight' />
              <span class='font-bold text-midnight'>Address</span>
            </div>

            <p class='mt-2 text-gray-600 text-sm'>
              We will ship your product to your default address. You can change your address in your
              <nuxt-link class='underline text-sky-600' to='/my-account'>account settings</nuxt-link>
            </p>
          </div>

          <!-- Add to cart -->
          <div class='flex mt-8 items-end space-x-4'>
            <div class='flex flex-col'>
              <span class='text-gray-600 text-xs'>Qty:</span>
              <select class='py-2 bg-white border-b border-b-black'>
                <option>1</option>
                <option>2</option>
                <option>3</option>
                <option>4</option>
                <option>5</option>
                <option>6</option>
                <option>7</option>
                <option>8</option>
                <option>9</option>
                <option>10</option>
              </select>
            </div>
            <button
              class='flex px-4 py-2 text-white bg-sky-600 rounded-full text-center justify-center mt-4 w-full'
            >
              Add to cart
            </button>
          </div>
        </div>
      </div>

      <hr class='my-8' />

      <div class='mt-8'>
        <h2 class='font-bold text-xl text-midnight'>
          Description
        </h2>
        <p class='mt-4 text-gray-600'>
          {{ product.description }}
        </p>
      </div>

      <div class='mt-8'>
        <h2 class='font-bold text-xl text-midnight'>
          Serialized
        </h2>
        <pre class='mt-4 text-gray-600'>
          {{ JSON.stringify(product, null, 2) }}
        </pre>
      </div>

      <ProductCarousel
        :items='featuredProducts as Product[]'
        title='Featured products'
      />

    </div>
  </div>

  <div v-else>
    Cannot find product with id {{ route.params.id }}
  </div>
</template>

<script lang='ts' setup>
import { Product, ProductDto, ProductsDto, ProductVariant } from '~/utils/dto';
import { BASE_URL } from '~/utils/api';
import { TruckIcon, StarIcon as EmptyStarIcon } from '@heroicons/vue/24/outline';
import { StarIcon as FilledStarIcon, MapPinIcon } from '@heroicons/vue/24/solid';
import clsx from 'clsx';
import { useProductMessage } from '~/composables/useProductMessage';
import { Swiper as SwiperClass } from 'swiper';
import { Swiper, SwiperSlide } from 'swiper/vue';

// Import Swiper styles
import 'swiper/css';

import 'swiper/css/free-mode';
import 'swiper/css/navigation';
import 'swiper/css/thumbs';

// import required modules
import { FreeMode, Navigation, Thumbs, Scrollbar } from 'swiper/modules';
import { getBreadcrumbLinksFromCategories, getCategoriesFromProduct } from '~/utils/category';

const thumbsSwiper = ref<SwiperClass | null>(null);
const modules = [Navigation, Thumbs, FreeMode];
const thumbsModules = [...modules, Scrollbar];

function setThumbsSwiper(swiper: SwiperClass) {
  thumbsSwiper.value = swiper;
}

const route = useRoute();

const {
  data,
} = await useFetch<ProductDto>(`${BASE_URL}/products/${route.params.id}`);

const {
  data: featuredProductsData,
} = await useFetch<ProductsDto>(`${BASE_URL}/products/featured`);

const featuredProducts = featuredProductsData.value?.data ?? [];

const product: Product | undefined = data.value?.data;

const categories = computed(() => getCategoriesFromProduct(product));

const breadcrumbLinks = computed(() => getBreadcrumbLinksFromCategories(categories.value));

const productMessage = product ? useProductMessage(product) : '';

// v is a shorthand for variant
const v = ref<ProductVariant | null>(product?.defaultVariant || null);

const isCriticalStock = computed(() => {
  if (!v || !v.value) {
    return false;
  }

  return v.value.inventory < 50;
});

const isFreeShipping = computed(() => {
  if (!v || !v.value) {
    return false;
  }

  return v.value.shippingPrice === 0;
});

const reviewFilledStarsCount = computed(() => {
  if (!product) {
    return 0;
  }

  const rs = 4.8;

  return Math.round(rs);
});

const reviewEmptyStarsCount = computed(() => {
  if (!product) {
    return 5;
  }

  return 5 - reviewFilledStarsCount.value;
});

const styles = computed(() => {
  if (!product) {
    return [];
  }

  const styleIds = new Set<string>();

  for (const variant of product.variants) {
    styleIds.add(variant.style.id);
  }

  const idsArr = [...styleIds];

  return (idsArr.map(id => product.variants.find(s => s.style.id === id)!)).map(it => it.style);
});

const sizes = computed(() => {
  if (!product) {
    return [];
  }

  const sizeIds = new Set<string>();

  for (const variant of product.variants) {
    sizeIds.add(variant.size.id);
  }

  const idsArr = [...sizeIds];

  return (idsArr.map(id => product.variants.find(s => s.size.id === id)!)).map(it => it.size);
});
const styleIndex = ref(0);
const sizeIndex = ref(0);
</script>
