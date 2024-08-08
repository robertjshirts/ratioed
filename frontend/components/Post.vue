<script setup lang="ts">
import * as data from '~/assets/mocks/post.json';
</script>

<template>
  <div class="flex p-4 border border-slate-800 w-auto">
    <UAvatar :src="data.user.profilePicture" alt="avatar" size="lg" />
    <div class="ms-4">
      <div class="flex">
        <p class="font-bold">{{ data.user.username }}</p>
        <p class="ms-2 text-gray-400">·</p>
        <p class="ms-2 text-gray-400">{{ data.body.timestamp }}</p>
      </div>
      <pre class="mt-2 font-sans">{{ data.body.content }}</pre>
      <img
        class="mt-2 rounded-xl border border-slate-800"
        :src="data.body.attachment.src"
        alt="image failed to load"
      />
    </div>
  </div>
</template>
