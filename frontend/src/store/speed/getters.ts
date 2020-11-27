import { GetterTree } from 'vuex';
import { SpeedState, SpeedInfo } from './types';
import { RootState } from '@/store/types'

export const getters: GetterTree<SpeedState, RootState> = {
  list: (state: SpeedState): SpeedInfo[] => {
    return state.otherList;
  },
  targetSpeed: (state: SpeedState): number => {
    return state.targetSpeed;
  },
}
