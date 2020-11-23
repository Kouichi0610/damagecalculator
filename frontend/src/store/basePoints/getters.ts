import { GetterTree } from 'vuex';
import { IBasePoint, BasePoints, BasePointsState } from './types';
import { RootState } from '@/store/types'

export const getters: GetterTree<BasePointsState, RootState> = {
  basePointsArray: (state: BasePointsState): IBasePoint[] => {
    return state.basePoints.array();
  },
  info: (state: BasePointsState): string => {
    console.log('' + state.basePoints.toString());
    return state.basePoints.toString();
  }
}
