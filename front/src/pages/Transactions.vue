<template>
  <warg-page shownav>
    <div class="flex justify-center items-start mt-16 w-full h-screen">
      <div class="w-[70vw]">
        <!-- Header -->
        <div class="flex justify-between">
          <v-card-title class="text-3xl font-bold mb-2">
            Transactions for {{ accountStore.getAccountById(account_id).name }}
          </v-card-title>
          <div class="flex justify-center items-center mr-4">
            <v-dialog v-model="newTransactionModal" width="800">
              <template v-slot:activator="{ props }">
                <v-btn v-bind="props" color="primary" variant="tonal">New Transaction</v-btn>
              </template>

              <NewTranscationModal @closeModal="newTransactionModal = false" :accountId="account_id" />
            </v-dialog>
          </div>
        </div>
        <!-- Search -->
        <div class="flex my-4 items-start">
          <!-- <v-select label="Type: All" class="mx-2" v-model="typeFilter"
                                            :items="[{ title: 'Deposit', value: 0 }, { title: 'Rembursment', value: 1 }, { title: 'Procurement', value: 2 }]"
                                            multiple variant="outlined">
                                            <template v-slot:selection="{ item }">
                                                <type-chip :type="item.value" />
                                            </template>
                                        </v-select>
                                        <v-select label="Status: All" class="mx-2" v-model="statusFilter"
                                            :items="[{ title: 'Created', value: 0 }, { title: 'Pending', value: 1 }, { title: 'Paid', value: 2 }, { title: 'Rejected', value: 3 }]"
                                            multiple variant="outlined">
                                            <template v-slot:selection="{ item }">
                                                <status-chip :type="item.value" />
                                            </template>
                                        </v-select> -->
        </div>
        <!-- Transactions -->
        <v-data-table :items="
          transactionStore.getFilteredTransactions(statusFilter, typeFilter)
        " :headers="headers">
          <template #item.type="{ item }">
            <type-chip :type="item.props.title.type" />
          </template>
          <template #item.status="{ item }">
            <status-chip :type="item.props.title.status" />
          </template>
          <template #item.approval_date="{ item }">
            {{ new Date(item.props.title.approval_date).toLocaleDateString() }}
          </template>
          <template #item.creation_date="{ item }">
            {{ new Date(item.props.title.creation_date).toLocaleDateString() }}
          </template>

          <template #item.rejected_date="{ item }">
            {{ new Date(item.props.title.rejected_date).toLocaleDateString() }}
          </template>
          <template #item.actions="{ item }">
            <v-btn icon="mdi-pencil" color="grey" variant="tonal" size="x-small" @click="
              () => {
                activeId = item.props.title.id;
                dialog = true;
              }
            "></v-btn>
          </template>
        </v-data-table>
      </div>
    </div>
    <v-dialog v-model="dialog" width="800">
      <edit-transaction-modal :accountId="account_id" :transactionId="activeId" @closeModal="dialog = false" />
    </v-dialog>
  </warg-page>
</template>

<script lang="ts">
import { defineComponent, onBeforeMount, ref } from "vue";
import NavBar from "../components/NavBar.vue";
import { useRouter } from "vue-router";
import { useTransactionStore } from "../store/transactions";
import { useAccountStore } from "../store/accounts";
import { useProfileStore } from "../store/profile";

import NewTranscationModal from "../components/NewTransactionModal.vue";
import TypeChip from "../components/TypeChip.vue";
import StatusChip from "../components/StatusChip.vue";
import EditTransactionModal from "../components/EditTransactionModal.vue";
import WargPage from "../components/WargPage.vue";
import { onMounted } from "vue";

export default defineComponent({
  components: {
    NavBar,
    NewTranscationModal,
    TypeChip,
    StatusChip,
    EditTransactionModal,
    WargPage,
  },
  setup() {
    const router = useRouter();
    const backToHome = () => {
      router.push({ name: "Home" });
    };

    const { account_id } = router.currentRoute.value.params;
    const transactionStore = useTransactionStore();
    const accountStore = useAccountStore();
    const profileStore = useProfileStore();

    onBeforeMount(() => {
      transactionStore.loadTransactions(account_id as string);
    });

    const statusFilter = ref([]);
    const typeFilter = ref([]);

    const newTransactionModal = ref(false);

    const activeId = ref("");
    const amountSort = ref(0);

    const dialog = ref(false);

    const headers = [
      { title: "Name", align: "start", sortable: false, key: "name" },
      { title: "Amount ($)", align: "center", key: "amount" },
      { title: "Type", align: "center", key: "type" },
      { title: "Status", align: "center", key: "status" },
      { title: "Approval Date", align: "center", key: "approval_date" },
      { title: "Creation Date", align: "center", key: "creation_date" },
      { title: "Rejected Date", align: "center", key: "rejected_date" },
    ];

    onMounted(() => {
      if (profileStore.getIsLead) {
        headers.push({ title: "Actions", align: "center", key: "actions" });
      }
    });

    return {
      activeId,
      backToHome,
      newTransactionModal,
      account_id: account_id as string,
      transactionStore,
      statusFilter,
      typeFilter,
      dialog,
      accountStore,
      amountSort,
      headers,
      profileStore,
    };
  },
});
</script>

