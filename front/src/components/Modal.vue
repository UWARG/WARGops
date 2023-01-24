<template>
    <v-card>
        <v-toolbar color="">
            <v-toolbar-title color="primary" class="text-3xl font-bold"> Create New Account</v-toolbar-title>
            <v-btn :disabled="true">Submit</v-btn>
        </v-toolbar>
        <v-card-text>
            <v-form>
                <div class="flex">
                    <v-text-field class="mr-2" label="Waterloo ID" required></v-text-field>
                    <v-text-field class="ml-2" label="Name" required></v-text-field>
                </div>

                <v-text-field label="Source" required></v-text-field>
                <v-text-field label="Allocation Date" required></v-text-field>
                <v-text-field label="Expiry Date" required></v-text-field>
                <v-switch color="primary" label="Active" required></v-switch>
                <v-text-field label="Creator" required></v-text-field>
                <v-text-field label="Point of Contact" required></v-text-field>
            </v-form>
        </v-card-text>
    </v-card>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
    setup() {


        return {};
    }
});
</script>

<style scoped>

</style>