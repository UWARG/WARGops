import { createApp } from 'vue';
import App from './App.vue';

import router from './router';
import 'v-calendar/dist/style.css';

//@ts-ignore
import VCalendar from 'v-calendar';

// Vuetify
import 'vuetify/styles';
import { createVuetify } from 'vuetify';
import * as components from 'vuetify/components';
import * as directives from 'vuetify/directives';
import '@mdi/font/css/materialdesignicons.css';
import './style.css';
import { createPinia } from 'pinia';

import { VDataTable } from 'vuetify/labs/VDataTable'

const wargDark = {
  dark: true,
  colors: {
    background: "#2d2d2d",
    'background-dark-1': "#2a2a2a",
    'background-light-1': "#525252",
    primary: "#FFEB3B",
    discord: "#5865F2",
    "warg-blue": "#081b2e",

  },
};

const wargLight = {
  dark: false,
  colors: {
    background: "#f4f4f4",
    primary: "#FFEB3B",
  },
};

const vuetify = createVuetify({
  components:{
    VDataTable,
    ...components,
  },
    
  directives,
  theme: {
    defaultTheme: 'wargDark',
    themes: {
      wargLight: wargLight,
      wargDark: wargDark,
    }
  },

});

const pinia = createPinia();


const app = createApp(App);
app.use(vuetify);
app.use(router);
app.use(VCalendar, {});
app.use(pinia);
app.mount('#app');
