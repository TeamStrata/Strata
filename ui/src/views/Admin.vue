<script setup>
import { RouterLink, RouterView } from 'vue-router';


</script>

<template>
    <div class="flex flex-row max-w-[1000px] mx-auto h-screen p-20">
        <!-- sidebar -->
        <div class="bg-red-400 w-3xs h-full p-5">
            <RouterLink to="/admin/users">Users</RouterLink>
        </div>
        <div class="w-full h-full p-5">
            <!-- focus area -->
            <RouterView></RouterView>
        </div>
    </div>
</template>