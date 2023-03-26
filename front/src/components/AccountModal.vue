<template>
  <v-card>
    <v-toolbar color="">
      <v-toolbar-title color="primary" class="text-3xl font-bold">
        Create New Account</v-toolbar-title
      >
      <v-btn :disabled="false" color="red" @click="closeModal">Cancel</v-btn>
      <v-btn :disabled="!valid" color="green" @click="submit">Submit</v-btn>
    </v-toolbar>
    <v-form v-model="valid" :lazy-validation="false" ref="form" class="p-4">
      <div class="flex">
        <v-text-field
          class="mr-2 flex-1"
          label="Waterloo ID"
          variant="outlined"
          v-model="accountInfo.waterloo_id"
          :rules="[rules.required]"
        />
        <v-text-field
          class="ml-2 flex-1"
          label="Point of Contact"
          variant="outlined"
          v-model="accountInfo.point_of_contact"
          :rules="[rules.required, rules.email]"
        />
      </div>
      <div class="flex mb-4">
        <div class="flex-1 mr-2">
          <v-text-field
            label="Name"
            variant="outlined"
            :rules="[rules.required]"
            v-model="accountInfo.name"
          /> 
          <v-autocomplete
            label="Term"
            variant="outlined"
            :rules="[rules.required]"
            v-model="accountInfo.allocation_date"
            :items="dates"
            placeholder="Select..."
          />

          <v-text-field
            label="Source"
            class="w-full"
            variant="outlined"
            :rules="[rules.required]"
            v-model="accountInfo.source"
          />
          <v-switch
            color="primary"
            label="Active"
            class="w-36 ml-8"
            v-model="accountInfo.active"
          ></v-switch>
        </div>
        <div class="flex-1 ml-2">
          <span class="text-gray-400">Expiry Date</span>
          <v-date-picker
            class="inline-block flex-1 w-full"
            v-model="accountInfo.expiry_date"
            color="yellow"
            is-dark
          />
        </div>
      </div>
    </v-form>
  </v-card>
</template>

<script lang="ts">
import { defineComponent, reactive, ref } from "vue";
import { rules } from "../helpers";
import WargDatePicker from "./DatePicker.vue";
import { useProfileStore } from "../store/profile";
import { useAccountStore } from "../store/accounts";

export default defineComponent({
  name: "AccountModal",
  components: { WargDatePicker },
  setup(props, context) {
    const closeModal = () => {
      context.emit("closeModal");
    };

    const valid = ref(false);
    const accountInfo = reactive({
      waterloo_id: "",
      name: "",
      source: "",
      active: true,
      allocation_date: "",
      expiry_date: new Date().toISOString(),
      creator: useProfileStore().profile.UserID,
      point_of_contact: "",
      //@ts-ignore
      id: crypto.randomUUID(),
    });
    const submit = async () => {
      await useAccountStore().createAccount(accountInfo);
      closeModal();
    };

    const date = ref(new Date());

    const terms = ["Spring", "Fall", "Winter"];
    // const dates = () => {
    //   let arr = [];
    //   let curr = new Date();
    //   for (let year = 2020; year <= curr.getFullYear()+1; year++) {
    //     for (let i = 0; i < 3; i++){
    //       arr.push(terms[(Math.floor(curr.getMonth()/4)+i)%3] + " " + year); 
    //     } 
    //   }
    //   return arr;
    // }
   
    let dates = [];
    let curr = new Date();
    for (let year = 2020; year <= curr.getFullYear()+1; year++) {
      for (let i = 0; i < 3; i++){
        dates.push(terms[(Math.floor(curr.getMonth()/4)+i)%3] + " " + year); 
      } 
    }
  

    return {
      closeModal,
      submit,
      accountInfo,
      valid,
      rules,
      date,
      dates,
    };
  },
});
</script>

<style lang="scss" scoped>
.vc-container {
  background-color: #2a2a2a;
  border-color: #2a2a2a;
}
</style>
