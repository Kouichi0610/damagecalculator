import { MutationTree } from 'vuex';
import { SpeedOrderState, SpeedInfo, SpeedCorrector } from './types';

export const mutations: MutationTree<SpeedOrderState> = {
  setTargetSpeed(state, payload: number) {
    state.targetSpeed = payload;
  },
  setInfos(state, payload: SpeedInfo[]) {
    state.list = payload;
  },
  abilityOwner(state, payload: SpeedCorrector) {
    state.abilityOwner = payload;
  },
  abilityOther(state, payload: SpeedCorrector) {
    state.abilityOther = payload;
  },
  item(state, payload: SpeedCorrector[]) {
    state.item = payload;
  }
}
