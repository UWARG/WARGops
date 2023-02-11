<template>
    <v-chip :color="color[type]" class="ma-1">
        {{ text[type] }}
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
        const text = [
            'Deposit',
            'Rembursment',
            'Procurement',
            'Unknown'
        ];

        const color = [
            'pink',
            'blue',
            'purple',
            'grey'
        ];

        return { text, color };
    }
});
</script>

<style scoped>

</style>