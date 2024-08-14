<template>
  <div class="fixed flex h-screen w-56 flex-col">
    <div class="w-24">
      <img
        src="https://cdn.pixabay.com/photo/2015/10/05/22/37/blank-profile-picture-973460_960_720.png"
        alt=""
        class="rounded-full bg-contain"
      />
    </div>
    <p class="mt-4 text-lg font-semibold">Generic User</p>
    <div class="mt-4 flex w-full flex-col">
      <NavLink to="/home" icon="carbon:home">Home</NavLink>
      <NavLink to="/profile" icon="carbon:user" class="mt-2">Profile</NavLink>
      <NavLink to="/settings" icon="carbon:settings" class="mt-2">
        Settings
      </NavLink>
    </div>
  </div>
</template>
