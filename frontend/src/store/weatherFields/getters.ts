import { GetterTree } from 'vuex';
import { WeatherFieldsState } from './types';
import { RootState } from '@/store/types'

export const getters: GetterTree<WeatherFieldsState, RootState> = {
  weathers: (state: WeatherFieldsState): string[] => {
    return state.weathers;
  },
  fields: (state: WeatherFieldsState): string[] => {
    return state.fields;
  }
}
