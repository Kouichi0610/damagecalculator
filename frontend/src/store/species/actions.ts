import { ActionTree } from 'vuex';
import axios from 'axios';
import { RootState } from '@/store/types';
import { SpeciesState } from './types'

export const actions: ActionTree<SpeciesState, RootState> = {
  /*
  initialize: ({ commit }) => {
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
  */
}
