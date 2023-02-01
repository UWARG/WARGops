<template>
    <v-chip color="pink" class="ma-1" v-if="type == 0">
        Deposit
    </v-chip>
    <v-chip color="blue" class="ma-1" v-else-if="type == 1">
        Rembursment
    </v-chip>
    <v-chip color="purple" class="ma-1" v-else-if="type == 2">
        Procurement
    </v-chip>
    <v-chip color="grey" class="ma-1" v-else>
        Unknown
    </v-chip>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
    props: {
        type: {
            type: Number,
            required: true
        }
    },
    setup() {
        return {};
    }
});
</script>

<style scoped>

</style>