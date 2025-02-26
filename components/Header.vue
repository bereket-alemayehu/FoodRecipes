<template>
  <header class="bg-black text-white p-4 flex items-center justify-between">
    <div class="flex items-center">
      <h1 class="text-2xl font-bold mr-6">Food.</h1>
      <nav class="flex space-x-4">
        <a href="#" class="hover:text-gray-300">Category</a>
        <a href="#" class="hover:text-gray-300">Recipes</a>
        <a href="#" class="hover:text-gray-300">Creators</a>
        <a href="#" class="hover:text-gray-300">Sign Up</a>
        <a href="#" class="hover:text-gray-300">Login</a>
      </nav>
    </div>
    <div class="flex items-center">
      <button class="mr-4 hover:text-gray-300">
        <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
        </svg>
      </button>
      <button class="hover:text-gray-300">
        <img v-if="userProfileImage" :src="userProfileImage" alt="User Profile" class="h-8 w-8 rounded-full" />
        <svg v-else xmlns="http://www.w3.org/2000/svg" class="h-8 w-8" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5.121 17.804A13.937 13.937 0 0112 16c2.5 0 4.847.796 6.879 2.196M15 15l2 2m-2-2a3 3 0 11-5.842-1.143 4.686 4.686 0 00-1.721 1.054M6 16.348a16.118 16.118 0 01-4 2.652m10.718-12.82A15.748 15.748 0 0119.5 8.296m-10.718 0a15.748 15.748 0 01-2.281 1.704" />
        </svg>
      </button>
    </div>
  </header>
</template>

<script setup>
import { ref } from 'vue';

// Replace with actual user profile image URL when available
const userProfileImage = ref(null); // or ref('/path/to/default-profile-image.jpg');
</script>

<style scoped>
/* You can add custom styles here if needed */
</style>