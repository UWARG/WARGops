<template>
  <div class="m-0 p-0">
    <router-view />
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted } from 'vue';
import { useAccountStore } from "./store/accounts";
import { useTransactionStore } from "./store/transactions";
import { useProfileStore } from "./store/profile";

export default defineComponent({
  setup() {
    const accountStore = useAccountStore();
    const transactionStore = useTransactionStore();
    const profileStore = useProfileStore();

    onMounted(() => {
      accountStore.loadAccounts();
    });

    return {};
  }
});
</script>

<style scoped>

</style>