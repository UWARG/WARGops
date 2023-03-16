import { defineStore } from "pinia";
import { Profile } from "../types";
import axios from "axios";

export const useProfileStore = defineStore("Profile", {
  state: () => ({
    profile: {} as Profile,
    external: false,
  }),
  getters: {
    // Returns true if the user is logged in
    getLoggedIn(state): boolean {
      console.log("called", state.profile);
      const localProfile = window.localStorage.getItem("profile");
      if (localProfile) {
        this.profile = JSON.parse(localProfile);
        return true;
      }
      return state.profile.Name !== undefined && state.profile.Name !== "";
    },
    // Getter for if user is a lead
    getIsLead(state): boolean {
      if (this.getLoggedIn && state.profile.RawData) {
        return (
          state.profile.RawData.roles.find(
            (role: string) => role === import.meta.env.VITE_LEAD_ROLE
          ) !== undefined
        );
      }
      return false;
    },
    getIsGuest(state): boolean {
      console.log(state.profile.RawData.roles);
      return (
        state.profile.Name === "Guest" ||
        state.profile.RawData.roles.length === 0
      );
    },
  },
  actions: {
    // Loads the profile of the logged in user based on the cookie
    async loadProfile(): Promise<void> {
      if (this.external) {
        return;
      }
      this.external = true;
      console.log("Loading profile...");
      return axios
        .get("/api/info")
        .then(({ data }) => {
          this.profile = data;
          localStorage.setItem("profile", JSON.stringify(data));
        })
        .catch((err) => {
          localStorage.removeItem("profile");
          this.profile = {} as Profile;
        });
    },
    // Wipes the profile cookie and logs out the user
    async logout() {
      console.log("Logging out...");
      this.profile = {} as Profile;
      document.cookie =
        "auth" + "=; Path=/; Expires=Thu, 01 Jan 1970 00:00:01 GMT;";
      localStorage.removeItem("profile");
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
        .then(({ data }) => {
          this.profile = data;
          localStorage.setItem("profile", JSON.stringify(data));
        })
        .catch((err) => {});
    },
  },
});
