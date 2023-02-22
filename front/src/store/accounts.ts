import { defineStore } from "pinia";
import { Account } from "@/types";
import axios from "axios";

export const useAccountStore = defineStore("Accounts", {
  state: () => ({
    accounts: [] as Account[],
    activeAccoutId: "",
  }),
  getters: {
    getAccounts(state) {
      return state.accounts;
    },
    getAccountById:
      (state) =>
        (id: string): Account => {
          return state.accounts.find((account) => account.id === id) as Account;
        },
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
    async loadAccounts() {
      const response = await fetch("/api/accounts");
      const accounts = await response.json();
      this.accounts = accounts;
    },
    async createAccount(account: any) {
      const response = await axios.post("/api/accounts", account);
      this.loadAccounts();
    },
  },
});
