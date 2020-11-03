import { Module } from 'vuex'
import { getters } from './getters';
import { actions } from './actions';
import { mutations } from './mutations';
import { SpeedOrderState } from './types';
import { RootState } from '@/store/types';

export const state: SpeedOrderState = {
  list: [],
}

const namespaced: boolean = true;

export const speedOrder: Module<SpeedOrderState, RootState> = {
  namespaced,
  state,
  getters,
  actions,
  mutations,
}
