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
        .get("http://localhost:8080/api/user")
        .then((res) => {
          this.profile = res.data;
        })
        .catch((err) => {
          console.log(err);
        });
    },
    getGuilds() {
      const url = new URL(
        "http://localhost:8080/api/guilds/776618956638388305/roles"
      );
      axios
        .get(url.toString())
        .then((res) => {
          console.log(res.data);
        })
        .catch((err) => {
          console.log(err);
        });
    },
    async checkProfile() {
      console.log("You are logged in:", this.profile.username !== undefined);
      return this.profile.username !== undefined;
    },
    logout() {
      console.log("Logging out...");
      this.profile = {} as Profile;
      const url = new URL("http://localhost:8080/api/logout");
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
