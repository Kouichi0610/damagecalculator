import { MutationTree } from 'vuex';
import { MovesState, MoveInfo, MoveInfos } from './types';

export const mutations: MutationTree<MovesState> = {
  setMoves(state, payload: MoveInfo[]){
    state.info = new MoveInfos(payload);
  },
}
