<template>
    <v-date-picker class="inline-block h-full" v-model="date" color="yellow" is-dark>
        <template v-slot="{ inputValue, togglePopover }">
            <div class="flex items-center">
                <v-text-field class="w-full" :value="inputValue" @click="togglePopover" @change="dateChanged"
                    prepend-inner-icon="mdi-calendar-range" label="Expiry Date" variant="outlined" readonly />
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

        const dateChanged = () => {
            context.emit('input', date.value)
        };

        return { openDatePicker, showDatePicker, date, dateChanged };
    }
});
</script>

<style scoped>
.vc-container {
    background-color: blue;
}
</style>