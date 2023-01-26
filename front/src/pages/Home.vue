<template>
    <v-app class="p-4">
        <v-app-bar>
            <div class=" w-full flex justify-between">
                <div class="text-3xl font-bold ml-8">
                    WARG ✈️
                </div>
                <div class="flex items-center">
                    <v-btn :icon="theme.global.name.value === 'wargDark' ? 'mdi-weather-sunny' : 'mdi-weather-night'"
                        @click="toggleTheme" />

                    <v-dialog v-model="dialog" width="800">
                        <template v-slot:activator="{ props }">
                            <v-btn v-bind="props" color="primary" class=""> Create New Account</v-btn>
                        </template>
                        <Modal @closeModal="dialog = false" />
                    </v-dialog>
                </div>
<<<<<<< Updated upstream

            </div>

        </v-app-bar>

        <v-main>

            <div class="flex">
                <div class="flex-1">
                    <v-card class="p-4 mr-4">
                        <v-card class="p-4" color="background-light-1">
=======
            </div>
        </v-app-bar>

        <v-main>
            <div class="flex">
                <div class="flex-1">
                    <v-card class="p-2" color="background-light-1">
                        <v-card class="p-4">
>>>>>>> Stashed changes
                            <v-card-title class="text-3xl font-bold">
                                Accounts
                            </v-card-title>
                            <v-expansion-panels color="background-dark-1">
<<<<<<< Updated upstream
                                <v-expansion-panel v-for="(account) in accounts" :key="account.id" :title="account.name">
=======
                                <v-expansion-panel v-for="(account) in accounts" :key="account.id"
                                    :title="account.name">
>>>>>>> Stashed changes
                                    {{ account }}
                                </v-expansion-panel>
                            </v-expansion-panels>
                        </v-card>
<<<<<<< Updated upstream


                    </v-card>

                </div>

            </div>



=======
                    </v-card>
                </div>
            </div>
>>>>>>> Stashed changes
        </v-main>
    </v-app>

</template>

<script lang="ts">
import { defineComponent, onMounted, ref } from 'vue';
import { useTheme } from 'vuetify';

import axios from 'axios';
import Modal from '../components/Modal.vue';

export default defineComponent({
    components: { Modal },
    setup() {
        const theme = useTheme();
        const dialog = ref(false);
        const toggleTheme = () => theme.global.name.value = theme.global.current.value.dark ? 'wargLight' : 'wargDark';
        const accounts = ref<any[]>([]);
        onMounted(() => {
            //@ts-ignore
            axios.get(`http://localhost:8080/accounts`)
                .then((res) => {
                    accounts.value = res.data;
                })
                .catch((err) => {
                    console.log(err);
                });
        });
<<<<<<< Updated upstream



=======
>>>>>>> Stashed changes
        return { dialog, toggleTheme, theme, accounts };
    }
});
</script>

<style scoped>

</style>