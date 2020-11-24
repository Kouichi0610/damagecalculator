import { MutationTree } from 'vuex';
import { IBasePoint, BasePoints, BasePointsState } from './types';

export const mutations: MutationTree<BasePointsState> = {
  reset(state) {
    state.basePoints = new BasePoints();
  },
  setHP(state, value: number) {
    state.basePoints.set(0, value);
  },
  setAttack(state, value: number) {
    state.basePoints.set(1, value);
  },
  setDefense(state, value: number) {
    state.basePoints.set(2, value);
  },
  setSpAttack(state, value: number) {
    state.basePoints.set(3, value);
  },
  setSpDefense(state, value: number) {
    state.basePoints.set(4, value);
  },
  setSpeed(state, value: number) {
    state.basePoints.set(5, value);
  },
}
