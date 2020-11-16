import { MutationTree } from 'vuex';
import { AttackerState, MoveCount } from './types';
import { DefendersResult } from '../../target/store/defenderDamages'
import { MoveInfo } from '@/components/moves/store/types';

export const mutations: MutationTree<AttackerState> = {
  setCurrent(state, payload: number) {
    if (payload < 0 || payload >= MoveCount) {
      return;
    }
    state.current = payload;
  },
  setMove(state, payload: MoveInfo) {
    switch(state.current) {
      case 0:
        state.moveA = payload;
        break;
      case 1:
        state.moveB = payload;
        break;
      case 2:
        state.moveC = payload;
        break;
      case 3:
        state.moveD = payload;
        break;
    }
  },
}
