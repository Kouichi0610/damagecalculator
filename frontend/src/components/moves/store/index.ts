import { Module } from 'vuex'
import { getters } from './getters';
import { actions } from './actions';
import { mutations } from './mutations';
import { MovesState, MoveInfos } from './types';
import { RootState } from '@/store/types';

export const state: MovesState = {
  info: new MoveInfos([]),
}

const namespaced: boolean = true;

export const moves: Module<MovesState, RootState> = {
  namespaced,
  state,
  getters,
  actions,
  mutations,
}
