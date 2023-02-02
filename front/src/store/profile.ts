import { defineStore } from 'pinia';
import { Profile } from '@/types';
import axios from 'axios';

export const useProfileStore = defineStore('Profile👨', {
    state: () => ({
        profile: {} as Profile
    }),
    getters: {
        getProfile(state) {
            return state.profile;
        },
        getProfileDefined(state): boolean {
            return state.profile !== undefined;
        },

    },
    actions: {
        setCode(code: string) {
            this.profile.code = code;
        },
        setProfile(profile: Profile) {
            this.profile = profile;
        },
        loadProfile(): Promise<void> {
            console.log("Loading profile...");
            return axios.get('http://localhost:8080/user')
                .then((res) => {
                    this.profile = res.data;
                })
                .catch((err) => {
                    console.log(err);
                });
        },
        getGuilds() {
            const url = new URL('http://localhost:8080/guilds/473584913497718824/roles');
            axios.get(url.toString());
        },
        async checkProfile() {
            console.log("You are logged in:", this.profile.username !== undefined);
            return this.profile.username !== undefined;
        }
    }
});