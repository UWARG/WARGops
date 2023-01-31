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
        }
    },
    actions: {
        setCode(code: string) {
            this.profile.code = code;
        },
        setProfile(profile: Profile) {
            this.profile = profile;
        },
        loadProfile(): Promise<void> {
            return axios.get('http://localhost:8080/user')
                .then((res) => {
                    this.profile = res.data;
                })
                .catch((err) => {
                    console.log(err);
                });
        }
    }
});