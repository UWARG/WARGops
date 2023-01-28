import { defineStore } from 'pinia';
import { Transaction } from '@/types';

export const useTransactionStore = defineStore('Transactions', {
    state: () => ({
        transactions: [
            {
                id: '1',
                amount: 100,
                approval_date: new Date(),
                creation_date: new Date(),
                notes: 'test',
                payment_date: new Date(),
                rejected_date: new Date(),
                status: 0,
                type: 0,
            },
            {
                id: '2',
                amount: 200,
                approval_date: new Date(),
                creation_date: new Date(),
                notes: 'test',
                payment_date: new Date(),
                rejected_date: new Date(),
                status: 1,
                type: 1,
            },
        ] as Transaction[],
    }),
    getters: {

    },
    actions: {
        async loadTransactions(accountId: string) {
            const response = await fetch(`http://localhost:8080/transactions/${accountId}/all`);
            const transaction = await response.json();
            this.transactions = transaction;
        }
    }
});