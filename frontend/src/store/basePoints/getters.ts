import { GetterTree } from 'vuex';
import { IBasePoint, BasePoints, BasePointsState } from './types';
import { RootState } from '@/store/types'

export const getters: GetterTree<BasePointsState, RootState> = {
  basePoints: (state: BasePointsState): BasePoints => {
    return state.basePoints;
  },
  hp: (state: BasePointsState): IBasePoint => {
    return state.basePoints.hp();
  },
  attack: (state: BasePointsState): IBasePoint => {
    return state.basePoints.attack();
  },
  defense: (state: BasePointsState): IBasePoint => {
    return state.basePoints.defense();
  },
  spAttack: (state: BasePointsState): IBasePoint => {
    return state.basePoints.spAttack();
  },
  spDefense: (state: BasePointsState): IBasePoint => {
    return state.basePoints.spDefense();
  },
  speed: (state: BasePointsState): IBasePoint => {
    return state.basePoints.speed();
  },
  basePointsArray: (state: BasePointsState): IBasePoint[] => {
    return state.basePoints.array();
  },
  info: (state: BasePointsState): string => {
    console.log('' + state.basePoints.toString());
    return state.basePoints.toString();
  }
}
