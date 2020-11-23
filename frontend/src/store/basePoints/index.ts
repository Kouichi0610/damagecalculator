import { Module } from 'vuex'
import { getters } from './getters';
import { mutations } from './mutations';
import { BasePoints, BasePointsState } from './types';
import { RootState } from '@/store/types';

export const state: BasePointsState = {
  basePoints: new BasePoints()
}

const namespaced: boolean = true;

export const basePointsState: Module<BasePointsState, RootState> = {
  namespaced,
  state,
  getters,
  mutations,
}
