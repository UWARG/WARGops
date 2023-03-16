<template>
  <v-app-bar color="warg-blue">
    <div class="w-full flex flex-row items-center justify-between px-4">
      <div class="flex-1">
        <img
          src="../assets/warg-logo.svg"
          class="h-8 cursor-pointer"
          @click="toHome"
        />
      </div>
      <div class="flex items-center justify-center gap-8">
        <p @click="toWarg" class="cursor-pointer">WARG</p>
        |
        <router-link to="/">Dashboard</router-link>
        |
        <router-link to="/about">About</router-link>
        |
        <p @click="toJoin">Join</p>
      </div>
      <div class="flex items-center flex-1 justify-end">
        <v-menu open-on-hover :close-on-content-click="false">
          <template v-slot:activator="{ props }">
            <v-btn icon v-bind="props">
              <v-avatar class="mx-4">
                <v-img
                  alt="Avatar"
                  :src="profileStore.profile.AvatarURL"
                  id="menu-activator"
                />
              </v-avatar>
            </v-btn>
          </template>
          <v-card class="items-center py-2 px-4">
            <div class="flex">
              <div class="inline mr-4">
                <p>Logged in as</p>
                <p class="font-bold">{{ profileStore.profile.Name }}</p>
              </div>
              <v-btn color="red" variant="text" @click="logout">Logout</v-btn>
            </div>
            <div>
              Permissions:
              {{
                profileStore.getIsGuest
                  ? "Guest"
                  : profileStore.getIsLead
                  ? "Lead"
                  : "Member"
              }}
            </div>
          </v-card>
        </v-menu>
      </div>
    </div>
  </v-app-bar>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { useTheme } from "vuetify";
import { useProfileStore } from "../store/profile";
import { useRouter } from "vue-router";

export default defineComponent({
  name: "NavBar",
  props: {
    profile: {
      type: Boolean,
      default: true,
    },
  },
  setup() {
    const theme = useTheme();
    const toggleTheme = () =>
      (theme.global.name.value = theme.global.current.value.dark
        ? "wargLight"
        : "wargDark");
    const profileStore = useProfileStore();
    const router = useRouter();
    const logout = () => {
      profileStore.logout().then(() => {
        router.push({ name: "Login" });
      });
    };

    const toHome = () => {
      router.push({ name: "Home" });
    };

    const toWarg = () => {
      window.location.href = "https://www.uwarg.com/";
    };

    const toJoin = () => {
      window.location.href = "https://discord.com/invite/rqMEV3m3hh";
    };

    return { toggleTheme, theme, profileStore, logout, toWarg, toJoin, toHome };
  },
});
</script>

<style scoped></style>
