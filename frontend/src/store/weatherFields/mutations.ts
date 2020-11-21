import { MutationTree } from 'vuex';
import { RootState } from '@/store/types'
import { WeatherFieldsState } from './types';

export const mutations: MutationTree<WeatherFieldsState> = {
  setWeathers(state, payload: string[]) {
    state.weathers = payload;
    state.currentWeather = payload[0];
  },
  setFields(state, payload: string[]) {
    state.fields = payload;
    state.currentField = payload[0];
  },
  setCurrentWeather(state, payload: string) {
    state.currentWeather = payload;
  },
  setCurrentField(state, payload: string) {
    state.currentField = payload;
  }
}
