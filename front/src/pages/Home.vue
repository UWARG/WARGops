<template>
    <v-app class="p-4">
        <nav-bar>
            <template #button>

            </template>
        </nav-bar>

        <v-main>
            <div class="flex">
                <div class="flex-1">
                    <v-card class="p-4">
                        <div class="flex">
                            <v-card class="p-4 flex-1 mr-2" color="background-light-1">
                                <div class="flex justify-between">
                                    <v-card-title class="text-3xl font-bold mb-2">
                                        Accounts
                                    </v-card-title>
                                    <v-dialog v-model="dialog" width="800">
                                        <template v-slot:activator="{ props }">
                                            <v-btn v-bind="props" color="primary" variant="tonal"> Create New Account</v-btn>
                                        </template>
                                        <account-modal @closeModal="dialog = false" />
                                    </v-dialog>
                                </div>

                                <v-text-field density="compact" variant="solo" label="Search Accounts"
                                    append-inner-icon="mdi-magnify" single-line hide-details v-model="filter" />
                                <div>
                                    <div class="cursor-pointer"
                                        v-for="(account, index) in accountStore.getFilteredAccounts(filter)"
                                        :key="account.id"
                                        @click="activeAccount = accountStore.getAccountById(account.id)">
                                        <div class="flex justify-between items-center">
                                            <div class="my-4"> {{ account.name }}</div>
                                            <v-chip color="green" @click="switchTransaction(account.id)">Open
                                                Transactions</v-chip>
                                        </div>
                                        <v-divider v-if="index != accountStore.accounts.length - 1" />
                                    </div>
                                </div>

                            </v-card>

                            <v-card class="p-4 flex-1 ml-2" color="background-light-1" v-if="activeAccount">
                                <v-card-title class="text-3xl font-bold mb-2 text-center">
                                    {{ activeAccount.name }}
                                </v-card-title>

                                <div v-for="(key) in Object.keys(activeAccount)" :key="key" class="flex">
                                    <div class="mb-2"> <span class="font-bold"> {{ key }}</span>: {{
                                        activeAccount[key]
                                    }}</div>
                                </div>
                            </v-card>
                        </div>
                    </v-card>
                </div>
            </div>
        </v-main>
    </v-app>

</template>

<script lang="ts">
import { defineComponent, ref, onBeforeMount } from 'vue';
import { useTheme } from 'vuetify';

import AccountModal from '../components/AccountModal.vue';
import { useRouter } from 'vue-router';
import NavBar from '../components/NavBar.vue';
import { useAccountStore } from '../store/accounts';
import { Account } from '../types';
import { useProfileStore } from '../store/profile';


export default defineComponent({
    components: { AccountModal, NavBar },
    setup() {
        const accountStore = useAccountStore();
        const theme = useTheme();
        const dialog = ref(false);
        const toggleTheme = () => theme.global.name.value = theme.global.current.value.dark ? 'wargLight' : 'wargDark';

        const router = useRouter();

        const switchTransaction = (account_id: string) => {
            console.log(account_id);
            router.push({ name: 'Transaction', params: { account_id } });
        };


        const filter = ref<string>('');

        const activeAccount = ref<Account | undefined>();
        return { dialog, toggleTheme, theme, accountStore, switchTransaction, activeAccount, filter };
    }
});
</script>

<style scoped>

</style>