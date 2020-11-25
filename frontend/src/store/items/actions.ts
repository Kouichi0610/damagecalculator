import { ActionTree } from 'vuex';
import axios from 'axios';
import { RootState } from '@/store/types';
import { Item, ItemState } from './types'

export const actions: ActionTree<ItemState, RootState> = {
  initialize: ({ commit }, name: string) => {
    if (name.length == 0) return;
    axios.get('usable_items', {
      params: {
        Name: name,
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
    })
    .catch((e) => {
      console.log('error' + e);
    })
  }
}
