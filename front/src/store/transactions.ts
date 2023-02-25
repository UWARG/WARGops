import { useProfileStore } from "./profile";
import { useAccountStore } from "./accounts";
import { defineStore } from "pinia";
import { Transaction } from "@/types";
import axios from "axios";

export const useTransactionStore = defineStore("Transactions", {
  state: () => ({
    transactions: [] as Transaction[],
  }),
  getters: {

    // Based off the given statuses and types, returns the filtered transactions
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

    // Returns the transaction with the given id
    getTransactionById:
      (state) =>
        (id: string): Transaction | undefined => {
          return state.transactions.find((transaction) => transaction.id === id);
        },
  },
  actions: {

    // Loads the transactions for the account with the given id
    async loadTransactions(accountId: string) {
      this.transactions = [];
      const response = await axios.get(`/api/transactions/${accountId}`);
      const transactions = await response.data;
      if (transactions != null) {
        this.transactions = transactions;
      }
      else {
        this.transactions = [];
      }
    },

    // Updates the transaction status with the given id
    updateTransaction(
      account_id: string,
      transaction: Transaction,
      updateType: string
    ) {
      axios
        .post(`/api/transactions/${account_id}/${transaction.id}:${updateType}`, {
          notes: transaction.notes,
          name: transaction.name,
          id: transaction.id,
          approver: useProfileStore().profile.UserID,
        })
        .then((response) => {
          this.loadTransactions(account_id);
        });
    },

    // Creates a new transaction
    createTransaction(transaction: any,) {
      return axios
        .post("/api/transactions", transaction).then(
          ()=>{
            useAccountStore().getAccountInfo(transaction.account_id)
          });
    },
  },
});
