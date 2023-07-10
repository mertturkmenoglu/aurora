// https://nuxt.com/docs/api/configuration/nuxt-config
// @ts-nocheck
export default defineNuxtConfig({
    devtools: {enabled: true},
    modules: [
        '@nuxtjs/tailwindcss',
        '@nuxtjs/i18n',
    ],
    i18n: {
        locales: ['en', 'tr'],
        defaultLocale: 'en',
        vueI18n: './i18n.config.ts',
    },
})
