<template>
    <v-app class="p-4">
        <nav-bar>
            <template v-slot:actions>

            </template>
        </nav-bar>

        <v-main class="flex align-center justify-center ">
            <v-card class="p-4">
                <v-card width="500" height="300" color="background-light-1"
                    class="flex flex-col justify-center items-center">
                    <h1 class="font-bold text-4xl mb-4">Welcome to WARG</h1>
                    <v-btn color="discord" class="text-m" @click="toDiscord">
                        <img src="../assets/discord_logo.svg" width="23" height="23" class="mr-2 -ml-2" />
                        Login with Discord
                    </v-btn>
                </v-card>
            </v-card>
        </v-main>

    </v-app>
</template>

<script lang="ts">
import { defineComponent, onBeforeMount, onMounted } from 'vue';
import NavBar from '../components/NavBar.vue';
import axios from 'axios';
import { useProfileStore } from '../store/profile';
import { useRouter } from 'vue-router';

export default defineComponent({
    name: 'Login',
    components: {
        NavBar,
    },
    setup() {
        //TODO: MAKE Environment Variable
        const toDiscord = () => {
            window.location.href = 'http://localhost:8080/auth?provider=discord';
        };
        const router = useRouter();

        onMounted(async () => {
            console.log("Loading Profile");
            await useProfileStore().loadProfile();
            if (useProfileStore().profile.Name) {
                console.log("Profile Loaded");
                router.push({ name: 'Home' });
            }

        });

        return {
            toDiscord,
        };
    }
});
</script>

<style scoped>

</style>