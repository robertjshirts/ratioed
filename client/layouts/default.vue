<template>
  <Topbar />
  <div class="relative mx-auto mt-24 max-w-screen-xl">
    <Sidebar />
    <div class="md:ps-72">
      <slot></slot>
    </div>
  </div>
  <PostModal />
</template>
