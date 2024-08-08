<script setup lang="ts">
const props = defineProps()
</script>

<template>
  <div class="flex p-4 border border-slate-800 w-auto">
    <UAvatar src="" alt="avatar" size="lg" />
    <div class="ms-4">
      <div class="flex">
        <p class="font-bold">username</p>
        <p class="ms-2 text-gray-400">·</p>
        <p class="ms-2 text-gray-400">timestamp</p>
      </div>
      <pre class="mt-2 font-sans">content</pre>
      <img
        class="mt-2 rounded-xl border border-slate-800"
        src=""
        alt="image failed to load"
      />
    </div>
  </div>
</template>
