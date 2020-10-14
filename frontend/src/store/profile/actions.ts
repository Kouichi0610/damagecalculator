import { ActionTree } from 'vuex';
import axios from 'axios';
import { ProfileState, User } from './types';
import { RootState } from '../types';


class UserData implements User {
    firstName: string = '茂';
    lastName: string = '田中';
    email: string = 'test@email.com';
    phone?: string = '123456789';
}

export const actions: ActionTree<ProfileState, RootState> = {
    fetchData({ commit }): any {
        // コミットで反映
        let user = new UserData();
        commit('profileLoaded', user);
/*
        // sample.
        axios({
            url: 'https://....'
        }).then((response) => {
            const payload: User = response && response.data;
            commit('profileLoaded', payload);
        }, (error) => {
            console.log('fetchData Error:' + error);
            commit('profileError');
        })
        */
    }
}

/*
// vuex/types/index.d.ts
export interface ActionTree<S, R> {
  [key: string]: Action<S, R>;
}
*/
