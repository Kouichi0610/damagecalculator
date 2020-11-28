import { GetterTree } from 'vuex';
import { SpeedState, SpeedInfo, SpeedRanking } from './types';
import { RootState } from '@/store/types'

export const getters: GetterTree<SpeedState, RootState> = {
  speedRanking: (state: SpeedState): SpeedInfo[] => {
    return new SpeedRanking().ranking(state.target, state.targetSpeed, state.trickRoom, state.otherList);
  },
  trickRoom: (state: SpeedState): boolean => {
    return state.trickRoom;
  },
  list: (state: SpeedState): SpeedInfo[] => {
    return state.otherList;
  },
  targetSpeed: (state: SpeedState): number => {
    return state.targetSpeed;
  },
}
