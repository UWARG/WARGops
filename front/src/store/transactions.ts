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
        },

        getTransactionById: (state) => (id: string): Transaction | undefined => {
            return state.transactions.find((transaction) => transaction.id === id);
        }
    },
    actions: {
        async loadTransactions(accountId: string) {
            const response = await axios.get(`http://localhost:8080/transactions/${accountId}`);
            const transactions = await response.data;
            this.transactions = transactions;
        },
        rejectTransaction(account_id: string, transactionId: string) {
            axios.post(`http://localhost:8080/transactions/${account_id}/${transactionId}:reject`).then(
                (response) => {
                    this.loadTransactions(account_id);
                }
            );
        },
        holdTransaction(account_id: string, transactionId: string) {
            axios.post(`http://localhost:8080/transactions/${account_id}/${transactionId}:hold`).then(
                (response) => {
                    this.loadTransactions(account_id);
                }
            );
        },
        payTransaction(account_id: string, transactionId: string) {
            axios.post(`http://localhost:8080/transactions/${account_id}/${transactionId}:pay`).then(
                (response) => {
                    this.loadTransactions(account_id);
                }
            );
        }

    }
});