<template>
    <warg-page :shownav="false" class="">
        <div class="flex flex-col justify-center items-center h-screen w-full">
            <h1 class="text-4xl text-red-500">404</h1>
            <h2 class="text-2xl">Page not found</h2>
        </div>
    </warg-page>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import WargPage from '../components/WargPage.vue';
export default defineComponent({
    components: { WargPage },
    setup() {


        return {};
    }
});
</script>

<style scoped>

</style>