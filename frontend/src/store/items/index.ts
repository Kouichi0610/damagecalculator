import { Module } from 'vuex'
import { getters } from './getters';
import { actions } from './actions';
import { mutations } from './mutations';
import { Item, ItemState } from './types';
import { RootState } from '@/store/types';

export const state: ItemState = {
  target: '',
  items: [],
  current: Item.default(),
}

const namespaced: boolean = true;

export const itemState: Module<ItemState, RootState> = {
  namespaced,
  state,
  getters,
  actions,
  mutations,
}
