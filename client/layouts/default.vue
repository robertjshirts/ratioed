<template>
  <Topbar />
  <div class="mx-auto max-w-6xl">
    <Sidebar />
    <div class="pl-72 pt-8">
      <slot></slot>
    </div>
  </div>
  <PostModal />
</template>
