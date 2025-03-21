<script setup>
import { RouterLink, RouterView } from 'vue-router';


</script>

<template>
    <div>
        <!-- sidebar -->
        <div>
            <RouterLink to="/admin/users">Users</RouterLink>
        </div>
        <!-- focus area -->
        <RouterView></RouterView>
    </div>
</template>