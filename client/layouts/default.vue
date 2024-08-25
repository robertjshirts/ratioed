<template>
  <Topbar />
  <div class="mx-auto max-w-screen-xl px-0 md:px-4">
    <Sidebar />
    <div class="pt-24 md:ps-72">
      <slot></slot>
    </div>
  </div>
  <PostModal />
</template>
