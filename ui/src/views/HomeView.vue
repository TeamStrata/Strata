<script setup>
import { RouterLink } from 'vue-router';

</script>

<template>
  <!-- main separator -->
  <div class="flex flex-row h-screen">
    <!-- sidebar -->
    <div class="w-xs h-full bg-amber-100 flex flex-col">
      <!-- logo -->
      <div class="">
        Logo Placeholder
      </div>
      <!-- admin controls -->
      <div>
        <RouterLink to="/admin">Admin button</RouterLink>
      </div>
      <!-- navigation stuff -->
      <div class="flex flex-col">
        Navigation stuff
        <RouterLink to="/">Link</RouterLink>
      </div>
    </div>
    <!-- right side -->
    <div class="flex flex-col">
      <!-- header bar -->
      <div>
        <button>Logout</button>
      </div>
      <!-- main content area -->
      <div>
        THERE SHOULD BE CONTENT HERE BUT WE HAVE NO CONTENT TO PUT HERE YET
      </div>
    </div>
  </div>
</template>



<!-- <template>
  <header>
    <div class="wrapper">
      <nav>
        <RouterLink to="/">Home</RouterLink>
        <RouterLink to="/about">About</RouterLink>
        <RouterLink to="/login">login</RouterLink>
        <RouterLink to="/admin">Admin</RouterLink>
      </nav>
    </div>
  </header>

  <p>Huge content placeholder</p> 
</template> -->
