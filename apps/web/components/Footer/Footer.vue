<template>
  <footer
      class="bg-sky-50 w-full px-8 space-y-8 lg:space-y-0 lg:px-16 py-16 flex flex-col lg:flex-row justify-around absolute mt-32">
    <div class="w-64 lg:w-32 xl:w-64">
      <div class="font-bold text-sky-600 text-4xl">Aurora</div>
    </div>

    <div class="grid grid-cols-2 min-w-[40vw] gap-8">
      <nav>
        <div class="flex flex-col">
          <h3 class="text-sky-600 font-bold text-xl mb-4">Company</h3>
          <a class="text-gray-600 hover:text-gray-900" href="#">About</a>
          <a class="text-gray-600 hover:text-gray-900" href="#">Blog</a>
          <a class="text-gray-600 hover:text-gray-900" href="#">Jobs</a>
          <a class="text-gray-600 hover:text-gray-900" href="#">Press</a>
          <a class="text-gray-600 hover:text-gray-900" href="#">Partnerships</a>
        </div>
      </nav>

      <nav>
        <div class="flex flex-col">
          <h3 class="text-sky-600 font-bold text-xl mb-4">Community</h3>
          <a class="text-gray-600 hover:text-gray-900" href="#">Developers</a>
          <a class="text-gray-600 hover:text-gray-900" href="#">Support</a>
          <a class="text-gray-600 hover:text-gray-900" href="#">Events</a>
          <a class="text-gray-600 hover:text-gray-900" href="#">Blog</a>
          <a class="text-gray-600 hover:text-gray-900" href="#">Newsletter</a>
        </div>
      </nav>

      <nav>
        <div class="flex flex-col">
          <h3 class="text-sky-600 font-bold text-xl mb-4">More</h3>
          <a class="text-gray-600 hover:text-gray-900" href="#">Help</a>
          <a class="text-gray-600 hover:text-gray-900" href="#">Accessibility</a>
          <a class="text-gray-600 hover:text-gray-900" href="#">Trust & Safety</a>
          <a class="text-gray-600 hover:text-gray-900" href="#">Legal</a>
        </div>
      </nav>

      <nav>
        <div class="flex flex-col">
          <h3 class="text-sky-600 font-bold text-xl mb-4">Social</h3>
          <a class="text-gray-600 hover:text-gray-900" href="#">Twitter</a>
          <a class="text-gray-600 hover:text-gray-900" href="#">Instagram</a>
          <a class="text-gray-600 hover:text-gray-900" href="#">Facebook</a>
          <a class="text-gray-600 hover:text-gray-900" href="#">YouTube</a>
        </div>
      </nav>
    </div>

    <div class="flex flex-col max-w-md">
      <h3 class="text-sky-600 font-bold text-xl mb-4">Newsletter</h3>
      <p class="text-gray-600">Sign up to receive updates about new products and news about Aurora.</p>
      <input class="border border-gray-400 rounded-md p-2 mt-4" placeholder="Email Address" type="text">
      <button class="bg-sky-600 text-white rounded-full px-4 py-2 mt-4">Subscribe</button>
    </div>
  </footer>
</template>
<script lang="ts" setup>
</script>