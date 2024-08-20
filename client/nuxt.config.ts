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
    types: "~/types/database.types.ts",
  },
  googleFonts: {
    families: {
      Outfit: true,
    },
  },
});
