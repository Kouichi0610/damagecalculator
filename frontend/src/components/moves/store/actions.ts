import { ActionTree } from 'vuex';
import axios from 'axios';
import { RootState } from '@/store/types';
import { MovesState, MoveInfo } from './types'

export const actions: ActionTree<MovesState, RootState> = {
  getMoves: ({commit}, name: string): Promise<string> => {
    console.log('get_moves...');

    return new Promise((resolve, reject) => {
      axios.get('moves', {
        params: {
          Name: name,
        }
      })
      .then((response) => {
        let json = JSON.stringify(response.data);
        console.log(json);
        let moveArray = JSON.parse(json);
        for (var i = 0; i < moveArray.length; i++) {
          let d = new MoveInfo(moveArray[i]);
          console.log('' + d.ToString());
        }
        resolve('');
      })
      .catch((e) => {
          console.log('failed:' + e);
          reject('');
      });
    });
  },
}

