<template>
    <v-chip color="green" class="ma-1" v-if="type == 0">
        Created
    </v-chip>
    <v-chip color="yellow" class="ma-1" v-else-if="type == 1">
        Pending
    </v-chip>
    <v-chip color="orange" class="ma-1" v-else-if="type == 2">
        Paid
    </v-chip>
    <v-chip color="red" class="ma-1" v-else-if="type == 3">
        Rejected
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