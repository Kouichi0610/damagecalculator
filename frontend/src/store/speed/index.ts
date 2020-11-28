import { Module } from 'vuex'
import { getters } from './getters';
import { actions } from './actions';
import { mutations } from './mutations';
import { SpeedState } from './types';
import { RootState } from '@/store/types';

export const state: SpeedState = {
  target: '',
  level: 0,
  otherList: [],
  targetSpeed: 0,
  trickRoom: false,
}

const namespaced: boolean = true;

export const speedState: Module<SpeedState, RootState> = {
  namespaced,
  state,
  getters,
  actions,
  mutations,
}
