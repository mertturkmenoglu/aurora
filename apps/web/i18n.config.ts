export default defineI18nConfig(() => ({
    legacy: false,
    locale: 'en',
    messages: {
        en: {
            header: {
                sliverTitle: "Welcome to Aurora",
                signIn: "Sign in",
                cart: "Cart",
                search: "Search",
            },
            home: {
                welcomeText: "Welcome",
            }
        },
        tr: {
            header: {
                sliverTitle: "Aurora'ya hoşgeldiniz",
                signIn: "Giriş Yap",
                cart: "Sepet",
                search: "Ara",
            },
            home: {
                welcomeText: "Hoşgeldiniz",
            }
        }
    }
}))