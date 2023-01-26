<template>
    <v-card>
        <v-toolbar color="">
            <v-toolbar-title color="primary" class="text-3xl font-bold"> Create New Account</v-toolbar-title>
            <v-btn :disabled="false" color="red" @click="closeModal">Cancel</v-btn>
            <v-btn :disabled="false" color="green" @click="submit">Submit</v-btn>
        </v-toolbar>
        <v-card-text>
            <v-form ref="myForm">
                <div class="flex">
                    <v-text-field class="mr-2" label="Waterloo ID" required variant="outlined"></v-text-field>
                    <v-text-field class="ml-2" label="Name" required variant="outlined"></v-text-field>
                </div>

                <div class="flex">
                    <v-text-field label="Source" required class="w-full" variant="outlined"></v-text-field>
                    <v-switch color="primary" label="Active" required class="w-36 ml-8"></v-switch>
                </div>

                <div class="flex">
                    <v-text-field prepend-inner-icon="mdi-calendar-range" @click:prepend-inner="openDatePicker"
                        label="Allocation Date" required class="mr-2" variant="outlined"></v-text-field>
                    <v-text-field prepend-inner-icon="mdi-calendar-range" @click:prepend-inner="openDatePicker"
                        label="Expiry Date" required class="ml-2" variant="outlined"></v-text-field>
                </div>


                <v-text-field label="Creator" required variant="outlined"></v-text-field>
                <v-text-field label="Point of Contact" required variant="outlined"></v-text-field>
            </v-form>
        </v-card-text>

    </v-card>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue';
import axios from 'axios';


export default defineComponent({
    setup(props, context) {
        const closeModal = () => {
            context.emit('closeModal');
        };
        const myForm = ref();
        const testForm = {
            Name: 'test',
            Id: 'test',
            Source: 'test',
            Active: true,
            AllocationDate: new Date(),
            ExpiryDate: new Date(),
            Creator: 'test',
            PointOfContact: 'ryegg360@gmail.com',
            WaterlooID: 'reggens'

        };
        const submit = () => {
            axios.post('http://localhost:8080/accounts', testForm);
        };
        const openDatePicker = () => {
            console.log('open date picker');
        };

        return {
            closeModal,
            submit,
            openDatePicker

        };
    }
});

</script>

<style lang="scss" scoped>

</style>