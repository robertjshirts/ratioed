<template>
  <div class="fixed flex h-screen w-72 flex-col">
    <div class="flex flex-col items-center">
      <img
        src="https://cdn.pixabay.com/photo/2015/10/05/22/37/blank-profile-picture-973460_960_720.png"
        alt=""
        class="w-24 rounded-full bg-contain"
      />
      <span class="mt-4 text-lg font-semibold">Generic User</span>
    </div>
    <div class="mt-8 flex w-full flex-col items-center">
      <NavLink to="/home" icon="carbon:home">Home</NavLink>
      <NavLink to="/following" icon="carbon:user-multiple" class="mt-2">
        Following
      </NavLink>
      <NavLink to="/profile" icon="carbon:user" class="mt-2">Profile</NavLink>
      <NavLink to="/settings" icon="carbon:settings" class="mt-2">
        Settings
      </NavLink>
    </div>
    <button class="w-full">Post</button>
  </div>
</template>
