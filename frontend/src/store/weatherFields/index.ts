import { Module } from 'vuex'
import { getters } from './getters';
import { actions } from './actions';
import { mutations } from './mutations';
import { WeatherFieldsState, Weather, Field } from './types';
import { RootState } from '@/store/types';

export const state: WeatherFieldsState = {
  target: '',
  weathers: [],
  fields: [],
  currentWeather: new Weather(0, '', ''),
  currentField: new Field(0, '', ''),
}

const namespaced: boolean = true;

export const weatherFieldsState: Module<WeatherFieldsState, RootState> = {
  namespaced,
  state,
  getters,
  actions,
  mutations,
}
