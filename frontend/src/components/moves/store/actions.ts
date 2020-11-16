import { ActionTree } from 'vuex';
import axios from 'axios';
import { RootState } from '@/store/types';
import { MovesState } from './types'

export const actions: ActionTree<MovesState, RootState> = {
  getMoves: ({commit}, name: string): Promise<string> => {
    return new Promise((resolve, reject) => {
      axios.get('moves', {
        params: {
          Name: name,
        }
      })
      .then((response) => {
        let json = JSON.stringify(response.data);
        let moveArray = JSON.parse(json);
        commit('setMoves', moveArray);
        resolve('');
      })
      .catch((e) => {
          console.log('failed:' + e);
          reject('');
      });
    });
  },
}

