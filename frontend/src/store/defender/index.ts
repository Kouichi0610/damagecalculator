import { Module } from 'vuex'
import { getters } from './getters';
import { actions } from './actions';
import { mutations } from './mutations';
import { DefenderState } from './types';
import { RootState } from '@/store/types';

export const state: DefenderState = {
}

const namespaced: boolean = true;

export const defenderState: Module<DefenderState, RootState> = {
  namespaced,
  state,
  getters,
  actions,
  mutations,
}
