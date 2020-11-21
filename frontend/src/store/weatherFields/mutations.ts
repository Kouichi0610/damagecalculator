import { MutationTree } from 'vuex';
import { RootState } from '@/store/types'
import { WeatherFieldsState } from './types';

export const mutations: MutationTree<WeatherFieldsState> = {
  setWeathers(state, payload: string[]) {
    state.weathers = payload;
  },
  setFields(state, payload: string[]) {
    state.fields = payload;
  }
}
