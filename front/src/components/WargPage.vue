<template>
    <v-app class="p-4" :scrollable="false">
        <nav-bar v-if="shownav"></nav-bar>
        <v-main >
            <slot></slot>
        </v-main>
    </v-app>

</template>

<script lang="ts">
import { defineComponent, watch, ref } from 'vue';
import NavBar from './NavBar.vue';
import { useProfileStore } from '../store/profile';

export default defineComponent({
    components: { NavBar },
    props: {
        shownav: {
            type: Boolean,
            default: true
        }
    },
    setup() {
        const profileStore = useProfileStore();
        const alert = ref();

        

        return { alert };
    }
});
</script>

<style scoped>

</style>