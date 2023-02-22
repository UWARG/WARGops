<template>
  <warg-page shownav>
    <div class="flex justify-center items-start mt-16 w-full h-screen">
      <div class="w-[70vw]">
        <div class="flex">
          <h3 class="text-4xl font-bold mr-16">Funding Accounts</h3>
          <v-text-field density="compact" variant="solo" label="Search Accounts" class="w-80"
            append-inner-icon="mdi-magnify" single-line hide-details v-model="filter" />
        </div>
        <v-dialog v-model="dialog" width="800">
          <template v-slot:activator="{ props }">
            <v-btn v-bind="props" color="primary" size="large" variant="tonal" class="my-4" v-if="profileStore.getIsLead">
              Create New Account
            </v-btn>
          </template>
          <account-modal @closeModal="dialog = false" />
        </v-dialog>
        <v-expansion-panels class="my-4">
          <v-expansion-panel v-for="(account, index) in accountStore.getFilteredAccounts(filter)" :key="index" class="">
            <v-expansion-panel-title>
              <div class="flex w-full items-center">
                {{ account.name }}
                <v-chip variant="tonal" color="green" class="ml-auto mr-4">$ 676</v-chip>
              </div>
            </v-expansion-panel-title>
            <v-expansion-panel-text>
              <div class="flex py-4">
                <div class="mr-4 flex-1 flex">
                  <div class="flex flex-col mr-2">
                    <p>
                      <span class="font-bold">Point of Contact:</span>
                      {{ account.point_of_contact }}
                    </p>
                    <p>
                      <span class="font-bold">Account Id:</span>
                      {{ account.id }}
                    </p>
                    <p>
                      <span class="font-bold">Account Creator:</span>
                      {{ account.creator }}
                    </p>
                    <p>
                      <span class="font-bold">Created On: </span>{{ account.allocation_date }}
                    </p>
                  </div>
                  <div class="flex flex-col ml-2">
                    <p>
                      <span class="font-bold">Allocated Funds: </span> Lorem
                    </p>
                    <p><span class="font-bold">Account Balance:</span> Lorem</p>
                    <p><span class="font-bold">Pending Balance:</span> Lorem</p>
                    <p>
                      <span class="font-bold">Total Transactions: </span>Lorem
                    </p>
                  </div>
                </div>
                <div class="flex flex-col justify-center ml-4">
                  <v-btn @click="openNewTransaction" class="my-2" color="warg-blue-light" variant="tonal">
                    New Transaction
                    <v-icon icon="mdi-plus" class="ml-1" />
                  </v-btn>
                  <v-btn class="my-2" color="warg-blue-light2" variant="tonal" @click="switchTransaction(account.id)">
                    View All Transactions
                    <v-icon icon="mdi-book-open-variant" class="ml-1" />
                  </v-btn>
                </div>
              </div>
            </v-expansion-panel-text>
          </v-expansion-panel>
        </v-expansion-panels>
      </div>
    </div>
    <v-dialog v-model="newTransactionModal" width="800">
      <NewTransactionModal :account-id="activeAccount.id" @closeModal="newTransactionModal = false" />
    </v-dialog>
  </warg-page>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted } from "vue";
import NavBar from "../components/NavBar.vue";
import WargPage from "../components/WargPage.vue";
import { useAccountStore } from "../store/accounts";
import { useProfileStore } from "../store/profile";
import AccountModal from "../components/AccountModal.vue";
import NewTransactionModal from "../components/NewTransactionModal.vue";
import { useRouter } from "vue-router";
import { Account } from "../types";
import axios from "axios";

export default defineComponent({
  name: "NewHome",
  components: {
    NavBar,
    WargPage,
    AccountModal,
    NewTransactionModal,
  },
  setup() {
    const accountStore = useAccountStore();
    const profileStore = useProfileStore();
    const dialog = ref(false);
    const filter = ref();
    const newTransactionModal = ref(false);
    const router = useRouter();

    const switchTransaction = (account_id: string) => {
      router.push({ name: "Transaction", params: { account_id } });
    };

    const openNewTransaction = (account_id: string) => {
      accountStore.getAccountById(account_id);
      newTransactionModal.value = true;
    };
    const activeAccount = ref<Account>({} as Account);

    onMounted(() => {
      accountStore.loadAccounts();
    });

    return {
      accountStore,
      dialog,
      filter,
      switchTransaction,
      openNewTransaction,
      newTransactionModal,
      activeAccount,
      profileStore,
    };
  },
});
</script>

<style scoped></style>
