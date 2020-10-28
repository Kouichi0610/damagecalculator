import { ActionTree } from 'vuex';
import axios from 'axios';
import { RootState } from '@/store/types';
import { NatureState } from './types';
import { Nature } from './types';

export const actions: ActionTree<NatureState, RootState> = {
  initialize: ({commit}) => {
    axios.get('nature_list')
    .then((response) => {
      let json = JSON.stringify(response.data);
      let descriptions = JSON.parse(json);
      let list: Nature[] = [];
      for (var i = 0; i < descriptions.length; i++) {
        let d = descriptions[i];
        let name = d.name;
        let desc = d.description;
        list.push({name: name, description: desc});
      }
      commit('setlist', list);
      return true;
    })
    .catch((e) => {
      console.log('failed:' + e);
      return false;
    });
  }

}
