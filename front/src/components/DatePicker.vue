<template>
    <v-date-picker class="inline-block h-full" v-model="date" color="yellow" is-dark>
        <template v-slot="{ inputValue, togglePopover }">
            <div class="flex items-center">
                <v-text-field class="w-full" :value="inputValue" @click="togglePopover"
                    prepend-inner-icon="mdi-calendar-range" label="Expiry Date" variant="outlined" readonl0 />
            </div>
        </template>
    </v-date-picker>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue';

export default defineComponent({
    name: 'WargDatePicker',
    setup(props, context) {
        const showDatePicker = ref(false);
        const openDatePicker = () => {
            showDatePicker.value = !showDatePicker.value;
        };
        const date = ref(new Date());

        const dateChanged = (date: Date) => {
            context.emit('input', date)
        };

        return { openDatePicker, showDatePicker, date };
    }
});
</script>

<style scoped>
.vc-container {
    background-color: blue;
}
</style>