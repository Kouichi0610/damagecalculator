import { GetterTree } from 'vuex';
import { SpeedListState, SpeedInfo } from './types';
import { RootState } from '@/store/types';

export const getters: GetterTree<SpeedListState, RootState> = {
  hasList: (state: SpeedListState) => {
    return state.list.length > 0;
  },
  listCount: (state: SpeedListState) => {
    return state.list.length;
  },
  list: (state: SpeedListState) => {
    return state.list;
  }
}
