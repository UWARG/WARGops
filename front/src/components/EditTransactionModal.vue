<template>
  <v-card>
    <v-toolbar color="">
      <v-toolbar-title color="primary" class="text-3xl font-bold"
        >Edit {{ transaction.name }} (
        <span class="text-green-700">${{ transaction.amount }}</span>
        )</v-toolbar-title
      >
      <v-btn :disabled="false" color="red" @click="closeModal">Close</v-btn>
    </v-toolbar>

    <div class="flex p-4">
      <v-card class="flex-1 mr-2 p-4" color="background-light-1">
        <v-table class="">
          <thead color="background-light-1">
            <tr>
              <th>Name</th>
              <th>Amount</th>
              <th>Type</th>
              <th>Status</th>
              <th>Approval Date</th>
              <th>Creation Date</th>
              <th>Rejected Date</th>
            </tr>
          </thead>
          <tbody v-if="transactionStore.transactions">
            <tr>
              <th>{{ transaction.name }}</th>
              <th>$ {{ transaction.amount }}</th>
              <th><type-chip :type="transaction.type" /></th>
              <th><status-chip :type="transaction.status" /></th>
              <th>
                {{ new Date(transaction.approval_date).toLocaleDateString() }}
              </th>
              <th>
                {{ new Date(transaction.creation_date).toLocaleDateString() }}
              </th>
              <th>
                {{ new Date(transaction.rejected_date).toLocaleDateString() }}
              </th>
            </tr>
          </tbody>
        </v-table>

        <div class="w-36 inline">
          <v-textarea
            label="Notes"
            class="mt-4"
            v-model="transaction.notes"
          ></v-textarea>
          <div class="flex justify-center" v-if="transaction.status == 0">
            <v-btn @click="approveTransaction" color="green" variant="tonal"
              >Approve</v-btn
            >
            <v-divider vertical class="mx-4"></v-divider>
            <v-btn @click="rejectTransaction" color="red" variant="tonal"
              >Decline</v-btn
            >
          </div>
          <div class="flex justify-center" v-else-if="transaction.status == 1">
            <v-btn @click="payTransaction" color="blue" variant="tonal"
              >Mark As Paid</v-btn
            >
            <v-divider vertical class="mx-4"></v-divider>
            <v-btn @click="holdTransaction" color="grey" variant="tonal"
              >Reset</v-btn
            >
          </div>
          <div class="flex justify-center" v-else-if="transaction.status == 2">
            <v-btn @click="holdTransaction" color="grey" variant="tonal"
              >Reset</v-btn
            >
          </div>
        </div>
      </v-card>
    </div>
  </v-card>
</template>

<script lang="ts">
import { defineComponent, ref, computed } from "vue";
import { useTransactionStore } from "../store/transactions";
import { useProfileStore } from "../store/profile";
import StatusChip from "./StatusChip.vue";
import TypeChip from "./TypeChip.vue";


export default defineComponent({
  name: "EditTransactionModal",
  components: {
    StatusChip,
    TypeChip,
  },
  props: {
    transactionId: {
      type: String,
      required: true,
    },
    accountId: {
      type: String,
      required: true,
    },
  },
  setup(props, context) {
    const transactionStore = useTransactionStore();
    const profileStore = useProfileStore();
    const transaction = computed(
      () => transactionStore.getTransactionById(props.transactionId)!
    );

    const closeModal = () => {
      context.emit("closeModal");
    };

    const submit = () => {
      console.log("Edit Transaction");
    };

    const approveTransaction = () => {
      transactionStore.updateTransaction(
        props.accountId,
        transaction.value,
        "approve"
      );
      closeModal();
    };

    const rejectTransaction = () => {
      transactionStore.updateTransaction(
        props.accountId,
        transaction.value,
        "reject"
      );
      closeModal();
    };

    const payTransaction = () => {
      transactionStore.updateTransaction(
        props.accountId,
        transaction.value,
        "pay"
      );
      closeModal();
    };

    const holdTransaction = () => {
      transactionStore.updateTransaction(
        props.accountId,
        transaction.value,
        "hold"
      );
      closeModal();
    };

    const valid = ref(true);
    return {
      closeModal,
      valid,
      transactionStore,
      submit,
      transaction,
      profileStore,
      approveTransaction,
      rejectTransaction,
      payTransaction,
      holdTransaction,
    };
  },
});
</script>

<style scoped></style>

