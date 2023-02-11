<template>
    <v-chip :color="color[type]" :size="size" class="ma-1">
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
        },
        size: {
            type: String,
            required: false,
            default: 'default'
        }
    },
    setup() {
        const color = [
            'green',
            'orange',
            'blue',
            'red',
            'grey'
        ];
        const text = [
            'Created',
            'Pending',
            'Paid',
            'Rejected',
            'Unknown'
        ];

        return { color, text };
    }
});
</script>

<style scoped>

</style>