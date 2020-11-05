/*

*/
import { Module } from 'vuex'
import { getters } from './getters';
import { actions } from './actions';
import { mutations } from './mutations';
import { TargetSelectState } from './types';
import { RootState } from '@/store/types';

export const state: TargetSelectState = {
  initialTotal: 500,
  initialType: 'すべて',
  candidates: [],
}

const namespaced: boolean = true;

export const targetSelect: Module<TargetSelectState, RootState> = {
  namespaced,
  state,
  getters,
  actions,
  mutations,
}
