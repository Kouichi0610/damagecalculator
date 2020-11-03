import { MutationTree } from 'vuex';
import { SpeedOrderState, SpeedInfo } from './types';

export const mutations: MutationTree<SpeedOrderState> = {
  setInfos(state, payload: SpeedInfo[]) {
    state.list = payload;
  }
}
