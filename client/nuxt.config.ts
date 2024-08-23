// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: {
    enabled: true,

    timeline: {
      enabled: true,
    },
  },
  compatibilityDate: "2024-08-07",
  modules: [
    "@nuxtjs/tailwindcss",
    "@nuxt/icon",
    "@nuxtjs/google-fonts",
    "@nuxtjs/supabase",
    "@nuxt/image",
    "@pinia/nuxt",
  ],
  supabase: {
    redirect: false,
    types: "~/types/database.d.ts",
  },
  googleFonts: {
    families: {
      Outfit: true,
    },
  },
  icon: {
    clientBundle: {
      icons: [
        "ph:house",
        "ph:house-fill",
        "ph:user",
        "ph:user-fill",
        "ph:gear",
        "ph:gear-fill",
      ],
    },
  },
  app: {
    pageTransition: { name: "page", mode: "out-in" },
  },
});
