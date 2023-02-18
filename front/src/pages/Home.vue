<template>

    <warg-page>
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
                                        <v-btn v-bind="props" color="primary" variant="tonal"> Create New
                                            Account</v-btn>
                                    </template>
                                    <account-modal @closeModal="dialog = false" />
                                </v-dialog>
                            </div>

                            <v-text-field density="compact" variant="solo" label="Search Accounts"
                                append-inner-icon="mdi-magnify" single-line hide-details v-model="filter" />
                            <div>
                                <div class="cursor-pointer"
                                    v-for="(account, index) in accountStore.getFilteredAccounts(filter)"
                                    :key="account.id" @click="activeAccount = accountStore.getAccountById(account.id)">
                                    <div class="flex justify-between items-center">
                                        <div class="my-4"> {{ account.name }}</div>
                                        <div>
                                            <v-btn variant="tonal" size="small" @click="openNewTransaction">
                                                New Transaction
                                                <v-icon icon="mdi-plus" class="ml-1" />
                                            </v-btn>
                                            <v-btn variant="tonal" size="small" class="ml-2">
                                                View All Transactions
                                                <v-icon icon="mdi-book-open-variant" class="ml-1" />
                                            </v-btn>
                                        </div>
                                    </div>
                                    <v-divider v-if="index != accountStore.accounts.length - 1" />
                                </div>
                            </div>

                        </v-card>

                        <v-card class="p-4 flex-1 ml-2 " color="background-light-1" v-if="activeAccount">
                            <v-card-title class="text-3xl font-bold mb-2 flex mx-8">
                                <span class="underline w-full text-center ml-12"> {{ activeAccount.name }}</span>
                                <div>
                                    <v-chip class="" :color="activeAccount.active ? 'green' : 'red'">
                                        {{ activeAccount.active ? 'Active' : 'Not Active' }}
                                    </v-chip>
                                </div>

                            </v-card-title>

                            <div class="flex text-center">
                                <div class="mx-2 flex-1">
                                    <h1 class="text-lg">Created Transactions</h1>
                                    <span>lorem</span>
                                </div>
                                <v-divider vertical></v-divider>
                                <div class="mx-2 flex-1">
                                    <h1 class="text-lg">Pending Transactions</h1>
                                    <span>lorem</span>
                                </div>
                                <v-divider vertical></v-divider>
                                <div class="mx-2 flex-1">
                                    <h1 class="text-lg">Paid Transactions</h1>
                                    <span>lorem</span>

                                </div>
                            </div>

                            <div>
                                <div>
                                    <div class="my-4"> <span class="font-bold">Point of Contact: </span>{{
                                        activeAccount.point_of_contact
                                    }}</div>
                                    <div class="my-4"> <span class="font-bold">Waterloo Id: </span>{{
                                        activeAccount.waterloo_id
                                    }}</div>
                                    <div class="my-4"> <span class="font-bold">Creator: </span>{{
                                        activeAccount.creator
                                    }}</div>
                                </div>
                            </div>
                            <div class="flex justify-center mt-6">
                                <v-btn color="primary" variant="outlined"
                                    @click="activeAccount ? switchTransaction(activeAccount.id) : null">
                                    Open Transactions For {{ activeAccount.name }}
                                </v-btn>
                            </div>
                        </v-card>
                        <v-card v-else class="flex-1 ml-2" color="background-light-1">

                        </v-card>
                    </div>
                </v-card>

            </div>
        </div>

        <v-dialog v-model="newTransactionModal" width="800">
            <NewTransactionModal :account-id="activeAccount?.id" @closeModal="newTransactionModal = false" />
        </v-dialog>

    </warg-page>


</template>

<script lang="ts">
import { defineComponent, ref } from 'vue';
import { useTheme } from 'vuetify';

import AccountModal from '../components/AccountModal.vue';
import { useRouter } from 'vue-router';
import NavBar from '../components/NavBar.vue';
import { useAccountStore } from '../store/accounts';
import { useProfileStore } from '../store/profile';
import { Account } from '../types';
import WargPage from '../components/WargPage.vue';
import axios from 'axios';
import NewTransactionModal from '../components/NewTransactionModal.vue';

export default defineComponent({
    components: { AccountModal, NavBar, WargPage, NewTransactionModal },
    setup() {
        const accountStore = useAccountStore();
        const profileStore = useProfileStore();


        const theme = useTheme();
        const dialog = ref(false);
        const toggleTheme = () => theme.global.name.value = theme.global.current.value.dark ? 'wargLight' : 'wargDark';

        const router = useRouter();
        const newTransactionModal = ref(false);

        const switchTransaction = (account_id: string) => {
            console.log(account_id);
            router.push({ name: 'Transaction', params: { account_id } });
        };

        const openNewTransaction = () => {
            console.log('open new transaction');
            newTransactionModal.value = true;
        };

        const filter = ref<string>('');


        // @ts-ignore
        const dev = import.meta.env.DEV;

        const activeAccount = ref<Account | undefined>();
        return { dialog, toggleTheme, theme, accountStore, switchTransaction, activeAccount, filter, profileStore, dev, newTransactionModal, openNewTransaction };
    }
});
</script>

<style scoped>

</style>