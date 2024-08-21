// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: { enabled: true },
  compatibilityDate: "2024-08-07",
  modules: [
    "@nuxtjs/tailwindcss",
    "@nuxt/icon",
    "@nuxtjs/google-fonts",
    "@nuxtjs/supabase",
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
});
