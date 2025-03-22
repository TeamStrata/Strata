<script setup>
import { ref } from 'vue';
// placeholder user data 
let users =
    [
        { name: "Alistair", stuff: "hello" },
        { name: "Sebastian", stuff: "hello" }
    ]
var newUsername = ref("");
var newPassword = ref("");

</script>

<template>
    <h1>Manage Users</h1>
    <!-- new user area -->
    <div>
        <label for="newUser">Username:</label>
        <input id="newUser" v-model="newUsername">
        <label for="newPass">Password:</label>
        <input id="newPass" type="password" v-model="newPassword">
        <button class="border-2 rounded-md p-2 cursor-pointer">Add User</button>
    </div>
    <ul>
        <li v-for="u in users">
            {{ u.name }}
        </li>
    </ul>

</template>