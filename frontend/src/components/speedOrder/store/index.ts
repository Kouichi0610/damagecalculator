import { Module } from 'vuex'
import { getters } from './getters';
import { actions } from './actions';
import { mutations } from './mutations';
import { SpeedOrderState } from './types';
import { RootState } from '@/store/types';
import { DefaultSpeedCorrector } from './types'

export const state: SpeedOrderState = {
  targetSpeed: 0,
  list: [],
  abilityOwner: DefaultSpeedCorrector(),
  abilityOther: DefaultSpeedCorrector(),
  item: [],
}

const namespaced: boolean = true;

export const speedOrder: Module<SpeedOrderState, RootState> = {
  namespaced,
  state,
  getters,
  actions,
  mutations,
}
