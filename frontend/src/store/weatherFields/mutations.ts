import { MutationTree } from 'vuex';
import { WeatherFieldsState, Weather, Field } from './types';

export const mutations: MutationTree<WeatherFieldsState> = {
  initialize(state, target: string) {
    state.target = target;
    if (state.fields.length == 0 || state.weathers.length == 0) return;
    state.currentWeather = state.weathers[0];
    state.currentField = state.fields[0];
  },
  setWeathers(state, payload: Weather[]) {
    state.weathers = payload;
    state.currentWeather = payload[0];
  },
  setFields(state, payload: Field[]) {
    state.fields = payload;
    state.currentField = payload[0];
  },
  setCurrentWeather(state, payload: Weather) {
    state.currentWeather = payload;
  },
  setCurrentField(state, payload: Field) {
    state.currentField = payload;
  }
}
