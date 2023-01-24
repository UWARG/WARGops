<template>
  <v-app class="p-4">
    <v-app-bar class="flex justify-between">
      <div class="text-3xl font-bold ml-8">
        WARG ✈️
      </div>
      <v-btn :icon="theme.global.name.value === 'wargDark' ? 'mdi-weather-sunny' : 'mdi-weather-night'"
        @click="toggleTheme" />
    </v-app-bar>

    <v-main>

      <div class="flex">
        <div class="flex-1">
          <v-card class="p-4 mr-4">
            <div class="flex justify-between items-center mb-4">
              <v-card-title class="text-3xl font-bold">
                Accounts
              </v-card-title>
              <v-dialog v-model="dialog" width="800">
                <template v-slot:activator="{ props }">
                  <v-btn v-bind="props" color="primary" class=""> Create New Account</v-btn>
                </template>
                <Modal @closeModal="dialog = false" />
              </v-dialog>
            </div>

            <v-expansion-panels color="background-dark-1">
              <v-expansion-panel title="WEEF" text="This is a cool account">
              </v-expansion-panel>
              <v-expansion-panel title="WARG" text="This is a cool account">
              </v-expansion-panel>
              <v-expansion-panel title="MEF" text="This is a cool account">
              </v-expansion-panel>
            </v-expansion-panels>

          </v-card>

        </div>
        <v-card class="flex-1 ml-4 p-4">
          <v-card-title class="text-3xl font-bold mb-4">
            Transactions for WEEF
          </v-card-title>
          <v-card color="background-light-1 ">
            Test
          </v-card>
        </v-card>

      </div>



    </v-main>
  </v-app>

</template>

<script lang="ts">
import { defineComponent, ref } from 'vue';
import { useTheme } from 'vuetify';

import axios from 'axios';
import Modal from './components/Modal.vue';

export default defineComponent({
  components: { Modal },
  setup() {
    const theme = useTheme();
    const accountInfo = ref({});
    const getAccounts = () => {
      axios.get('http://localhost:8080/accounts')
        .then((response) => {
          accountInfo.value = response.data;
        })
        .catch((error) => {
          console.log(error);
        });
    };

    const testAccount = {
      //@ts-ignore
      Id: crypto.randomUUID(),
      WaterlooId: 123456789,
      Name: 'test',
      Source: 'test',
      AllocationDate: new Date(),
      ExpiryDate: new Date(),
      Active: true,
      Creator: 'test',
      PointOfContact: 'test',
    };

    const createAccountResponse = ref({});

    const createAccount = () => {
      axios.post('http://localhost:8080/accounts', testAccount)
        .then((response) => {
          createAccountResponse.value = response.data;
        })
        .catch((error) => {
          console.log(error);
        });
    };

    const dialog = ref(false);

    const toggleTheme = () => theme.global.name.value = theme.global.current.value.dark ? 'wargLight' : 'wargDark';



    return { accountInfo, getAccounts, createAccountResponse, createAccount, dialog, toggleTheme, theme };
  }
});
</script>

<style scoped>

</style>