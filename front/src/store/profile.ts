import { defineStore } from "pinia";
import { Profile } from "../types";
import axios from "axios";
import { useRouter } from 'vue-router';


export const useProfileStore = defineStore("Profile", {
  state: () => ({
    profile: {} as Profile,
  }),
  getters: {
    // Returns true if the user is logged in
    getLoggedIn(state): boolean {
      console.log(state.profile.Name);
      return state.profile.Name !== undefined && state.profile.Name !== "";
    },
    // Getter for if user is a lead
    getIsLead(state): boolean {
      if (this.getLoggedIn) {
        return state.profile.RawData.roles.find((role: string) => role === import.meta.env.VITE_LEAD_ROLE) !== undefined;
      }
      return false;
    }
  },
  actions: {
    // Loads the profile of the logged in user based on the cookie
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
    // Wipes the profile cookie and logs out the user
    async logout() {
      console.log("Logging out...");
      this.profile = {} as Profile;
      document.cookie = "auth" + '=; Path=/; Expires=Thu, 01 Jan 1970 00:00:01 GMT;';
      return axios
        .get("/api/logout")
        .then((res) => {
          console.log(res.data);
        })
        .catch((err) => {
          console.log(err);
        });
    },

    async signInAsGuest(): Promise<void> {
      console.log("Signing in as guest...");
      return axios
        .get("/api/guest")
        .then((res) => {
          this.profile = res.data;
        })
        .catch((err) => {
        });
    }
  },
});
