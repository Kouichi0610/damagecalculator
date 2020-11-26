import { MutationTree } from 'vuex';
import { AttackerState, Move } from './types';
import { Result } from './sendDamage'

export const mutations: MutationTree<AttackerState> = {
  reset(state) {
    for (var i = 0; i < state.selectedMove.length; i++) {
      state.selectedMove[i] = '';
    }
    state.result = [];
  },
  setCurrent(state, current: number) {
    state.currentIndex = current;
  },
  setMove(state, move: string) {
    state.selectedMove[state.currentIndex] = move;
  },
  setAttacker(state, payload: string) {
    state.attacker = payload;
  },
  setMoves(state, payload: Move[]) {
    state.physicals = [];
    state.specials = [];
    for (var i = 0; i < payload.length; i++) {
      let move = payload[i];
      if (move.isPhysical()) {
        state.physicals.push(move);
      }
      if (move.isSpecial()) {
        state.specials.push(move);
      }
    }
  },
  setResults(state, payload: Result[]) {
    state.result = payload;
  }
}
