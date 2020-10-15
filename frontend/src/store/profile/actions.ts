import { ActionTree } from 'vuex';
import axios from 'axios';
import { ProfileState, User } from './types';
import { RootState } from '../types';


export const actions: ActionTree<ProfileState, RootState> = {
    fetchData({ commit }): any {
        // コミットで反映
        commit('profileLoaded', {
            firstName: '茂',
            lastName: '田中',
            email: 'test@email.com',
            phone: '123456789',
        });
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
