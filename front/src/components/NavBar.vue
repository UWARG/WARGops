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
            </div>
        </div>
    </v-app-bar>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { useTheme } from 'vuetify';

export default defineComponent({
    name: 'NavBar',
    setup() {
        const theme = useTheme();
        const toggleTheme = () => theme.global.name.value = theme.global.current.value.dark ? 'wargLight' : 'wargDark';


        return { toggleTheme, theme};
    }
});
</script>

<style scoped>

</style>