import { MutationTree } from 'vuex';
import { Item, ItemState } from './types';

export const mutations: MutationTree<ItemState> = {
  setTarget(state, payload: string) {
    state.target = payload;
  },
  setList(state, payload: Item[]) {
    state.items = payload;
    state.current = payload[0];
  },
  setCurrent(state, payload: Item) {
    state.current = payload;
  },
  clearCurrent(state) {
    if (state.items.length == 0) return;
    state.current = state.items[0];
  }
}
