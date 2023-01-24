import { createApp } from 'vue';
import App from './App.vue';

// Vuetify
import 'vuetify/styles';
import { createVuetify } from 'vuetify';
import * as components from 'vuetify/components';
import * as directives from 'vuetify/directives';

import '@mdi/font/css/materialdesignicons.css';

import './style.css';

const wargDark = {
  dark: true,
  colors: {
    background: "#3a3a3a",
    'background-dark-1': "#2a2a2a",
    'background-light-1': "#525252",
    primary: "#FFEB3B",

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
  components,
  directives,
  theme: {
    defaultTheme: 'wargDark',
    themes: {
      wargLight: wargLight,
      wargDark: wargDark,
    }
  },

});

createApp(App).use(vuetify).mount('#app');
