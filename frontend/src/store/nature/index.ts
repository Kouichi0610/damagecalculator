import { Module } from 'vuex'
import { getters } from './getters';
import { actions } from './actions';
import { mutations } from './mutations';
import { Nature, NatureState } from './types';
import { RootState } from '@/store/types';

export const state: NatureState = {
  natures: [],
  current: Nature.default(),
}

const namespaced: boolean = true;

export const natureState: Module<NatureState, RootState> = {
  namespaced,
  state,
  getters,
  actions,
  mutations,
}
