import { ActionTree } from 'vuex';
import axios from 'axios';
import { RootState } from '@/store/types';
import { Nature, NatureState } from './types'
import { state } from './index'

export const actions: ActionTree<NatureState, RootState> = {
  initialize: ({commit}, target: string) => {
    if (state.target == target) return;
    commit('setTarget', target);
    commit('clearCurrent');
  },
  loadNatures: ({ commit }) => {
    axios.get('nature_list')
    .then((response) => {
      let json = JSON.stringify(response.data);
      let descriptions = JSON.parse(json);
      let list: Nature[] = [];

      for (var i = 0; i < descriptions.length; i++) {
        let d = descriptions[i];
        list.push(new Nature(d.name, d.description));
      }

      commit('setList', list);
    })
    .catch((e) => {
      console.log('error' + e);
    })
  }
}
