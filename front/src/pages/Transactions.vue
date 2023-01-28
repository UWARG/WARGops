<template>
    <v-app class="p-4">
        <nav-bar>
            <template #button>
                <v-btn color="red" @click="backToHome"> Back To Accounts </v-btn>
            </template>
        </nav-bar>
        <v-main>
            <v-card class="p-4">
                <v-card color="background-light-1" class="p-4">


                    <!-- Header -->
                    <div class="flex justify-between">
                        <v-card-title class="text-3xl font-bold mb-2">
                            Transactions
                        </v-card-title>
                        <div class="flex justify-center items-center mr-4">
                            <v-dialog v-model="newTransactionModal" width="800">
                                <template v-slot:activator="{ props }">
                                    <v-btn v-bind="props" color="primary" class=""> Add Income</v-btn>
                                </template>

                                <NewTranscationModal @closeModal="newTransactionModal = false"
                                    :accountId="account_id" />
                            </v-dialog>
                            <v-btn color="background-dark-1" class=""> Add Expense </v-btn>
                        </div>
                    </div>
                    <!-- Search -->
                    <div class="flex my-4">
                        <v-select label="Status: All" class="mx-2"
                            :items="[{ title: 'Deposit', value: 0 }, { title: 'Rembursment', value: 1 }, { title: 'Procurement', value: 2 }]"
                            multiple variant="outlined">
                            <template v-slot:selection="{ item }">
                                <type-chip :type="item.value" />
                            </template>
                        </v-select>
                        <v-select label="Type: All" class="mx-2"
                            :items="[{ title: 'Created', value: 0 }, { title: 'Pending', value: 1 }, { title: 'Paid', value: 2 }, { title: 'Rejected', value: 3      }]"
                            multiple variant="outlined">
                            <template v-slot:selection="{ item }">
                                <status-chip :type="item.value" />
                            </template>
                        </v-select>
                        <v-select label="Acccount: All" class="mx-2">

                        </v-select>
                        <v-select label="From: " class="mx-2"></v-select>
                        <v-select label="To: " class="mx-2"></v-select>
                    </div>
                    <!-- Transactions -->
                    <div>
                        <v-table color="background-light-1">
                            <thead>
                                <tr>
                                    <th>Amount</th>
                                    <th>Type</th>
                                    <th>Status</th>
                                    <th>Approval Date</th>
                                    <th>Creation Date</th>
                                    <th>Rejected Date</th>
                                </tr>
                            </thead>
                            <tbody>
                                <tr v-for="(transaction, index) in transactionStore.transactions" :key="index">
                                    <th> $ {{ transaction.amount }}</th>
                                    <th> <type-chip :type="transaction.type" /></th>
                                    <th> <status-chip :type="transaction.status" /></th>
                                    <th> {{ transaction.approval_date.toLocaleDateString() }}</th>
                                    <th> {{ transaction.creation_date.toLocaleDateString() }}</th>
                                    <th> {{ transaction.rejected_date.toLocaleDateString() }}</th>
                                </tr>
                            </tbody>
                        </v-table>
                    </div>
                </v-card>
            </v-card>
        </v-main>
    </v-app>
</template>

<script lang="ts">
import { defineComponent, onMounted, ref } from 'vue';
import NavBar from '../components/NavBar.vue';
import { useRouter } from 'vue-router';
import { useTransactionStore } from '../store/transactions';
import NewTranscationModal from '../components/NewTransactionModal.vue';
import TypeChip from '../components/TypeChip.vue';
import StatusChip from '../components/StatusChip.vue';

export default defineComponent({
    components: { NavBar, NewTranscationModal, TypeChip, StatusChip },
    setup() {
        const router = useRouter();
        const backToHome = () => {
            router.push({ name: 'Home' });
        };

        const { account_id } = router.currentRoute.value.params;
        const transactionStore = useTransactionStore();


        const newTransactionModal = ref(false);
        return { backToHome, newTransactionModal, account_id: account_id as string, transactionStore };
    }
});
</script>

<style scoped>

</style>
