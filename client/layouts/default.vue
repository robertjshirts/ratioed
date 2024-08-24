<template>
  <Topbar />
  <div class="mx-4 max-w-6xl lg:mx-auto">
    <Sidebar />
    <div class="pt-8">
      <slot></slot>
    </div>
  </div>
  <PostModal />
</template>
