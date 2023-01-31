<template>
    <v-app-bar>
        <div class=" w-full flex justify-between">
            <div class="text-3xl font-bold ml-8">
                WARG ✈️
            </div>
            <div class="flex items-center">
                <v-btn :icon="theme.global.name.value === 'wargDark' ? 'mdi-weather-sunny' : 'mdi-weather-night'"
                    @click="toggleTheme" />
                <slot name="button"></slot>



                <v-menu open-on-hover>
                    <template v-slot:activator="{ props }">
                        <v-btn icon v-bind="props">
                            <v-avatar class="mx-4">
                                <v-img alt="Avatar" :src="profileStore.profile.AvatarURL" id="menu-activator" />
                            </v-avatar>
                        </v-btn>
                    </template>
                    <v-card width="250" class="text-center flex flex-col items-center">
                        <v-card-title>
                            {{ profileStore.profile.Name }}
                        </v-card-title>
                        <v-card-text>
                            {{ profileStore.profile.Email }}
                        </v-card-text>
                        <v-card-actions>
                            <v-btn color="red" variant="outlined" @click="profileStore.logout">Logout</v-btn>
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

        return { toggleTheme, theme, profileStore };
    }
});
</script>

<style scoped>

</style>