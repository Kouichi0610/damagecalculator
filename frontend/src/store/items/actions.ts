import { ActionTree } from 'vuex';
import axios from 'axios';
import { RootState } from '@/store/types';
import { Item, ItemState } from './types'
import { state } from './index'

export const actions: ActionTree<ItemState, RootState> = {
  initialize: ({ commit }, target: string) => {
    if (target.length == 0) return;
    if (target == state.target) {
      return;
    }
    axios.get('usable_items', {
      params: {
        Name: target,
      }
    })
    .then((response) => {
      let json = JSON.stringify(response.data);
      let descriptions = JSON.parse(json);
      let list: Item[] = [];

      for (var i = 0; i < descriptions.length; i++) {
        let d = descriptions[i];
        list.push(new Item(d.name, d.description));
      }
      commit('setList', list);
      commit('setTarget', target);
      commit('clearCurrent');
    })
    .catch((e) => {
      console.log('error' + e);
    })
  }
}
