import { defineStore } from "pinia";
import { Profile } from "@/types";
import axios from "axios";

export const useProfileStore = defineStore("Profile👨", {
  state: () => ({
    profile: {} as Profile,
    alert: { alert: false },
  }),
  getters: {
    getProfile(state) {
      return state.profile;
    },
    getProfileDefined(state): boolean {
      return state.profile !== undefined;
    },
    getAlert(state): any {
      return state.alert;
    },
    getLoggedIn(state): boolean {
      console.log(state.profile.Name);
      return state.profile.Name !== undefined && state.profile.Name !== "";
    },
  },
  actions: {
    toggleAlert() {
      this.alert.alert = !this.alert.alert;
      setTimeout(() => {
        this.alert.alert = !this.alert.alert;
      }, 3000);
    },

    setCode(code: string) {
      this.profile.code = code;
    },
    setProfile(profile: Profile) {
      this.profile = profile;
    },
    loadProfile(): Promise<void> {
      console.log("Loading profile...");
      return axios
        .get("/api/info")
        .then((res) => {
          this.profile = res.data;
        })
        .catch((err) => {
          console.log(err);
        });
    },
    getGuilds() {
      const url = new URL("/api/guilds/776618956638388305/roles");
      axios
        .get(url.toString())
        .then((res) => {
          console.log(res.data);
        })
        .catch((err) => {
          console.log(err);
        });
    },
    logout() {
      console.log("Logging out...");
      this.profile = {} as Profile;
      const url = new URL("/api/logout");
      axios
        .get(url.toString())
        .then((res) => {
          console.log(res.data);
        })
        .catch((err) => {
          console.log(err);
        });
    },
  },
});
