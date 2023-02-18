<template>
    <warg-page :shownav="false" class="-m-8 overflow-clip">
        <div class="flex flex-col w-full items-center p-40 z-10 relative">
            <img src="../assets/warg-logo.svg" class="w-[50vw] md:w-[25vw] text-white" />
            <h3 class="text-4xl font-bold mt-[13vh] mb-[3vh]">Welcome to WARGops</h3>
            <h4 class="my-2">Sign in with Discord to Access All of WARGops</h4>
            <v-btn color="discord" class="text-m" @click="toDiscord">
                <img src="../assets/discord_logo.svg" width="23" height="23" class="mr-2 -ml-2" />
                Login with Discord
            </v-btn>
            <p class="my-2">Alternativly <span class="text-yellow-600"> Continue as Guest</span></p>
        </div>
        <img src="../assets/wave-bottom.svg" class="absolute -bottom-20 right-0 w-[120vw] z-0" />
        <img src="../assets/wave-top.svg" class="absolute -bottom-20 right-0 w-[120vw] z-0" />
    </warg-page>

</template>

<script lang="ts">
import { defineComponent, onBeforeMount } from 'vue';
import NavBar from '../components/NavBar.vue';
import { useProfileStore } from '../store/profile';
import { useRouter } from 'vue-router';
import WargPage from '../components/WargPage.vue';


export default defineComponent({
    name: 'Login',
    components: {
        NavBar,
        WargPage,
    },
    setup() {
        //TODO: MAKE Environment Variable
        const toDiscord = () => {
            window.location.href = 'http://localhost:8080/auth?provider=discord';
        };
        const router = useRouter();

        onBeforeMount(async () => {
            await useProfileStore().loadProfile();
            if (useProfileStore().profile.username != "") {
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