// チェック対象 候補一覧
import { Module } from 'vuex';
import { getters } from './getters';
import { actions } from './actions';
import { mutations } from './mutations';
import { SpeedListState } from './types';
import { RootState } from '@/store/types';

export const state: SpeedListState = {
  list: [],
}

const namespaced: boolean = true;

export const speedList: Module<SpeedListState, RootState> = {
  namespaced,
  state,
  getters,
  actions,
  mutations,
}
