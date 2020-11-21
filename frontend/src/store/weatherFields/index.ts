import { Module } from 'vuex'
import { getters } from './getters';
import { actions } from './actions';
import { mutations } from './mutations';
import { WeatherFieldsState } from './types';
import { RootState } from '@/store/types';

export const state: WeatherFieldsState = {
  weathers: [],
  fields: [],
}

const namespaced: boolean = true;

export const weatherFields: Module<WeatherFieldsState, RootState> = {
  namespaced,
  state,
  getters,
  actions,
  mutations,
}
