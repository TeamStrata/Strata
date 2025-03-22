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


function addUser() {
    // submit user and pass to form
    fetch('http://localhost:8080/signup', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            username: newUsername.value,
            password: newPassword.value
        }),
    })
        .then(response => {
            if (!response.ok) {
                error = true;
                throw new Error('Something went wrong');
            }
            let data = response.json();
            console.log(data)
        })
        .catch(error => {
            console.error('Error:', error);
        });
}
</script>

<template>
    <h1 class="font-bold">Manage Users</h1>
    <!-- new user area -->
     add user placeholder
    <div class="flex flex-col">
        <div>
            <label for="newUser">Username:</label>
            <input id="newUser" v-model="newUsername">
        </div>
        <div>
            <label for="newPass">Password:</label>
            <input id="newPass" type="password" v-model="newPassword">
        </div>
        <button class="bg-blue-100 rounded-sm w-fit px-2 cursor-pointer" @click="addUser">Add User</button>
    </div>

    Users list placeholder:
    <ul>
        <li v-for="u in users">
            {{ u.name }}
        </li>
    </ul>

</template>

<style scoped>
input {
    background-color: rgb(230, 230, 230);
    margin: 1px;
}
</style>