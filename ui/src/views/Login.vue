<script setup>
import { ref } from 'vue';

var user = ref("");
var pass = ref("");
var error = ref(false);

async function login() {
    // submit user and pass to form
    // fetch("http://localhost:2300/login", Headers{}) //i forget how to do this
}

</script>

<template>
    <div>
        <label for="username">Username</label>
        <input id="usernam" v-model="user">
        <label for="password">Password</label>
        <input id="password" v-model="pass">
        <button @click="login">Login</button>
        <p v-if="error">Something went wrong</p>
    </div>
</template>