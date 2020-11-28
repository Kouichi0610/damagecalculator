import { GetterTree } from 'vuex';
import { WeatherFieldsState, Weather, Field } from './types';
import { RootState } from '@/store/types'

export const getters: GetterTree<WeatherFieldsState, RootState> = {
  isLoaded: (state: WeatherFieldsState): boolean => {
    return state.weathers.length > 0 && state.fields.length > 0;
  },
  weathers: (state: WeatherFieldsState): Weather[] => {
    return state.weathers;
  },
  fields: (state: WeatherFieldsState): Field[] => {
    return state.fields;
  },
  currentWeather: (state: WeatherFieldsState): Weather => {
    return state.currentWeather;
  },
  currentField: (state: WeatherFieldsState): Field => {
    return state.currentField;
  }
}
