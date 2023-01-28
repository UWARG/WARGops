<template>
    <v-card>
        <v-toolbar color="">
            <v-toolbar-title color="primary" class="text-3xl font-bold"> Create New Account</v-toolbar-title>
            <v-btn :disabled="false" color="red" @click="closeModal">Cancel</v-btn>
            <v-btn :disabled="!valid" color="green" @click="submit">Submit</v-btn>
        </v-toolbar>
        <v-form v-model="valid" :lazy-validation="false" ref="form" class="p-4">
            <div class="flex">
                <v-text-field class="mr-2" label="Waterloo ID" variant="outlined" v-model="accountInfo.waterloo_id"
                    :rules="[rules.required]" />
                <v-text-field class="ml-2" label="Name" variant="outlined" :rules="[rules.required]"
                    v-model="accountInfo.name" />
            </div>
            <div class="flex">
                <v-text-field label="Allocation Date" class="w-full mr-2" variant="outlined" :rules="[rules.required]"
                    v-model="accountInfo.allocation_date" />
                <v-text-field label="Expiry Date" class="w-full ml-2" variant="outlined" :rules="[rules.required]"
                    v-model="accountInfo.expiry_date" />
                <!-- <warg-date-picker class="flex-1 mr-2"> </warg-date-picker>
                <warg-date-picker class="flex-1 ml-2"> </warg-date-picker> -->
            </div>
            <div class="flex">
                <v-text-field label="Source" class="w-full" variant="outlined" :rules="[rules.required]"
                    v-model="accountInfo.source" />
                <v-switch color="primary" label="Active" class="w-36 ml-8" v-model="accountInfo.active"></v-switch>
            </div>
            <v-text-field label="Creator" variant="outlined" v-model="accountInfo.creator" :rules="[rules.required,]" />
            <v-text-field label="Point of Contact" variant="outlined" v-model="accountInfo.point_of_contact"
                :rules="[rules.required, rules.email]" />


        </v-form>
    </v-card>
</template>
        
<script lang="ts">
import { defineComponent, reactive, ref } from 'vue';
import axios from 'axios';
import { rules } from "../helpers";
import WargDatePicker from './DatePicker.vue';


export default defineComponent({
    name: 'AccountModal',
    components: { WargDatePicker },
    setup(props, context) {
        const closeModal = () => {
            context.emit('closeModal');
        };

        const valid = ref(false);
        const accountInfo = reactive({
            waterloo_id: '',
            name: '',
            source: '',
            active: false,
            allocation_date: new Date().toISOString(),
            expiry_date: new Date().toISOString(),
            creator: '',
            point_of_contact: '',
            //@ts-ignore
            id: crypto.randomUUID()
        });
        const submit = () => {
            axios.post('http://localhost:8080/accounts', accountInfo);
        };
        const openDatePicker = () => {
            console.log('open date picker');
        };

        return {
            closeModal,
            submit,
            accountInfo,
            valid,
            rules

        };
    }
});

</script>

<style lang="scss" scoped>

</style>