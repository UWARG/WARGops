<template>
  <div>
    <router-view />
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted } from 'vue';
import { useAccountStore } from "./store/accounts";
import { useTransactionStore } from "./store/transactions";

export default defineComponent({
  setup() {
    const accountStore = useAccountStore();
    const transactionStore = useTransactionStore();

    onMounted(() => {
      accountStore.loadAccounts();
    });

    return {};
  }
});
</script>

<style scoped>

</style>