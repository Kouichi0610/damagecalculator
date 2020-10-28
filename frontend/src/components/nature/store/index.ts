import { Module } from 'vuex'
import { getters } from './getters';
import { actions } from './actions';
import { mutations } from './mutations';
import { NatureState } from './types';
import { RootState } from '@/store/types';

export const state: NatureState = {
  current: '',
  list: [],
}

const namespaced: boolean = true;

export const nature: Module<NatureState, RootState> = {
  namespaced,
  state,
  getters,
  actions,
  mutations,
}
