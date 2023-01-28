<template>
    <v-card>
        <v-toolbar color="">
            <v-toolbar-title color="primary" class="text-3xl font-bold"> Create New Transaction</v-toolbar-title>
            <v-btn :disabled="false" color="red" @click="closeModal">Cancel</v-btn>
            <v-btn :disabled="false" color="green" @click="submit">Submit</v-btn>
        </v-toolbar>
        <v-form class="p-4" v-model="valid">
            <div class="flex">
                <v-text-field class="mr-2" label="Amount" variant="outlined" v-model="newTransaction.amount"
                    :rules="[rules.required]" />
                <v-select class="ml-2" label="Name" variant="outlined" :rules="[rules.required]"
                    v-model="newTransaction.type" />
            </div>
        
        </v-form>
    </v-card>
</template>

<script lang="ts">
import { defineComponent, reactive, ref } from 'vue';
import { Transaction } from '../types';
import { rules } from '../helpers';
import axios from 'axios';
export default defineComponent({
    name: 'NewTransactionModal',
    props: {
        accountId: {
            type: String,
            required: true,
        },
    },
    setup(props, context) {
        const close = () => {
            context.emit('close');
        };

        const closeModal = () => {
            context.emit('closeModal');
        };

        const submit = () => {
            axios.post('http://localhost:8080/transactions', newTransaction).then((res) => {
                context.emit('closeModal');
            });
        };

        const valid = ref(false);

        const newTransaction = reactive({
            //@ts-ignore
            id: crypto.randomUUID(),
            amount: 0,
            type: 0,
            account_id: props.accountId,
            status: 0,
        });
        return { close, newTransaction, submit, valid, closeModal, rules };
    }
});
</script>

<style scoped>

</style>