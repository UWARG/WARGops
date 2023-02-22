<template>
  <div class="m-0 p-0">
    <router-view />
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  setup() {
    return {};
  },
});
</script>

<style lang="scss">
.v-application .primary--text {
  color: red !important;
}
</style>

