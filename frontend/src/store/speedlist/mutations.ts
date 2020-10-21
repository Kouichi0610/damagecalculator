import { MutationTree } from 'vuex';
import { SpeedListState, SpeedInfo } from './types';

export const mutations: MutationTree<SpeedListState> = {
  setInfos(state, payload: SpeedInfo[]) {
    state.list = payload;
  }
}
