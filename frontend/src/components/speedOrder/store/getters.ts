import { GetterTree } from 'vuex';
import { SpeedOrderState } from './types';
import { RootState } from '@/store/types'

export const getters: GetterTree<SpeedOrderState, RootState> = {
  hasList: (state: SpeedOrderState) => {
    return state.list.length > 0;
  },
  list: (state: SpeedOrderState) => {
    return state.list;
  }

}
