import { GetterTree } from 'vuex';
import { IBasePoint, IBasePoints, BasePointsState } from './types';
import { RootState } from '@/store/types'

export const getters: GetterTree<BasePointsState, RootState> = {
  basePoints: (state: BasePointsState): IBasePoints => {
    return state.basePoints;
  },
  basePointsArray: (state: BasePointsState): IBasePoint[] => {
    return state.basePoints.array();
  },
  info: (state: BasePointsState): string => {
    console.log('' + state.basePoints.toString());
    return state.basePoints.toString();
  }
}
