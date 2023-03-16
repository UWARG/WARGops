<template>
  <router-view />
</template>

<script lang="ts">
import { defineComponent, onMounted } from "vue";
import { useProfileStore } from "./store/profile";
import router from "./router";

export default defineComponent({
  setup() {
    onMounted(async () => {
      if (!useProfileStore().getLoggedIn) {
        return;
      }
      try {
        await useProfileStore().loadProfile();
      } finally {
        console.log(useProfileStore().getLoggedIn);
        if (!useProfileStore().getLoggedIn) {
          router.push({ name: "Login" });
        }
      }
    });

    return {};
  },
});
</script>
