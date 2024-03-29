// https://nuxt.com/docs/api/configuration/nuxt-config
// @ts-nocheck
export default defineNuxtConfig({
    routeRules: {
        "/my-account/**": {ssr: false},
        "/": {ssr: false},
        "/products/**": {ssr: false},
        "/cart": {ssr: false}
    },
    devtools: {enabled: true},
    modules: [
        '@nuxtjs/tailwindcss',
        '@nuxtjs/i18n',
        '@nuxt/image',
        '@formkit/nuxt'
    ],
    i18n: {
        locales: ['en', 'tr'],
        defaultLocale: 'en',
        vueI18n: './i18n.config.ts',
    },
})
