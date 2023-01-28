<template>
  <div>
    <router-view />
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted } from 'vue';
import { useAccountStore } from "./store/accounts";

export default defineComponent({
  setup() {
    const accountStore = useAccountStore();
    onMounted(() => {
      accountStore.loadAccounts();
    });

    return {};
  }
});
</script>

<style scoped>

</style>