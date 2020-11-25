import { GetterTree } from 'vuex';
import { Item, ItemState } from './types';
import { RootState } from '@/store/types'

export const getters: GetterTree<ItemState, RootState> = {
  items: (state: ItemState): Item[] => {
    return state.items;
  },
  current: (state: ItemState): Item => {
    return state.current;
  }
}
