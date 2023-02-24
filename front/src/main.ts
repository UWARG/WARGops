import { createApp } from "vue";
import App from "./App.vue";

import router from "./router";
import "v-calendar/dist/style.css";

//@ts-ignore
import VCalendar from "v-calendar";

// Vuetify
import "vuetify/styles";
import { createVuetify } from "vuetify";
import * as components from "vuetify/components";
import * as directives from "vuetify/directives";
import "@mdi/font/css/materialdesignicons.css";
import "./style.css";
import "./scss/main.scss"
import { createPinia } from "pinia";

import { VDataTable } from "vuetify/labs/VDataTable";

const wargDark = {
  dark: true,
  colors: {
    primary: "#FFEB3B",
    background: "#2d2d2d",
    "background-dark-1": "#2a2a2a",
    "background-light-1": "#525252",
    discord: "#5865F2",
    "warg-blue": "#081b2e",
    "warg-blue-light": "#9ccaf8",
    "warg-blue-light2": "#5aacfe",
    "warg-accent": "#FCE2D3",

  },
  
};
const vuetify = createVuetify({
  components: {
    VDataTable,
    ...components,
  },

  directives,
  theme: {
    defaultTheme: "wargDark",
  },
});

const pinia = createPinia();

const app = createApp(App);
app.use(vuetify);
app.use(router);
app.use(VCalendar, {});
app.use(pinia);
app.mount("#app");
