import { GetterTree } from 'vuex';
import { WeatherFieldsState } from './types';
import { RootState } from '@/store/types'

export const getters: GetterTree<WeatherFieldsState, RootState> = {
  isInitialized: (state: WeatherFieldsState): boolean => {
    return state.weathers.length > 0 && state.fields.length > 0;
  },
  weathers: (state: WeatherFieldsState): string[] => {
    return state.weathers;
  },
  fields: (state: WeatherFieldsState): string[] => {
    return state.fields;
  },
  currentWeather: (state: WeatherFieldsState): string => {
    return state.currentWeather;
  },
  currentField: (state: WeatherFieldsState): string => {
    return state.currentField;
  }
}
