import { useProfileStore } from "./profile";
import { defineStore } from "pinia";
import { Transaction } from "@/types";
import axios from "axios";

export const useTransactionStore = defineStore("Transactions", {
  state: () => ({
    transactions: [] as Transaction[],
  }),
  getters: {
    getTransactions(state): Transaction[] {
      return state.transactions;
    },

    getFilteredTransactions:
      (state) =>
      (statuses: number[], types: number[]): Transaction[] => {
        if (statuses.length === 0 && types.length === 0) {
          return state.transactions;
        } else if (statuses.length === 0) {
          return state.transactions.filter((transaction) =>
            types.includes(transaction.type)
          );
        } else if (types.length === 0) {
          return state.transactions.filter((transaction) =>
            statuses.includes(transaction.status)
          );
        }
        return state.transactions.filter(
          (transaction) =>
            statuses.includes(transaction.status) &&
            types.includes(transaction.type)
        );
      },

    getTransactionById:
      (state) =>
      (id: string): Transaction | undefined => {
        return state.transactions.find((transaction) => transaction.id === id);
      },
  },
  actions: {

    async loadTransactions(accountId: string) {
      this.transactions = [];
      const response = await axios.get(`/api/transactions/${accountId}`);
      const transactions = await response.data;
      if (transactions != null) {
        this.transactions = transactions;
      }
      else{
        this.transactions = [];
      }
    },

    updateTransaction(
      account_id: string,
      transaction: Transaction,
      updateType: string
    ) {
      axios
        .post(`/transactions/${account_id}/${transaction.id}:${updateType}`, {
          notes: transaction.notes,
          name: transaction.name,
          id: transaction.id,
          approver: useProfileStore().profile.id,
        })
        .then((response) => {
          this.loadTransactions(account_id);
        });
    },
  },
});
