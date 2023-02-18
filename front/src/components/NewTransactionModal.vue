<template>
    <v-card>
        <v-toolbar color="">
            <v-toolbar-title color="primary" class="text-3xl font-bold"> Create New Transaction</v-toolbar-title>
            <v-btn :disabled="false" color="red" @click="closeModal">Cancel</v-btn>
            <v-btn :disabled="!valid" color="green" @click="submit">Submit</v-btn>
        </v-toolbar>
        <v-form class="p-4" v-model="valid">
            <div class="flex">
                <v-text-field class="mr-2 flex-1" label="Name" variant="outlined" v-model="newTransaction.name"
                    :rules="[rules.required]"></v-text-field>
                <v-text-field class="ml-2 flex-1" label="Amount" variant="outlined" v-model="newTransaction.amount"
                    :rules="[rules.required, rules.money]"> <span class="mr-2">$</span> </v-text-field>
            </div>

            <div class="flex">
                <v-select label="Type" class="mr-2 flex-1" v-model="newTransaction.type"
                    :items="[{ title: 'Deposit', value: 0 }, { title: 'Rembursment', value: 1 }, { title: 'Procurement', value: 2 }]"
                    variant="outlined">
                    <template v-slot:selection="{ item }">
                        <type-chip :type="item.value" />
                    </template>
                </v-select>

                <v-select label="Status" class="ml-2 flex-1" v-model="newTransaction.status"
                    :items="[{ title: 'Created', value: 0 }, { title: 'Pending', value: 1 }, { title: 'Paid', value: 2 }, { title: 'Rejected', value: 3 }]"
                    variant="outlined">
                    <template v-slot:selection="{ item }">
                        <status-chip :type="item.value" />
                    </template>
                </v-select>

                
            </div>
            <v-textarea label="Notes" variant="outlined" v-model="newTransaction.notes" auto-grow>
                    
            </v-textarea>
        </v-form>
    </v-card>
</template>

<script lang="ts">
import { defineComponent, reactive, ref } from 'vue';
import { Transaction, NewTransaction } from '../types';
import { rules } from '../helpers';
import axios from 'axios';
import StatusChip from './StatusChip.vue';
import TypeChip from './TypeChip.vue';
import { useTransactionStore } from '../store/transactions';

export default defineComponent({
    name: 'NewTransactionModal',
    props: {
        accountId: {
            type: String,
            required: true,
        },
    },
    components: {
        StatusChip,
        TypeChip,
    },
    setup(props, context) {
        const close = () => {
            context.emit('close');
        };

        const closeModal = () => {
            context.emit('closeModal');
        };

        const submit = () => {
            newTransaction.amount = Number(newTransaction.amount);
            axios.post('http://localhost:8080/transactions', newTransaction).then((res) => {
                context.emit('closeModal');
                useTransactionStore().loadTransactions(props.accountId);
            });
        };

        const valid = ref(false);

        const newTransaction = reactive<NewTransaction>({
            //@ts-ignore
            id: crypto.randomUUID(),
            name: '',
            amount: 0,
            notes: '',
            type: 0,
            account_id: props.accountId,
            status: 0,
        });
        return { close, newTransaction, submit, valid, closeModal, rules };
    }
});
</script>