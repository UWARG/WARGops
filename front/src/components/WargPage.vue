<template>
    <v-app class="p-4">
        <nav-bar></nav-bar>
        <v-main :scrollable="false" >
            <slot></slot>
        </v-main>
        <!-- <div class="flex justify-center items-center">

            <div class="w-[30vw]">
                <v-alert icon="mdi-information" density="comfortable" type="error" variant="tonal" ref="alert" id="alert" class="opacity-0">
                    Test
                </v-alert>
            </div>
        </div> -->
    </v-app>

</template>

<script lang="ts">
import { defineComponent, watch, ref } from 'vue';
import NavBar from './NavBar.vue';
import { useProfileStore } from '../store/profile';
import { alertUP, alertDOWN } from '../animations';

export default defineComponent({
    components: { NavBar },
    setup() {
        const profileStore = useProfileStore();
        const alert = ref();

        watch(profileStore.getAlert, () => {
            console.log("alert changed:" + profileStore.getAlert.alert )
            if (profileStore.getAlert.alert) {
                alertUP(document.getElementById("alert")!);
            } else {
                alertDOWN(document.getElementById("alert")!);
            }
        });

        return { alert };
    }
});
</script>

<style scoped>

</style>