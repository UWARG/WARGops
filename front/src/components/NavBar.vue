<template>
    <v-app-bar color="warg-blue">
        <div class="w-full flex flex-row justify-between">
            <img src="../assets/warg-logo.svg" class="w-1/12 text-white ml-8" />
            <div class="flex items-center justify-center">
                <router-link to="/" class="mx-2">WARG</router-link>
                |
                <router-link to="/dashboard" class="mx-2">DashBoard</router-link>
                |
                <router-link to="/about" class="mx-2">About</router-link>
                |
                <router-link to="/join" class="mx-2">Join</router-link>
            </div>
            <div class="flex items-center ">
                <v-menu open-on-hover>
                    <template v-slot:activator="{ props }">
                        <v-btn icon v-bind="props">
                            <v-avatar class="mx-4">
                                <v-img alt="Avatar"
                                    :src="`https://cdn.discordapp.com/avatars/${profileStore.getProfile.id}/${profileStore.getProfile.avatar}.png?size=1024`"
                                    id="menu-activator" />
                            </v-avatar>
                        </v-btn>
                    </template>
                    <v-card width="250" class="text-center flex flex-col items-center">
                        <v-card-title>
                            {{ profileStore.profile.username }}
                        </v-card-title>
                        <v-card-actions>
                            <v-btn color="red" variant="outlined" @click="logout">Logout</v-btn>
                        </v-card-actions>
                    </v-card>
                </v-menu>

            </div>
        </div>
    </v-app-bar>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { useTheme } from 'vuetify';
import { useProfileStore } from '../store/profile';
import { useRouter } from 'vue-router';

export default defineComponent({
    name: 'NavBar',
    props: {
        profile: {
            type: Boolean,
            default: true
        }
    },
    setup() {
        const theme = useTheme();
        const toggleTheme = () => theme.global.name.value = theme.global.current.value.dark ? 'wargLight' : 'wargDark';
        const profileStore = useProfileStore();
        const router = useRouter();
        const logout = () => {
            profileStore.logout();
            router.push('/login');
        };

        return { toggleTheme, theme, profileStore, logout };
    }
});
</script>

<style scoped>

</style>