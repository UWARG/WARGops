import { defineStore } from 'pinia';
import { Transaction } from '@/types';
import axios from 'axios';

export const useTransactionStore = defineStore('Transactions', {
    state: () => ({
        transactions: [] as Transaction[],
    }),
    getters: {
        getTransactions(state): Transaction[] {
            return state.transactions;
        },

        getFilteredTransactions: (state) => (statuses: number[], types: number[]): Transaction[] => {
            if (statuses.length === 0 && types.length === 0) {
                return state.transactions;
            }
            else if (statuses.length === 0) {
                return state.transactions.filter((transaction) => types.includes(transaction.type));
            }
            else if (types.length === 0) {
                return state.transactions.filter((transaction) => statuses.includes(transaction.status));
            }
            return state.transactions.filter((transaction) => statuses.includes(transaction.status) && types.includes(transaction.type));
        }
    },
    actions: {
        async loadTransactions(accountId: string) {
            const response = await axios.get(`http://localhost:8080/transactions/${accountId}`);
            const transactions = await response.data;
            this.transactions = transactions;
        }
    }
});