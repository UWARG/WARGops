import { defineStore } from "pinia";
import { Account } from "@/types";
import axios from "axios";

export const useAccountStore = defineStore("Accounts", {
  state: () => ({
    accounts: [] as Account[],
    activeAccoutId: "",
  }),
  getters: {
    // Returns the account with the given id
    getAccountById:
      (state) =>
      (id: string): Account => {
        return state.accounts.find((account) => account.id === id) as Account;
      },
    // Based on search filter, returns the filtered accounts
    getFilteredAccounts:
      (state) =>
      (filter: string): Account[] => {
        return filter
          ? state.accounts.filter((account) =>
              account.name.toLowerCase().includes(filter.toLowerCase())
            )
          : state.accounts;
      },
  },
  actions: {
    // Loads all the accounts
    async loadAccounts() {
      const response = await fetch("/api/accounts");
      const accounts = await response.json();
      this.accounts = accounts;
    },
    // Creates a new account
    async createAccount(account: any) {
      const response = await axios.post("/api/accounts", account);
      this.loadAccounts();
    },

    //Gets more information about the account
    async getAccountInfo(accountId: string) {
      const response = await axios.get(`/api/accounts/${accountId}`);
      const account = await response.data;
      this.accounts = this.accounts.map((acc) => {
        if (acc.id === accountId) {
          account.external = true;
          return account;
        }
        return acc;
      });
    },
  },
});
