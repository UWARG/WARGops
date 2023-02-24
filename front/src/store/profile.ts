import { defineStore } from "pinia";
import { Profile } from "../types";
import axios from "axios";

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
      return state.profile.RawData.roles.find((role: string) => role === "1076602113590841467") !== undefined;
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
  },
});
